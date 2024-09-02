package service

import (
	"database/sql"
	"fmt"

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
// - add database postgresql
// - implement GetBookings method
// - implement DeleteBooking method
// - add logging
// - add middleware
// - dockerfile/ docker-compose

// CreateBooking creates a new booking.
func (s *bookingService) CreateBooking(request models.BookingRequest) (models.BookingResponse, error) {
	gte, lt := utils.GetRangeQueryValues(request.LaunchDate)

	body := models.RequestBody{
		Query: map[string]interface{}{
			"launchpad": request.LaunchpadID,
			"date_utc": map[string]interface{}{
				"$gte": gte,
				"$lt":  lt,
			},
		},
		Options: models.Options{
			Select: "id",
		},
	}
	available, err := s.externalClient.CheckLaunchpadAvailability(body)
	if err != nil {
		return models.BookingResponse{}, err
	}
	if !available {
		return models.BookingResponse{}, fmt.Errorf("launchpad is not available")
	}

	return models.BookingResponse{
		ID:          123,
		LaunchpadID: request.LaunchpadID,
	}, nil
}
