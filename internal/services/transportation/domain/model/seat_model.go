package transportationmodel

import (
	"time"

	commonenum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
)

// Seat model
type Seat struct {
	SeatId        uint64              // primary key
	BusId         uint64              // bus id
	SeatNumber    string              // seat number (ex: "A1", "B2")
	SeatRowNo     uint8               // seat row number (ex: 1, 2, 3, ...)
	SeatColumnNo  uint8               // seat column number (ex: 1, 2, 3, ...)
	SeatFloorNo   uint8               // seat floor number (ex: 1, 2, 3, ...)
	SeatType      commonenum.SeatType // seat type (ex: 1: normal, 2: VIP, 3: wheelchair)
	SeatCreatedAt time.Time           // created at (ex: 2021-01-01 00:00:00)
	SeatUpdatedAt time.Time           // updated at (ex: 2021-01-01 00:00:00)
	SeatDeletedAt *time.Time          // deleted at (ex: 2021-01-01 00:00:00)

	// Navigation property - lazy loading
	Bus *Bus // one to many
}
