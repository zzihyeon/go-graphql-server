package main

import (
	"github.com/zzihyeon/go-graphql-server/controller/graphql"
	"github.com/zzihyeon/go-graphql-server/controller/rest"
	"github.com/zzihyeon/go-graphql-server/infrastructure/db"
)

func main() {
	defaultGraphqlPort := "8081"
	defaultRestPort := "8080"
	go rest.Setup(defaultRestPort)
	go graphql.Setup(defaultGraphqlPort)
	go db.Setup()
	for {
	}
}
