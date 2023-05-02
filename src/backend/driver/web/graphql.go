package web

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/rs/zerolog/log"
)

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

	srv.SetRecoverFunc(func(ctx context.Context, err any) error {
		// FIXME:
		if e, ok := err.(error); ok {
			log.Err(e).Msg("[PANIC] failed to graphQL service")
		} else {
			e := fmt.Errorf("%v", err)
			log.Err(e).Msg("[PANIC] failed to graphQL service")
		}
		return nil
	})

	return srv
}
