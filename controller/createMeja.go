package controller

import (
	"github.com/gin-gonic/gin"
)

type CreateMejaReq struct {
	Nomor int32 `json:"nomor" binding:"required"`
}

func (ctr *Controller) CreateMeja(ctx *gin.Context) {
	var req CreateMejaReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status": "bad request",
		})
		return
	}

	err := ctr.Repo.Queries.CreateMeja(ctx, req.Nomor)
	if err != nil {
		ctx.JSON(403, gin.H{
			"status": "forbidden",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status": "created",
	})
}
