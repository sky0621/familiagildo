package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
)

const (
	normalEnd   = 0
	abnormalEnd = -1
)

func main() {
	os.Exit(execMain())
}

func execMain() int {
	/*
	 * setup app
	 */
	cfg := newConfig()

	app, shutdownFunc, err := buildApp(context.Background(), cfg)
	if err != nil {
		log.Err(err).Msgf("failed to buildApp: %+v", err)
		return abnormalEnd
	}
	defer shutdownFunc()

	log.Info().Msg(cfg.String())

	go func() {
		q := make(chan os.Signal, 1)
		signal.Notify(q, os.Interrupt, syscall.SIGTERM)
		<-q
		shutdownFunc()
		os.Exit(abnormalEnd)
	}()

	/*
	 * start app
	 */

	return 0
}

func buildApp(ctx context.Context, cfg config) (*app, func(), error) {
	if cfg.IsLocal() {
		return buildLocal(ctx, cfg)
	}
	return build(ctx, cfg)
}
