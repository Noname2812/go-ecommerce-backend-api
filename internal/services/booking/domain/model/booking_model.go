package bookingmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

// Booking model
type Booking struct {
	BookingId           string          // primary key
	TripId              uint64          // trip id
	UserId              *uint64         // user id
	BookingTotalPrice   decimal.Decimal // total price (ex: 1000000 VND)
	BookingStatus       int8            // status (ex: 1 = pending, 2 = confirmed, 3 = cancelled)
	BookingContactName  string          // contact name (ex: "John Doe")
	BookingContactPhone string          // contact phone (ex: "0909090909")
	BookingContactEmail string          // contact email (ex: "john.doe@example.com")
	BookingNote         *string         // note (ex: "I have a special request")
	BookingCreatedAt    time.Time       // created at (ex: 2021-01-01 00:00:00)
	BookingUpdatedAt    time.Time       // updated at (ex: 2021-01-01 00:00:00)
	BookingDeletedAt    *time.Time      // deleted at (ex: 2021-01-01 00:00:00)
}
