package controller

import (
	"github.com/gin-gonic/gin"
)

func (Controller) PingPong(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"Message": "pong",
	})
}
