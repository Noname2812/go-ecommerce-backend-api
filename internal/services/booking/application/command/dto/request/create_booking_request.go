package bookingcommandrequest

import (
	"fmt"
	"strings"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SeatInfo struct {
	SeatNumber     string  `json:"seat_number" validate:"required"`          // seat number (e.g., A1)
	Price          float64 `json:"price" validate:"required"`                // price of seat
	PassengerName  string  `json:"passenger_name" validate:"required"`       // passenger name
	PassengerPhone string  `json:"passenger_phone" validate:"required,e164"` // passenger phone
}

type CreateBookingRequest struct {
	UserID     *string    `json:"user_id,omitempty"`               // user id
	TotalPrice float64    `json:"total_price" validate:"required"` // total price
	Seats      []SeatInfo `json:"seats" validate:"required"`       // seats
	Note       *string    `json:"note,omitempty"`                  // note
	Name       string     `json:"name" validate:"required"`        // name
	Phone      string     `json:"phone" validate:"required,e164"`  // phone
	Email      string     `json:"email" validate:"required,email"` // email
}

func (req *CreateBookingRequest) Validate(ctx *gin.Context) map[string]string {
	validate := ctx.Value("validation").(*validator.Validate)
	errors := utils.ValidateStructWithValidatorTags(validate, req)
	// Custom validations
	req.validateSeats(errors)
	req.validateTotalPrice(errors)
	req.validateName(errors)
	req.validateNote(errors)
	if len(errors) == 0 {
		return nil
	}
	return errors
}

// validateSeats validate seats
func (req *CreateBookingRequest) validateSeats(errors map[string]string) {
	if len(req.Seats) == 0 {
		errors["seats"] = "At least 1 seat must be selected"
		return
	}

	if len(req.Seats) > 10 {
		errors["seats"] = "Cannot book more than 10 seats at once"
		return
	}

	seatNumbers := make(map[string]bool)

	for i, seat := range req.Seats {
		fieldPrefix := fmt.Sprintf("seats[%d]", i)
		// validate duplicate seat number
		if seatNumbers[seat.SeatNumber] {
			errors[fieldPrefix+".seat_number"] = fmt.Sprintf("Seat %s is duplicated", seat.SeatNumber)
		}
		seatNumbers[seat.SeatNumber] = true

		// validate seat price
		if seat.Price <= 0 {
			errors[fieldPrefix+".price"] = "Seat price must be greater than 0"
		}
		if seat.Price > 10000000 { // 10 million
			errors[fieldPrefix+".price"] = "Seat price cannot exceed 10,000,000"
		}

		// validate passenger name
		if len(strings.TrimSpace(seat.PassengerName)) < 2 {
			errors[fieldPrefix+".passenger_name"] = "Passenger name must be at least 2 characters"
		}
		if len(seat.PassengerName) > 100 {
			errors[fieldPrefix+".passenger_name"] = "Passenger name cannot exceed 100 characters"
		}
	}
}

// validateTotalPrice validate total price
func (req *CreateBookingRequest) validateTotalPrice(errors map[string]string) {
	if req.TotalPrice <= 0 {
		errors["total_price"] = "Total price must be greater than 0"
		return
	}

	// calculate total price from seats
	calculatedTotal := 0.0
	for _, seat := range req.Seats {
		calculatedTotal += seat.Price
	}

	// allow small error due to floating point
	tolerance := 0.01
	if calculatedTotal-req.TotalPrice > tolerance || req.TotalPrice-calculatedTotal > tolerance {
		errors["total_price"] = fmt.Sprintf("Total price does not match. Calculated: %.2f, Received: %.2f", calculatedTotal, req.TotalPrice)
	}
}

// validateName validate name
func (req *CreateBookingRequest) validateName(errors map[string]string) {
	if len(strings.TrimSpace(req.Name)) < 2 {
		errors["name"] = "Name must be at least 2 characters"
	}
	if len(req.Name) > 100 {
		errors["name"] = "Name cannot exceed 100 characters"
	}
}

// validateNote validate note
func (req *CreateBookingRequest) validateNote(errors map[string]string) {
	if req.Note != nil && len(*req.Note) > 500 {
		errors["note"] = "Note cannot exceed 500 characters"
	}
}
