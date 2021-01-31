package graphql

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/zzihyeon/go-graphql-server/graph"
	"github.com/zzihyeon/go-graphql-server/graph/generated"
)

type GraphqlController struct{}

func Setup(defaultPort string) {
	var gc GraphqlController
	gc.setup(defaultPort)
}

func (g *GraphqlController) setup(defaultPort string) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	log.Fatal(http.ListenAndServe(":"+defaultPort, nil))
}
