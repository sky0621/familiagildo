package web

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/sky0621/familiagildo/cmd/setup"
	"time"
)

func router(es graphql.ExecutableSchema, env setup.Env) (*chi.Mux, error) {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	//r.Use(requestCtxLogger())

	r.Use(middleware.Timeout(60 * time.Second))

	if env.IsLocal() {
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

		r.HandleFunc("/pg", playground.Handler("GraphQL playground", "/query"))
	}

	r.Handle("/query", graphQlServer(es))

	return r, nil
}
