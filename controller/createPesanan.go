package controller

import (
	"kasir/database/postgres"

	"github.com/gin-gonic/gin"
)

type DetailPesanan struct {
	Menu_id int32 `json:"menu_id" binding:"required"`
	Jumlah  int32 `json:"jumlah" binding:"required"`
}

type CreatePesannReq struct {
	Kode       string          `json:"kode" binding:"required"`
	Meja_nomor int32           `json:"meja_nomor" binding:"required"`
	Pesanan    []DetailPesanan `form:"colors[]" json:"pesanan" binding:"required"`
}

func (ctr *Controller) CreatePesanan(ctx *gin.Context) {
	var req CreatePesannReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status": "bad request",
			"pesan":  err,
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
		rbErr := tx.Rollback()
		if rbErr != nil {
			ctx.JSON(500, gin.H{
				"status": "internal server error",
			})
			return
		}
		ctx.JSON(403, gin.H{
			"status": "forbidden",
		})
		return
	}

	pesanan_id, _ := q.GetPesananID(ctx, req.Kode)

	for _, v := range req.Pesanan {
		err := q.CreateDetailPesanan(
			ctx, postgres.CreateDetailPesananParams{
				PesananID: pesanan_id,
				MenuID:    v.Menu_id,
				Jumlah:    v.Jumlah,
			},
		)
		if err != nil {
			rbErr := tx.Rollback()
			if rbErr != nil {
				ctx.JSON(500, gin.H{
					"status": "internal server error",
				})
				return
			}
			ctx.JSON(403, gin.H{
				"status": "forbidden",
			})
			return
		}
	}

	tx.Commit()

	ctx.JSON(201, gin.H{
		"status": "created",
	})
}
