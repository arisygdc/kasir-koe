package controller

import (
	"github.com/gin-gonic/gin"
)

func (ctr Controller) GetMeja(ctx *gin.Context) {
	meja, err := ctr.Repo.Queries.GetMejaAll(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status": "internal server error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"status": "ok",
		"data":   meja,
	})
}
