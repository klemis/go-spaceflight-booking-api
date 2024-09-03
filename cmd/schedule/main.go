package main

import (
	"log"
	"os"

	"github.com/klemis/go-spaceflight-booking-api/internal/database"
	"github.com/klemis/go-spaceflight-booking-api/internal/external"
	"github.com/klemis/go-spaceflight-booking-api/models"
)

func main() {
	log.Println("Initiating the schedule setup process for launchpads...")
	//dbConnectionString := "host=localhost port=5432 user=admin password=admin dbname=bookings sslmode=disable"
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := database.InitDB(databaseURL)
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer func(db *database.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("failed to close database connection: %v", err)
		}
	}(db)

	externalClient := external.NewSpaceXAPIClient("https://api.spacexdata.com/v4/")

	options := models.Options{
		Select: map[string]int{
			"id": 1,
		},
	}

	body := models.RequestBody{
		Query: map[string]interface{}{
			"state": "active",
		},
		Options: options,
	}

	_, err = externalClient.GetActiveLaunchpads(body)
	if err != nil {
		log.Fatalf("failed to fetch active launchpads: %v", err)
	}

}
