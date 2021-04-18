package rest

import "github.com/gin-gonic/gin"

type RestController struct {
	*gin.Engine
	*gin.Context
	gin.HandlerFunc
}

func Initialize(defaultPort string) {
	var rc RestController
	rc.register(defaultPort)
}
