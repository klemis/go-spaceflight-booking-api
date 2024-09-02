package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	result, err := h.BookingService.CreateBooking(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create booking: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
