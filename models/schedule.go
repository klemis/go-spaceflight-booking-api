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

type Schedule struct {
	ID          uint
	LaunchpadID uint
	Destination Destination
	DayOfWeek   time.Weekday
}
