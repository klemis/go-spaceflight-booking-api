package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/klemis/go-spaceflight-booking-api/internal/api"
)

func main() {
	log.Println("Starting API server...")

	r := gin.New()

	api.SetupRouter(r)

	log.Println("API server listening on port 8080...")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
