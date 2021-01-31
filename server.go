package main

import (
	"github.com/zzihyeon/go-graphql-server/controller/graphql"
	"github.com/zzihyeon/go-graphql-server/infrastructure/db"
)

func main() {
	defaultGraphqlPort := "8080"
	// defaultRestPort := "8080"
	go graphql.Setup(defaultGraphqlPort)
	go db.Setup()
	// go rest.Setup(defaultRestPort)
	for {
	}
}
