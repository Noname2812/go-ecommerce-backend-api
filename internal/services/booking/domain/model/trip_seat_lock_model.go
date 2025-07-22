package bookingmodel

import (
	"time"

	commonenum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
)

type TripSeatLock struct {
	LockId            uint64
	TripId            uint64
	SeatNumber        string
	LockedByBookingId *uint64
	LockStatus        commonenum.SeatLockStatus
	LockExpiresAt     *time.Time
	BookingCreatedAt  time.Time
	BookingUpdatedAt  time.Time
}
