package service

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/klemis/go-spaceflight-booking-api/internal/external"
	"github.com/klemis/go-spaceflight-booking-api/internal/utils"
	"github.com/klemis/go-spaceflight-booking-api/models"
)

// BookingService provides methods for booking operations.
type BookingService interface {
	CreateBooking(request models.BookingRequest) (models.BookingResponse, error)
}

// bookingService is an implementation of BookingService.
type bookingService struct {
	externalClient *external.SpaceXAPIClient
	db             *sql.DB
}

// NewBookingService creates a new instance of bookingService with the external client.
func NewBookingService(externalClient *external.SpaceXAPIClient, db *sql.DB) BookingService {
	return &bookingService{
		externalClient: externalClient,
		db:             db,
	}
}

// TODO:
// - implement GetBookings method
// - implement DeleteBooking method
// - add logging
// - add middleware

// CreateBooking creates a new booking.
func (s *bookingService) CreateBooking(request models.BookingRequest) (models.BookingResponse, error) {
	state, err := s.externalClient.CheckLaunchpadState(request.LaunchpadID)
	if err != nil {
		return models.BookingResponse{}, err
	}
	if state != "active" {
		return models.BookingResponse{}, fmt.Errorf("launchpad is not available")
	}

	body := prepareRequestBody(request.LaunchpadID, request.LaunchDate)

	launches, err := s.externalClient.CheckScheduledLaunches(body)
	if err != nil {
		return models.BookingResponse{}, err
	}
	if len(launches.Docs) != 0 {
		return models.BookingResponse{}, fmt.Errorf("launchpad has already been reserved")
	}

	query := `SELECT destination_id FROM schedules WHERE launchpad_id = $1 AND day_of_week = $2;`
	row := s.db.QueryRow(query, request.LaunchpadID, request.LaunchDate.Weekday())

	var schedule models.Schedule
	err = row.Scan(&schedule.Destination)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.BookingResponse{}, nil
		}

		return models.BookingResponse{}, err
	}

	return models.BookingResponse{
		ID:          123,
		LaunchpadID: request.LaunchpadID,
		LaunchDate:  request.LaunchDate,
		Destination: utils.String(schedule.Destination),
	}, nil
}

// prepareRequestBody constructs a RequestBody with extended options.
func prepareRequestBody(launchpadId string, launchDate time.Time) models.RequestBody {
	gte, lt := utils.GetRangeQueryValues(launchDate)

	options := models.Options{
		Select: map[string]int{
			"id": 1,
		},
	}

	body := models.RequestBody{
		Query: map[string]interface{}{
			"launchpad": launchpadId,
			"date_utc": map[string]string{
				"$gte": gte,
				"$lt":  lt,
			},
		},
		Options: options,
	}

	return body
}
