package main

import (
	"kasir/config"
	"kasir/controller"
	"kasir/server"
	"log"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	controller := controller.NewController()

	r := server.SetupServer(config)
	r.Server.GET("/ping", controller.PingPong)
	r.Run()
}
