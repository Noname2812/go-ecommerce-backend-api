package bookingmodel

import (
	"time"

	commonenum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
)

// TripSeatLock model
type TripSeatLock struct {
	TripSeatLockId         uint64                    // primary key
	TripId                 uint64                    // trip id
	TripSeatLockSeatNumber string                    // seat number (ex: "A1", "B2")
	LockedByBookingId      *uint64                   // locked by booking id
	TripSeatLockStatus     commonenum.SeatLockStatus // lock status (ex: 1 = locked, 2 = unlocked)
	TripSeatLockExpiresAt  *time.Time                // lock expires at (ex: 2021-01-01 00:00:00)
	TripSeatLockCreatedAt  time.Time                 // created at (ex: 2021-01-01 00:00:00)
	TripSeatLockUpdatedAt  time.Time                 // updated at (ex: 2021-01-01 00:00:00)

	// Navigation property - lazy loading
	Booking *Booking // one to one
}
