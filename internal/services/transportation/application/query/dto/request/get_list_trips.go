package transportationqueryrequest

import (
	"time"
)

type GetListTripsRequest struct {
	DepartureDate string `json:"departure_date" binding:"required"` // Departure date
	FromLocation  string `json:"from_location" binding:"required"`  // From location
	ToLocation    string `json:"to_location" binding:"required"`    // To location
	Page          int    `json:"page" binding:"required"`           // Page
	PageSize      int    `json:"page_size" binding:"required"`      // Page size
}

func (r *GetListTripsRequest) Validate() map[string]string {
	errors := make(map[string]string)

	// Validate departure_date format
	departureDate, err := time.Parse("2006-01-02", r.DepartureDate)
	if err != nil {
		errors["departure_date"] = "must be in format YYYY-MM-DD"
	} else {
		today := time.Now().Truncate(24 * time.Hour)
		if departureDate.Before(today) {
			errors["departure_date"] = "must not be before today"
		}
	}
	if len(errors) > 0 {
		return errors
	}
	return nil
}
