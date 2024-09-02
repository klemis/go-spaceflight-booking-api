package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/klemis/go-spaceflight-booking-api/internal/api"
	"github.com/klemis/go-spaceflight-booking-api/internal/database"
	"github.com/klemis/go-spaceflight-booking-api/internal/external"
	"github.com/klemis/go-spaceflight-booking-api/internal/service"
)

func main() {
	log.Println("Starting API server...")
	// Initialize the database.
	dbConnectionString := "host=localhost port=5432 user=admin password=admin dbname=bookings sslmode=disable"
	db, err := database.InitDB(dbConnectionString)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("failed to close database connection: %v", err)
		}
	}(db)
	// Initialize spacex client.
	externalClient := external.NewSpaceXAPIClient("https://api.spacexdata.com/v4/")
	// Initialize the booking service with the spacex external client.
	bookingService := service.NewBookingService(externalClient, db.DB)
	// Initialize the handler with the booking service.
	handler := api.NewHandler(bookingService)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.POST("/bookings", handler.CreateBooking)

	log.Println("API server listening on port 8080...")
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
