package main

import (
	"log"

	"github.com/WahtuAstrawan/go-ecommerce-api/cmd/api"
)

func main() {
	server := api.NewApiServer(":3000", nil)
	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
