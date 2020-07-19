package main

import (
	"github.com/zzihyeon/go-graphql-server/controller/graphql"
)

func main() {
	defaultGraphqlPort := "8080"
	// defaultRestPort := "8080"
	go graphql.Setup(defaultGraphqlPort)
	// go rest.Setup(defaultRestPort)
	for {
	}
}
