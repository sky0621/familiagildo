package graphql

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/sky0621/kaubandus/graph"
)

func NewServer() *handler.Server {
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return srv
}
