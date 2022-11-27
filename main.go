package main

import (
	"fmt"

	"github.com/amit16110/urlshortener/api"
)

func main() {
	server := api.Server{}

	start := server.Router()

	err := start.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the server - Error: %v", err))
	}

}
