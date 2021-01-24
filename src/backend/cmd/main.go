package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.opencensus.io/trace"
	"gocloud.dev/gcp"
	"gocloud.dev/server"
	"gocloud.dev/server/health"
	"gocloud.dev/server/sdserver"
	"golang.org/x/xerrors"
)

const (
	normalEnd   = 0
	abnormalEnd = -1
)

func main() {
	os.Exit(execMain())
}

func execMain() int {
	ctx := context.Background()

	cfg := newConfig()

	db, closeDB, err := setupRDB(cfg)
	if err != nil {
		log.Err(err).Msgf("failed to setupRDB: %+v", err)
		return abnormalEnd
	}
	defer closeDB()

	// FIXME:
	fmt.Println(db != nil)

	svr, shutdownServer, err := setupServer(ctx, cfg, handler.New(nil))
	if err != nil {
		log.Err(err).Msgf("failed to setupServer: %+v", err)
		return abnormalEnd
	}
	defer shutdownServer()

	go func() {
		q := make(chan os.Signal, 1)
		signal.Notify(q, os.Interrupt, syscall.SIGTERM)
		<-q
		closeDB()
		shutdownServer()
		os.Exit(abnormalEnd)
	}()

	/*
	 * start app
	 */
	if err := svr.ListenAndServe((":" + cfg.WebPort)); err != nil {
		log.Err(err).Msgf("failed to start server: %+v", err)
		return abnormalEnd
	}

	return normalEnd
}

func setupRDB(cfg config) (boil.ContextExecutor, func(), error) {
	db, err := sqlx.Connect("postgres", cfg.Dsn())
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to sqlx.Connect: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, nil, xerrors.Errorf("failed to ping: %w", err)
	}

	return db, func() {
		if db != nil {
			if err := db.Close(); err != nil {
				log.Err(err).Send()
			}
		}
	}, nil
}

func setupRouter(es graphql.ExecutableSchema) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	//r.Use(requestCtxLogger())

	r.Use(middleware.Timeout(60 * time.Second))

	r.Handle("/query", graphQlServer(es))

	var workDir string
	{
		var err error
		workDir, err = os.Getwd()
		if err != nil {
			return nil, xerrors.Errorf("failed to Getwd: %w", err)
		}
	}
	log.Info().Msgf("workDir:%s", workDir)

	filesDir := http.Dir(filepath.Join(workDir, "frontend"))
	log.Info().Msgf("filesDir:%#+v", filesDir)

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Str("Host", r.Host).Str("RequestURI", r.RequestURI).Msg("REQUEST")

		ctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(ctx.RoutePattern(), "/*")
		log.Info().Msgf("pathPrefix:%s", pathPrefix)

		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
	})

	return r, nil
}

func setupServer(ctx context.Context, cfg config, es graphql.ExecutableSchema) (*server.Server, func(), error) {
	r, err := setupRouter(es, cfg)
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to setupRouter: %w", err)
	}

	credentials, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to DefaultCredentials: %w", err)
	}

	tokenSource := gcp.CredentialsTokenSource(credentials)

	var projectID gcp.ProjectID
	{
		var err error
		projectID, err = gcp.DefaultProjectID(credentials)
		if err != nil {
			return nil, nil, xerrors.Errorf("failed to DefaultProjectID: %w", err)
		}
	}

	var exporter trace.Exporter
	if cfg.Trace {
		log.Info().Msg("Exporting traces to Stack driver")
		mr := GlobalMonitoredResource{projectID: string(projectID)}
		exporter, _, err = sdserver.NewExporter(projectID, tokenSource, mr)
		if err != nil {
			return nil, nil, xerrors.Errorf("failed to NewExporter: %w", err)
		}
	}

	healthCheck := new(customHealthCheck)
	time.AfterFunc(10*time.Second, func() {
		healthCheck.mu.Lock()
		defer healthCheck.mu.Unlock()
		healthCheck.healthy = true
	})

	options := &server.Options{
		RequestLogger: sdserver.NewRequestLogger(),
		HealthChecks:  []health.Checker{healthCheck},
		TraceExporter: exporter,

		// In production you will likely want to use trace.ProbabilitySampler
		// instead, since AlwaysSample will start and export a trace for every
		// request - this may be prohibitively slow with significant traffic.
		DefaultSamplingPolicy: trace.AlwaysSample(),
		Driver:                &server.DefaultDriver{},
	}

	svr := server.New(r, options)
	return svr, func() {
		if svr != nil {
			if err := svr.Shutdown(ctx); err != nil {
				log.Err(err).Send()
			}
		}
	}, nil
}

type GlobalMonitoredResource struct {
	projectID string
}

func (g GlobalMonitoredResource) MonitoredResource() (string, map[string]string) {
	return "global", map[string]string{"project_id": g.projectID}
}

type customHealthCheck struct {
	mu      sync.RWMutex
	healthy bool
}

func (h *customHealthCheck) CheckHealth() error {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if !h.healthy {
		return errors.New("not ready yet")
	}
	return nil
}

func graphQlServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	var mb int64 = 1 << 20
	srv.AddTransport(transport.MultipartForm{
		MaxMemory:     128 * mb,
		MaxUploadSize: 100 * mb,
	})

	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		// FIXME:
		log.Err(err).Msgf("failed to graphQL service: %+v", err)
		return graphql.DefaultErrorPresenter(ctx, err)
	})

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		// FIXME:
		return xerrors.Errorf("panic occurred: %w", err)
	})

	return srv
}
