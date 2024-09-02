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

// PopulateOption represents the populate option for embedded documents.
type PopulateOption struct {
	Path   string         `json:"path"`
	Select map[string]int `json:"select"`
}

// Options specifies additional query options.
type Options struct {
	Select   map[string]int   `json:"select"`
	Populate []PopulateOption `json:"populate"`
}

type RequestBody struct {
	Query   map[string]interface{} `json:"query"`
	Options Options                `json:"options"`
}

// LaunchpadFiltered represents the launchpad id information.
type LaunchpadFiltered struct {
	ID string `json:"id"`
}

// LaunchesResponse represents the response containing a list of launches.
type LaunchesResponse struct {
	Docs []LaunchpadFiltered `json:"docs"`
}

// Launchpad represents the details of a launchpad.
type Launchpad struct {
	Images struct {
		Large []string `json:"large"`
	} `json:"images"`
	Name            string   `json:"name"`
	FullName        string   `json:"full_name"`
	Locality        string   `json:"locality"`
	Region          string   `json:"region"`
	Latitude        float64  `json:"latitude"`
	Longitude       float64  `json:"longitude"`
	LaunchAttempts  int      `json:"launch_attempts"`
	LaunchSuccesses int      `json:"launch_successes"`
	Rockets         []string `json:"rockets"`
	Timezone        string   `json:"timezone"`
	Launches        []string `json:"launches"`
	Status          string   `json:"status"`
	Details         string   `json:"details"`
	ID              string   `json:"id"`
}
