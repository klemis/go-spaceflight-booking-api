package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/klemis/go-spaceflight-booking-api/internal/api"
	"github.com/klemis/go-spaceflight-booking-api/internal/external"
	"github.com/klemis/go-spaceflight-booking-api/internal/service"
)

func main() {
	log.Println("Starting API server...")
	// Initialize spacex client.
	externalClient := external.NewSpaceXAPIClient("https://api.spacexdata.com/v4/")
	// Initialize the booking service with the spacex external client.
	bookingService := service.NewBookingService(externalClient)
	// Initialize the handler with the booking service.
	handler := api.NewHandler(bookingService)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.POST("/bookings", handler.CreateBooking)

	log.Println("API server listening on port 8080...")
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
