package main

import (
	"kasir/config"
	"kasir/controller"
	"kasir/database"
	"kasir/server"
	"log"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgres(config)
	if err != nil {
		log.Fatal(err)
	}

	controller := controller.NewController(repo)

	r := server.SetupServer(config)
	s := r.Server.Group("API").Use(r.Authorization())
	s.GET("/ping", controller.PingPong)
	s.GET("/meja", controller.GetMeja)
	s.GET("/menu", controller.GetMenu)
	s.GET("/riwayat/pesanan", controller.GetPesananHistory)
	s.POST("/meja", controller.CreateMeja)
	s.POST("/kategori", controller.CreateKategori)
	s.POST("/menu", controller.CreateMenu)
	s.POST("/pesan", controller.CreatePesanan)
	s.POST("/bayar", controller.CreatePembayaran)
	r.Run()
}
