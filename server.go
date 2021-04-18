package main

import (
	"go-graphql-server/controller"
	"go-graphql-server/infrastructure"
	"go-graphql-server/service"
)

func main() {
	initialize()
	for {
	}
}

func initialize() {
	controller.Initialize()
	service.Initialize()
	infrastructure.Initialize()
}
