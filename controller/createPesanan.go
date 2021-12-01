package controller

import (
	"kasir/database/postgres"

	"github.com/gin-gonic/gin"
)

// type DetailPesanan struct {
// 	Menu_id int32 `json:"menu_id" binding:"required"`
// 	Jumlah  int32 `json:"jumlah" binding:"required"`
// }

type CreatePesannReq struct {
	Kode       string          `json:"kode" binding:"required"`
	Meja_nomor int32           `json:"meja_nomor" binding:"required"`
	Pesanan    map[int32]int32 `json:"pesanan" binding:"required"` // Pesanan map[menu_id]jumlah
}

func (ctr *Controller) CreatePesann(ctx *gin.Context) {
	var req CreatePesannReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status": "bad request",
		})
		return
	}

	tx, err := ctr.Repo.Conn.BeginTx(ctx, nil)
	if err != nil {
		ctx.JSON(403, gin.H{
			"status": "forbidden",
		})
		return
	}

	q := postgres.New(tx)
	err = q.CreatePesanan(
		ctx, postgres.CreatePesananParams{
			Kode:      req.Kode,
			MejaNomor: req.Meja_nomor,
		},
	)
	if err != nil {
		ctx.JSON(403, gin.H{
			"status": "forbidden",
		})
		return
	}

	for m, j := range req.Pesanan {
		err := q.CreateDetailPesanan(
			ctx, postgres.CreateDetailPesananParams{
				PesananID: m,
				MenuID:    m,
				Jumlah:    j,
			},
		)
		if err != nil {
			ctx.JSON(403, gin.H{
				"status": "forbidden",
			})
			return
		}
	}

	ctx.JSON(201, gin.H{
		"status": "created",
	})
}
