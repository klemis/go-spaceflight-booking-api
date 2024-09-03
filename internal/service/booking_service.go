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
	GetDestinationID(launchpadID string, launchDate time.Time) (models.Destination, error)
	GetLaunchpadID(destinationID models.Destination, launchDate time.Time) (string, error)
	InsertBooking(request models.BookingRequest, launchpadID string) (uint, error)
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
// - add request params validation
// - describe destination_id mapping in docs
// - add booking state ex. "cancelled"
// - implement GetBookings method
// - implement DeleteBooking method
// - add logging
// - add middleware

// CreateBooking creates a new booking.
func (s *bookingService) CreateBooking(request models.BookingRequest) (models.BookingResponse, error) {
	// This function is designed based on the assumption that the destination is more crucial for the user than the launchpad.
	// I created a separate binary for generating schedules (`GenerateSchedules`), that creates schedule only for active launchpads.
	// To simplify, I removed the `LaunchpadID` parameter from the request. Instead, the function retrieves the relevant launchpad
	// from the current schedules. It selects the appropriate launchpad based on the `DestinationID` and `LaunchDate`.
	launchpadID, err := s.GetLaunchpadID(request.DestinationID, request.LaunchDate)
	if err != nil {
		return models.BookingResponse{}, err
	}

	body := prepareRequestBody(launchpadID, request.LaunchDate)
	launches, err := s.externalClient.CheckScheduledLaunches(body)
	if err != nil {
		return models.BookingResponse{}, err
	}
	if len(launches.Docs) != 0 {
		return models.BookingResponse{}, fmt.Errorf("launchpad has already been reserved")
	}

	// Insert booking to bookings table.
	id, err := s.InsertBooking(request, launchpadID)
	if err != nil {
		return models.BookingResponse{}, err
	}

	return models.BookingResponse{
		ID:          id,
		LaunchpadID: launchpadID,
		LaunchDate:  request.LaunchDate,
		Destination: utils.String(request.DestinationID),
	}, nil
}

func (s *bookingService) GetDestinationID(launchpadID string, launchDate time.Time) (models.Destination, error) {
	query := `SELECT destination_id FROM schedules WHERE launchpad_id = $1 AND day_of_week = $2;`
	row := s.db.QueryRow(query, launchpadID, launchDate.Weekday())

	var schedule models.Schedule
	err := row.Scan(&schedule.Destination)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("missing destination for the provided launchpad")
		}

		return 0, err
	}

	return schedule.Destination, nil
}

func (s *bookingService) GetLaunchpadID(destinationID models.Destination, launchDate time.Time) (string, error) {
	query := `SELECT launchpad_id FROM schedules WHERE destination_id = $1 AND day_of_week = $2;`
	row := s.db.QueryRow(query, destinationID, launchDate.Weekday())

	var schedule models.Schedule
	err := row.Scan(&schedule.LaunchpadID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("missing destination for the provided launchpad")
		}

		return "", err
	}

	return schedule.LaunchpadID, nil
}

func (s *bookingService) InsertBooking(request models.BookingRequest, launchpadID string) (uint, error) {
	query := `
        INSERT INTO bookings (first_name, last_name, gender, birthday, launchpad_id, destination_id, launch_date)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING id`

	var id uint
	err := s.db.QueryRow(query,
		request.FirstName,
		request.LastName,
		request.Gender,
		request.Birthday,
		launchpadID,
		request.DestinationID,
		request.LaunchDate,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to insert booking: %w", err)
	}

	return id, nil
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
