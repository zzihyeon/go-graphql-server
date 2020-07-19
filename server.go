package main

import (
	"github.com/zzihyeon/go-graphql-server/controller/graphql"
	"github.com/zzihyeon/go-graphql-server/controller/rest"
)

func main() {
	defaultGraphqlPort := "8081"
	defaultRestPort := "8080"
	go graphql.Setup(defaultGraphqlPort)
	go rest.Setup(defaultRestPort)
	for {
	}
}
