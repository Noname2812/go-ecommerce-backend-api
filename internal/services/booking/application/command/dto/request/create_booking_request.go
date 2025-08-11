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

		// validate passenger name
		if len(strings.TrimSpace(seat.PassengerName)) < 2 {
			errors[fieldPrefix+".passenger_name"] = "Passenger name must be at least 2 characters"
		}
		if len(seat.PassengerName) > 100 {
			errors[fieldPrefix+".passenger_name"] = "Passenger name cannot exceed 100 characters"
		}
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
