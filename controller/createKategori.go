package controller

import "github.com/gin-gonic/gin"

type CreateKategoriReq struct {
	Kategori string `json:"kategori" binding:"required"`
}

func (ctr *Controller) CreateKategori(ctx *gin.Context) {
	var req CreateKategoriReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status": "bad request",
		})
		return
	}

	err := ctr.Repo.CreateKategori(ctx, req.Kategori)
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
