package rest

import (
	"github.com/gin-gonic/gin"
)

type RestController struct {
	*gin.Engine
	*gin.Context
	gin.HandlerFunc
}

func Setup(defaultPort string) {
	var rc RestController
	rc.setup(defaultPort)
}

func (r *RestController) Default() *gin.Engine {
	return gin.Default()
}

func (rc *RestController) setup(defaultPort string) {
	r := rc.Default()
	r.GET("/ping", rc.pingPong)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func (rc RestController) pingPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
