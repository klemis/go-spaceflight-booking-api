package models

import "time"

type BookingRequest struct {
	ID            uint      `json:"id"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Gender        string    `json:"gender"`
	Birthday      time.Time `json:"birthday"`
	LaunchpadID   string    `json:"launchpad_id"`
	DestinationID int       `json:"destination_id"`
	LaunchDate    time.Time `json:"launch_date"`
}

type BookingResponse struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	LaunchpadID string    `json:"launchpad_id"`
	LaunchDate  time.Time `json:"launch_date"`
	Destination string    `json:"destination"`
}
