package server

import (
	"kasir/config"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Config config.Config
	Server *gin.Engine
}

func SetupServer(config config.Config) (r *Router) {
	server := gin.Default()

	r = &Router{
		Config: config,
		Server: server,
	}
	return
}

func (R *Router) Run() {
	addr := R.Config.ServerAddress
	R.Server.Run(addr)
}

func (R *Router) Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.Request.Header.Get("API-KEY")
		if apiKey != R.Config.Api_key {
			c.AbortWithStatusJSON(451, gin.H{
				"status": "unavailable",
			})
		}
		c.Next()
	}
}
