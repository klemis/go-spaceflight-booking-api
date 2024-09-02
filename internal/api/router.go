package api

import (
	"github.com/gin-gonic/gin"

	"github.com/klemis/go-spaceflight-booking-api/internal/service"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	bookingGroup := v1.Group("/bookings")
	// Create a new booking
	bookingGroup.POST("", service.CreateBooking)
	// Get all bookings
	bookingGroup.GET("", service.GetBookings)
}
