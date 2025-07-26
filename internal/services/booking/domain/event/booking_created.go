package bookingdomainevent

import (
	commonkafka "github.com/Noname2812/go-ecommerce-backend-api/internal/common/kafka"
)

type SeatInfo struct {
	SeatNumber     string  `json:"seat_number"`     // seat number (e.g., A1)
	Price          float64 `json:"price"`           // price of seat
	PassengerName  string  `json:"passenger_name"`  // passenger name
	PassengerPhone string  `json:"passenger_phone"` // passenger phone
}

type BookingCreatedPayload struct {
	TripID     string     `json:"trip_id"`           // trip id
	UserID     *string    `json:"user_id,omitempty"` // user id
	TotalPrice float64    `json:"total_price"`       // total price
	Seats      []SeatInfo `json:"seats"`             // seats
	Note       *string    `json:"note,omitempty"`    // note
}

type BookingCreatedEvent struct {
	EventID   string                `json:"event_id"`   // Unique Event ID (UUID)
	BookingID string                `json:"booking_id"` // Booking ID
	Timestamp int64                 `json:"timestamp"`  // Unix timestamp
	Payload   BookingCreatedPayload `json:"payload"`    // payload
}

// EventName
func (e *BookingCreatedEvent) EventName() string {
	return commonkafka.TOPIC_BOOKING_CREATED
}
