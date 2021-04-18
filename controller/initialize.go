package controller

import (
	"go-graphql-server/controller/graphql"
	"go-graphql-server/controller/rest"
)

func Initialize() {
	defaultGraphqlPort := "8081"
	defaultRestPort := "8080"
	go rest.Initialize(defaultRestPort)
	go graphql.Initialize(defaultGraphqlPort)
}
