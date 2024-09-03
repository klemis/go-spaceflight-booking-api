package utils

import (
	"fmt"
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
