package controller

import (
	"kasir/database/postgres"

	"github.com/gin-gonic/gin"
)

type CreateMenuReq struct {
	Kategori_id int32  `json:"kategori_id" binding:"required"`
	Menu        string `json:"menu" binding:"required"`
	Foto        string `json:"foto" binding:"required"`
	Deskripsi   string `json:"deskripsi" binding:"required"`
	Tersedia    bool   `json:"tersedia" binding:"required"`
	Harga       int32  `json:"harga" binding:"required"`
}

func (ctr *Controller) CreateMenu(ctx *gin.Context) {
	var req CreateMenuReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status": "bad request",
		})
		return
	}

	err := ctr.Repo.Queries.CreateMenu(
		ctx,
		postgres.CreateMenuParams{
			KategoriID: req.Kategori_id,
			Menu:       req.Menu,
			Foto:       req.Foto,
			Deskripsi:  req.Deskripsi,
			Tersedia:   req.Tersedia,
			Harga:      req.Harga,
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
