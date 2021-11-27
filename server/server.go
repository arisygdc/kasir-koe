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
