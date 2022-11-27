package main

import (
	"github.com/amit16110/urlshortener/api"
)

func main() {
	server := api.Server{}

	start := server.Router()

	start.Run(":8080")
}
