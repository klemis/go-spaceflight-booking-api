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

// DestinationStrings map to hold string representations
var DestinationStrings = map[Destination]string{
	Mars:         "Mars",
	Moon:         "Moon",
	Pluto:        "Pluto",
	AsteroidBelt: "Asteroid Belt",
	Europa:       "Europa",
	Titan:        "Titan",
	Ganymede:     "Ganymede",
}

type Schedule struct {
	ID          uint
	LaunchpadID string
	Destination Destination
	DayOfWeek   time.Weekday
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
