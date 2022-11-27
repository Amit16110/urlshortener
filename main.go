package main

import (
	"fmt"
	"log"

	"github.com/amit16110/urlshortener/api"
)

func main() {
	server := api.Server{}

	start := server.Router()
	fmt.Println("showw")
	log.Panic(start.Run(":8080"), "server running at 8080")

}
