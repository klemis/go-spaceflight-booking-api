package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	
	"github.com/klemis/go-spaceflight-booking-api/internal/database"
	"github.com/klemis/go-spaceflight-booking-api/internal/external"
	"github.com/klemis/go-spaceflight-booking-api/models"
)

func main() {
	log.Println("Initiating the schedule setup process for launchpads...")

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

	body := prepareRequestBody()
	availableLaunchpads, err := externalClient.GetActiveLaunchpads(body)
	if err != nil {
		log.Fatalf("failed to fetch active launchpads: %v", err)
	}

	schedules := generateSchedule(availableLaunchpads)
	// Insert schedule into database
	if err := insertSchedules(db, schedules); err != nil {
		log.Fatalf("failed to insert schedules: %v", err)
	}
}

// insertSchedules inserts a list of schedules into the database.
func insertSchedules(db *database.DB, schedules []models.Schedule) error {
	query := `
		INSERT INTO schedules (launchpad_id, destination_id, day_of_week, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (launchpad_id, day_of_week) DO UPDATE
		SET destination_id = EXCLUDED.destination_id,
		    updated_at = EXCLUDED.updated_at;` // Update fields in case of conflict

	for _, schedule := range schedules {
		_, err := db.Exec(query,
			schedule.LaunchpadID,
			schedule.Destination,
			schedule.DayOfWeek,
			time.Now(), // created_at
			time.Now(), // updated_at
		)
		if err != nil {
			return fmt.Errorf("failed to insert schedule: %w", err)
		}
	}

	return nil
}

func generateSchedule(availableLaunchpads []models.Filtered) []models.Schedule {
	daysOfWeek := []time.Weekday{
		time.Sunday,
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}

	availableDestinations := []models.Destination{
		models.Mars,
		models.Moon,
		models.Pluto,
		models.AsteroidBelt,
		models.Europa,
		models.Titan,
		models.Ganymede,
	}

	// Create a new random number generator with a source seeded by the current time
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	schedule := make([]models.Schedule, 0)
	for _, launchpad := range availableLaunchpads {
		// Create a copy of the destinations to shuffle
		destinationsCopy := make([]models.Destination, len(availableDestinations))
		copy(destinationsCopy, availableDestinations)
		shuffleDestinations(destinationsCopy, rng)

		for i, day := range daysOfWeek {
			destination := destinationsCopy[i]
			schedule = append(schedule, models.Schedule{
				ID:          uint(len(schedule) + 1),
				LaunchpadID: launchpad.ID,
				Destination: destination,
				DayOfWeek:   day,
			})
		}
	}

	return schedule
}

// shuffleDestinations shuffles the slice of destinations in place.
func shuffleDestinations(destinations []models.Destination, rng *rand.Rand) {
	rng.Shuffle(len(destinations), func(i, j int) {
		destinations[i], destinations[j] = destinations[j], destinations[i]
	})
}

// prepareRequestBody constructs a RequestBody for active launchpads.
func prepareRequestBody() models.RequestBody {
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

	return body
}
