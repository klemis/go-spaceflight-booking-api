package models

import "time"

type Destination uint

const (
	Mars         Destination = 1
	Moon         Destination = 2
	Pluto        Destination = 3
	AsteroidBelt Destination = 4
	Europa       Destination = 5
	Titan        Destination = 6
	Ganymede     Destination = 7
)

type BookingRequest struct {
	ID            uint        `gorm:"primary_key" json:"id"`
	FirstName     string      `json:"first_name"`
	LastName      string      `json:"last_name"`
	Gender        string      `json:"gender"`
	Birthday      time.Time   `json:"birthday"`
	LaunchpadID   string      `json:"launchpad_id"`
	DestinationID Destination `json:"destination_id"`
	LaunchDate    time.Time   `json:"launch_date"`
}

type BookingResponse struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	LaunchpadID string `json:"launchpad_id"`
}

type Options struct {
	Select string `json:"select"`
}

type RequestBody struct {
	Query   map[string]interface{} `json:"query"`
	Options Options                `json:"options"`
}

type Launches struct {
	Docs []struct {
		ID string `json:"id"`
	} `json:"docs"`
}
