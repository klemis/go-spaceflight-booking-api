package models

import "time"

type Booking struct {
	ID            uint
	FirstName     string
	LastName      string
	Gender        string
	Birthday      time.Time
	LaunchpadID   string
	DestinationID Destination
	LaunchDate    time.Time
}

type BookingRequest struct {
	FirstName     string      `json:"first_name" validate:"required,min=2,max=50"`
	LastName      string      `json:"last_name" validate:"required,min=2,max=50"`
	Gender        string      `json:"gender" validate:"required,min=2,max=50"`
	Birthday      time.Time   `json:"birthday" validate:"required"`
	DestinationID Destination `json:"destination_id" validate:"required,gte=1,lte=7"`
	LaunchDate    time.Time   `json:"launch_date" validate:"required"`
}

type BookingResponse struct {
	ID          uint      `json:"id"`
	LaunchpadID string    `json:"launchpad_id"`
	LaunchDate  time.Time `json:"launch_date"`
	Destination string    `json:"destination"`
}
