package controller

import (
	"errors"
	"kasir/database/postgres"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type CreatePembayaranReq struct {
	Kode  string `json:"kode" binding:"required"`
	Bayar int32  `json:"bayar" binding:"required"`
}

type CreatePembayaranRes struct {
	MejaNomor    int32
	DipesanPada  time.Time
	Dibayar_pada time.Time
	Item         int64
	Total        int64
	Bayar        int32
	Kembalian    int32
}

func (ctr *Controller) CreatePembayaran(ctx *gin.Context) {
	var req CreatePembayaranReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"status": "bad request",
		})
		return
	}

	response, err := bayar(ctx, ctr.Repo.Queries, req)
	log.Println(err)
	if err != nil {
		ctx.JSON(403, gin.H{
			"status": "forbiden",
			"error":  err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"status": "created",
		"data":   response,
	})
}

func bayar(ctx *gin.Context, q *postgres.Queries, req CreatePembayaranReq) (res CreatePembayaranRes, err error) {
	var pesanan postgres.GetPesananByKDRow
	var pembayaranId int32

	pesanan, err = q.GetPesananByKD(ctx, req.Kode)
	if err != nil {
		log.Println("GetPesananByKD")
		return
	}

	if req.Bayar < int32(pesanan.Total) {
		err = errors.New("jumlah bayar kurang dari total yang harus dibayar")
		log.Println("check jumlah")
		return
	}

	pembayaranId, _ = q.GetPembayaranID(ctx, pesanan.ID)
	if pembayaranId != 0 {
		err = errors.New("pengulangan transaksi")
		return
	}

	dibayarPada := time.Now().UTC()
	err = q.CreatePembayaran(
		ctx, postgres.CreatePembayaranParams{
			PesananID:   pesanan.ID,
			Bayar:       req.Bayar,
			DibayarPada: dibayarPada,
		})

	if err != nil {
		log.Println("create pembayaran")
		return
	}

	res = CreatePembayaranRes{
		MejaNomor:    pesanan.MejaNomor,
		DipesanPada:  pesanan.DipesanPada,
		Dibayar_pada: dibayarPada,
		Item:         pesanan.Item,
		Total:        pesanan.Total,
		Bayar:        req.Bayar,
		Kembalian:    req.Bayar - int32(pesanan.Total),
	}
	log.Println("Selesai")
	return
}
