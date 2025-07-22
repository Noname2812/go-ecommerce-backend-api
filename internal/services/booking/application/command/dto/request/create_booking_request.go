package bookingcommandrequest

type SeatInfo struct {
	SeatNumber     string  `json:"seat_number" binding:"required"`          // seat number (e.g., A1)
	Price          float64 `json:"price" binding:"required"`                // price of seat
	PassengerName  string  `json:"passenger_name" binding:"required"`       // passenger name
	PassengerPhone string  `json:"passenger_phone" binding:"required,e164"` // passenger phone
}

type CreateBookingRequest struct {
	UserID     *string    `json:"user_id,omitempty"`              // user id
	TotalPrice float64    `json:"total_price" binding:"required"` // total price
	Seats      []SeatInfo `json:"seats" binding:"required"`       // seats
	Note       *string    `json:"note,omitempty"`                 // note
	Name       string     `json:"name" binding:"required"`        // name
	Phone      string     `json:"phone" binding:"required,e164"`  // phone
	Email      string     `json:"email" binding:"required,email"` // email
}
