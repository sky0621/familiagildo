package setup

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"github.com/rs/zerolog/log"
	"go.opencensus.io/trace"
	"gocloud.dev/gcp"
	"gocloud.dev/server"
	"gocloud.dev/server/health"
	"gocloud.dev/server/sdserver"
	"golang.org/x/xerrors"
	"sync"
	"time"
)

func Server(ctx context.Context, cfg Config, es graphql.ExecutableSchema) (*server.Server, func(), error) {
	r, err := Router(es, cfg.Env)
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
		// instead, since AlwaysSample will start and export graphql_generated.go trace for every
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
