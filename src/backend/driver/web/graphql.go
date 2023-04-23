package web

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/rs/zerolog/log"
	"github.com/sky0621/familiagildo/adapter/controller"
	"github.com/vektah/gqlparser/v2/gqlerror"
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

	srv.SetErrorPresenter(func(ctx context.Context, err error) *gqlerror.Error {
		// FIXME:
		log.Err(err).Msgf("failed to graphQL service: %+v", err)
		controller.AddGraphQLError(ctx, err)
		return nil
	})

	srv.SetRecoverFunc(func(ctx context.Context, err any) error {
		// FIXME:
		controller.AddGraphQLError(ctx, err)
		return nil
	})

	return srv
}
