package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	"github.com/klemis/go-spaceflight-booking-api/internal/service"
	"github.com/klemis/go-spaceflight-booking-api/models"
)

type Handler struct {
	BookingService service.BookingService
}

// NewHandler creates a new Handler with the provided BookingService.
func NewHandler(bookingService service.BookingService) *Handler {
	return &Handler{
		BookingService: bookingService,
	}
}

// CreateBooking handles the creation of a new booking.
func (h *Handler) CreateBooking(c *gin.Context) {
	var booking models.BookingRequest
	if err := c.BindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		c.Abort()
		return
	}

	validate := validator.New()
	err := validate.Struct(booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body: " + err.Error()})
		c.Abort()
		return
	}

	result, err := h.BookingService.CreateBooking(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create booking: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetBookings handles the retrieval of a list of bookings.
func (h *Handler) GetBookings(c *gin.Context) {
	bookings, err := h.BookingService.GetBookings()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "No bookings found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve bookings: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// DeleteBooking handles the deletion of a booking.
func (h *Handler) DeleteBooking(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		c.Abort()
		return
	}

	err = h.BookingService.DeleteBooking(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Booking not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete booking: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
