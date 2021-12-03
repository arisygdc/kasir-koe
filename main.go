package main

import (
	"kasir/config"
	"kasir/controller"
	"kasir/database"
	"kasir/server"
	"log"
	"os"
)

func main() {
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	err = publicDir()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := database.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	controller := controller.NewController(repo)

	r := server.SetupServer(cfg)
	s := r.Server.Group("API").Use(r.Authorization())
	s.GET("/ping", controller.PingPong)
	s.GET("/meja", controller.GetMeja)
	s.GET("/menu", controller.GetMenu)
	s.POST("/menu/upload", controller.UploadFotoMenu)
	s.GET("/riwayat/pesanan", controller.GetPesananHistory)
	s.POST("/meja", controller.CreateMeja)
	s.POST("/kategori", controller.CreateKategori)
	s.POST("/menu", controller.CreateMenu)
	s.POST("/pesan", controller.CreatePesanan)
	s.POST("/bayar", controller.CreatePembayaran)
	r.Run()
}

func publicDir() (err error) {
	err = mkdirP(config.GetPathImageMenu())
	return
}

func mkdirP(path string) (err error) {
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.FileMode(0755))
	}
	return
}
