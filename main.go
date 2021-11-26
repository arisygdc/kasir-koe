package main

import (
	"kasir/config"
	"log"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(config)
}
