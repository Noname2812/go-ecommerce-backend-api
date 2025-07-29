package bookingmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

// SeatBooking model
type SeatBooking struct {
	SeatBookingId         uint64          // primary key
	BookingId             uint64          // booking id
	SeatBookingSeatNumber string          // seat number (ex: "A1", "B2")
	SeatBookingPrice      decimal.Decimal // price (ex: 1000000 VND)
	PassengerName         string          // passenger name (ex: "John Doe")
	PassengerPhone        string          // passenger phone (ex: "0909090909")
	SeatBookingCreatedAt  time.Time       // created at (ex: 2021-01-01 00:00:00)
	SeatBookingUpdatedAt  time.Time       // updated at (ex: 2021-01-01 00:00:00)
	SeatBookingDeletedAt  *time.Time      // deleted at (ex: 2021-01-01 00:00:00)

	// Navigation property - lazy loading
	Booking *Booking // one to one
}
