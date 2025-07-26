package bookingmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

type SeatBooking struct {
	SeatBookingId         uint64
	BookingId             uint64
	SeatBookingSeatNumber string
	SeatBookingPrice      decimal.Decimal
	PassengerName         string
	PassengerPhone        string
	SeatBookingCreatedAt  time.Time
	SeatBookingUpdatedAt  time.Time
	SeatBookingDeletedAt  *time.Time
}
