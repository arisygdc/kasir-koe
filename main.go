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
	s := r.Server.Group("API")
	s.GET("/ping", controller.PingPong)
	s.POST("/meja", controller.CreateMeja)
	r.Run()
}
