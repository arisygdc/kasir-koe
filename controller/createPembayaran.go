package controller

import (
	"kasir/database/postgres"

	"github.com/gin-gonic/gin"
)

type CreatePembayaranReq struct {
	Kode  string `json:"kode" binding:"required"`
	Bayar int32  `json:"bayar" binding:"required"`
}

func (ctr *Controller) CreatePembayaran(ctx *gin.Context) {
	var req CreatePembayaranReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status": "bad request",
		})
		return
	}

	pesanan_id, err := ctr.Repo.Queries.GetPesananID(ctx, req.Kode)
	if err != nil {
		ctx.JSON(500, gin.H{
			"status": "forbiden",
		})
		return
	}

	err = ctr.Repo.Queries.CreatePembayaran(
		ctx, postgres.CreatePembayaranParams{
			PesananID: pesanan_id,
			Bayar:     req.Bayar,
		})

	if err != nil {
		ctx.JSON(403, gin.H{
			"status": "forbiden",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status": "created",
	})
}
