package models

import "time"

type Booking struct {
	ID            uint      `gorm:"primary_key" json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Gender        string    `json:"gender"`
	Birthday      time.Time `json:"birthday"`
	LaunchpadID   uint      `json:"launchpad_id"`
	DestinationID uint      `json:"destination_id"`
	LaunchDate    time.Time `json:"launch_date"`
}
