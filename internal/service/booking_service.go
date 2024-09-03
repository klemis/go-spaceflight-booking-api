package service

import (
	"fmt"
	"time"

	"github.com/klemis/go-spaceflight-booking-api/internal/database"
	"github.com/klemis/go-spaceflight-booking-api/internal/external"
	"github.com/klemis/go-spaceflight-booking-api/internal/utils"
	"github.com/klemis/go-spaceflight-booking-api/models"
)

// BookingService provides methods for booking operations.
type BookingService interface {
	GetBookings() ([]models.Booking, error)
	CreateBooking(request models.BookingRequest) (models.Booking, error)
	DeleteBooking(id int) error
}

// bookingService is an implementation of BookingService.
type bookingService struct {
	externalClient *external.SpaceXAPIClient
	db             database.DBInterface
}

// NewBookingService creates a new instance of bookingService with the external client.
func NewBookingService(externalClient *external.SpaceXAPIClient, db database.DBInterface) BookingService {
	return &bookingService{
		externalClient: externalClient,
		db:             db,
	}
}

// TODO:
// - create readme
// - describe destination_id mapping in docs
// - add unit tests

// CreateBooking creates a new booking.
func (s *bookingService) CreateBooking(request models.BookingRequest) (models.Booking, error) {
	// This function is designed based on the assumption that the destination is more crucial for the user than the launchpad.
	// I created a separate binary for generating schedules (`GenerateSchedules`), that creates schedule only for active launchpads.
	// To simplify, I removed the `LaunchpadID` parameter from the request. Instead, the function retrieves the relevant launchpad
	// from the current schedules. It selects the appropriate launchpad based on the `DestinationID` and `LaunchDate`.
	// FIXME: Also consider if it should be a "cancelled" flight.
	launchpadID, err := s.db.GetLaunchpadID(request.DestinationID, request.LaunchDate)
	if err != nil {
		return models.Booking{}, err
	}

	body := prepareRequestBody(launchpadID, request.LaunchDate)
	launches, err := s.externalClient.CheckScheduledLaunches(body)
	if err != nil {
		return models.Booking{}, err
	}
	if len(launches.Docs) != 0 {
		// FIXME: Can we assume that this is a "cancelled" flight? Currently its rather not created at all.
		// - Extend Booking struct by State and set state = "cancelled" here.
		return models.Booking{}, fmt.Errorf("launchpad has already been reserved")
	}

	// Insert booking to bookings table.
	id, err := s.db.InsertBooking(request, launchpadID)
	if err != nil {
		return models.Booking{}, err
	}

	return models.Booking{
		ID:            id,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Gender:        request.Gender,
		Birthday:      request.Birthday,
		LaunchpadID:   launchpadID,
		DestinationID: request.DestinationID,
		LaunchDate:    request.LaunchDate,
	}, nil
}

func (s *bookingService) DeleteBooking(id int) error {
	err := s.db.DeleteBooking(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *bookingService) GetBookings() ([]models.Booking, error) {
	bookings, err := s.db.GetBookings()
	if err != nil {
		return []models.Booking{}, err
	}

	return bookings, nil
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
