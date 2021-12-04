package controller

import (
	"kasir/database/postgres"
	"sync"

	"github.com/gin-gonic/gin"
)

var mux sync.Mutex

type DetailPesanan struct {
	Menu_id int32 `json:"menu_id" binding:"required"`
	Jumlah  int32 `json:"jumlah" binding:"required"`
}

type CreatePesannReq struct {
	Kode       string          `json:"kode" binding:"required"`
	Meja_nomor int32           `json:"meja_nomor" binding:"required"`
	Pesanan    []DetailPesanan `form:"colors[]" json:"pesanan" binding:"required"`
}

type CreatePesananRes struct {
	Kode  string
	Total int32
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
		ctx.JSON(500, gin.H{
			"status": "internal server error",
		})
		return
	}

	q := postgres.New(tx)
	var total int32
	total, err = stmtPesanan(ctx, q, req)
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

	tx.Commit()

	ctx.JSON(201, gin.H{
		"status": "created",
		"data": CreatePesananRes{
			Kode:  req.Kode,
			Total: total,
		},
	})
}

func stmtPesanan(ctx *gin.Context, q *postgres.Queries, req CreatePesannReq) (total int32, err error) {
	mux.Lock()
	defer mux.Unlock()

	err = q.CreatePesanan(
		ctx, postgres.CreatePesananParams{
			Kode:      req.Kode,
			MejaNomor: req.Meja_nomor,
		},
	)

	if err != nil {
		return
	}

	var pesanan_id, harga int32
	pesanan_id, err = q.GetPesananID(ctx, req.Kode)
	if err != nil {
		return
	}

	for _, v := range req.Pesanan {
		harga, err = q.GetHarga(ctx, v.Menu_id)
		if err != nil {
			return
		}

		err = q.CreateDetailPesanan(
			ctx, postgres.CreateDetailPesananParams{
				PesananID: pesanan_id,
				MenuID:    v.Menu_id,
				Harga:     harga,
				Jumlah:    v.Jumlah,
			},
		)
		if err != nil {
			return
		}
		total += harga * v.Jumlah
	}

	return
}
