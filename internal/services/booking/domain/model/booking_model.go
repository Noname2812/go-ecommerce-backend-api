package bookingmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

type Booking struct {
	BookingId           uint64
	TripId              uint64
	UserId              *uint64
	BookingTotalPrice   decimal.Decimal
	BookingStatus       int8
	BookingContactName  string
	BookingContactPhone string
	BookingContactEmail string
	BookingNote         *string
	BookingCreatedAt    time.Time
	BookingUpdatedAt    time.Time
	BookingDeletedAt    *time.Time
}
