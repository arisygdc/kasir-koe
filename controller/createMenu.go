package controller

import (
	"kasir/database/postgres"

	"github.com/gin-gonic/gin"
)

type CreateMenuReq struct {
	Kategori_id int32  `json:"kategori_id" binding:"required"`
	Menu        string `json:"menu" binding:"required"`
}

func (ctr *Controller) CreateMenu(ctx *gin.Context) {
	var req CreateMenuReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status": "bad request",
		})
		return
	}

	err := ctr.Repo.CreateMenu(
		ctx,
		postgres.CreateMenuParams{
			KategoriID: req.Kategori_id,
			Menu:       req.Menu,
		},
	)

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
