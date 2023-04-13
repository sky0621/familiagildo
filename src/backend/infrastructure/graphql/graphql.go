package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/sky0621/kaubandus/adapter/controller"
)

func NewServer() *handler.Server {
	srv := handler.NewDefaultServer(controller.NewExecutableSchema(controller.Config{Resolvers: &controller.Resolver{}}))

	return srv
}
