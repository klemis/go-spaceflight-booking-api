package utils

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/klemis/go-spaceflight-booking-api/models"
)

// String function to get string representation of Destination.
func String(destination models.Destination) string {
	if str, ok := models.DestinationStrings[destination]; ok {
		return str
	}

	return "Unknown"
}

// GetRangeQueryValues creates a range query for a given date to cover the entire day.
func GetRangeQueryValues(dateTime time.Time) (gte, lt string) {
	// Truncate to the start of the day.
	startOfDay := dateTime.Truncate(24 * time.Hour)

	// Add 1 day to get the start of the next day.
	endOfDay := startOfDay.Add(24 * time.Hour)

	gte = fmt.Sprintf(`%s`, startOfDay.Format(time.RFC3339))
	lt = fmt.Sprintf(`%s`, endOfDay.Format(time.RFC3339))

	return
}

func GenerateSchedule(availableLaunchpads []models.Filtered) []models.Schedule {
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
