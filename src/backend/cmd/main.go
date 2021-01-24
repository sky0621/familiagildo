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
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sky0621/kaubandus/adapter/web"
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

	setupLogger(cfg)

	db, closeDB, err := setupRDB(cfg)
	if err != nil {
		log.Err(err).Msgf("failed to setupRDB: %+v", err)
		return abnormalEnd
	}
	defer closeDB()

	// FIXME:
	fmt.Println(db != nil)

	svr, shutdownServer, err := setupServer(ctx, cfg, web.NewExecutableSchema(web.Config{Resolvers: web.NewResolver()}))
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

func setupLogger(cfg config) {
	if cfg.IsCloud() {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.LevelFieldName = "severity"
		log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Caller().Logger()
	} else {
		// Pretty logging
		output := zerolog.ConsoleWriter{Out: os.Stderr}
		output.FormatLevel = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("| %-6s |", i))
		}
		log.Logger = zerolog.New(output).With().Timestamp().Caller().Logger()
	}
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

// サーバーのセッティング（gocloudを使う）
func setupServer(ctx context.Context, cfg config, es graphql.ExecutableSchema) (*server.Server, func(), error) {
	r, err := setupRouter(es, cfg.IsCloud())
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to setupRouter: %w", err)
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

		Driver: &server.DefaultDriver{},
	}

	if cfg.Trace {
		traceExporter, err := setupTraceExporter(ctx)
		if err != nil {

		}
		options.TraceExporter = traceExporter

		// In production you will likely want to use trace.ProbabilitySampler
		// instead, since AlwaysSample will start and export a trace for every
		// request - this may be prohibitively slow with significant traffic.
		options.DefaultSamplingPolicy = trace.AlwaysSample()
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

func setupRouter(es graphql.ExecutableSchema, isCloud bool) (*chi.Mux, error) {
	/*
	 * ルーターのベースセッティング
	 */
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	//r.Use(requestCtxLogger())

	r.Use(middleware.Timeout(60 * time.Second))

	/*
	 * on Cloud は Cloud Run を想定。 nuxt build 結果をルーティング
	 */
	if isCloud {
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
	} else {
		/*
		 * ローカルではフロントエンドを別ポート起動で動作確認する想定なのでCORSを有効にしておく。
		 */
		r.Use(cors.Handler(cors.Options{
			// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
			AllowedOrigins: []string{"*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		/*
		 * GraphQL動作確認用
		 */
		r.HandleFunc("/pg", playground.Handler("GraphQL playground", "/query"))
	}

	r.Handle("/query", graphQlServer(es))

	return r, nil
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

func setupTraceExporter(ctx context.Context) (trace.Exporter, error) {
	credentials, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		return nil, xerrors.Errorf("failed to DefaultCredentials: %w", err)
	}

	projectID, err := gcp.DefaultProjectID(credentials)
	if err != nil {
		return nil, xerrors.Errorf("failed to DefaultProjectID: %w", err)
	}

	tokenSource := gcp.CredentialsTokenSource(credentials)

	mr := GlobalMonitoredResource{projectID: string(projectID)}
	exporter, _, err := sdserver.NewExporter(projectID, tokenSource, mr)
	if err != nil {
		return nil, xerrors.Errorf("failed to NewExporter: %w", err)
	}

	return exporter, err
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
