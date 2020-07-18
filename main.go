package main

import (
	"log"
	"net/http"

	types "github.com/zzihyeon/go-graphql-server/types"

	"github.com/friendsofgo/graphiql"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type query struct{}

func (_ *query) Hello() types.StockType {
	return types.StockType{
		Name:   "jh",
		Per:    "100%",
		Volume: "100000",
	}
}

func main() {
	s := `
                type Query {
                        hello: String!
				}

				type Stocks {
					name: String!
					per: String!
					volume: String!
				}
        `
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})

	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
	if err != nil {
		panic(err)
	}
	http.Handle("/", graphiqlHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
