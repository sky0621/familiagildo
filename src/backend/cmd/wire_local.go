package main

import (
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gocloud.dev/server"
	"golang.org/x/xerrors"
)

// ローカルマシン上で動かす際の固有設定
func buildLocal(ctx context.Context, cfg config) (*app, func(), error) {
	wire.Build(
		setupLocalRDB,
		appSet,
		setupLocalServer,
	)
	return nil, nil, nil
}

func setupLocalRDB(cfg config) (boil.ContextExecutor, func(), error) {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBUser, cfg.DBPass)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to sqlx.Connect: %w", err)
	}

	boil.DebugMode = true

	var loc *time.Location
	loc, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, nil, xerrors.Errorf("failed to time.LoadLocation: %w", err)
	}
	boil.SetLocation(loc)

	return db, func() {
		if db != nil {
			if err := db.Close(); err != nil {
				log.Err(err).Send()
			}
		}
	}, nil
}

func setupLocalServer(ctx context.Context, cfg config, resolver *web.Resolver) (*server.Server, error) {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)

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

	r.HandleFunc("/pg", playground.Handler("GraphQL playground", "/query"))

	r.Handle("/query", graphQlServer(resolver))

	return server.New(r, nil), nil
}
