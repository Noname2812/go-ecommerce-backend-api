package transportationqueryrequest

import (
	"time"

	"github.com/gin-gonic/gin"
)

type GetListTripsRequest struct {
	DepartureDate string `json:"departure_date" validate:"required"` // Departure date
	FromLocation  string `json:"from_location" validate:"required"`  // From location
	ToLocation    string `json:"to_location" validate:"required"`    // To location
	Page          int    `json:"page"`                               // Page
}

func (r *GetListTripsRequest) Validate(ctx *gin.Context) map[string]string {
	errors := make(map[string]string)

	// Validate departure_date format
	departureDate, err := time.Parse(time.DateOnly, r.DepartureDate)
	if err != nil {
		errors["departure_date"] = "Departure date must be in format YYYY-MM-DD"
	} else {
		today := time.Now().Truncate(24 * time.Hour)
		if departureDate.Before(today) {
			errors["departure_date"] = "Departure date must not be before today"
		}
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}
