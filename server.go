package main

import (
	"github.com/zzihyeon/hackernews/controller/graphql"
	"github.com/zzihyeon/hackernews/controller/rest"
)

func main() {
	defaultGraphqlPort := "8081"
	defaultRestPort := "8080"
	go graphql.Setup(defaultGraphqlPort)
	go rest.Setup(defaultRestPort)
	for {
	}
}
