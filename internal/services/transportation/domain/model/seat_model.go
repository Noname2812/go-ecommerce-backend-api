package transportationmodel

import "time"

// Seat model
type Seat struct {
	SeatId        uint64     // primary key
	BusId         uint64     // bus id
	SeatNumber    string     // seat number (ex: "A1", "B2")
	SeatAvailable bool       // available (ex: true, false)
	SeatCreatedAt time.Time  // created at (ex: 2021-01-01 00:00:00)
	SeatUpdatedAt time.Time  // updated at (ex: 2021-01-01 00:00:00)
	SeatDeletedAt *time.Time // deleted at (ex: 2021-01-01 00:00:00)

	// Navigation property - lazy loading
	Bus *Bus // one to many
}

func NewSeat(busId uint64, seatNumber string) *Seat {
	return &Seat{
		BusId:         busId,
		SeatNumber:    seatNumber,
		SeatAvailable: true,
		SeatCreatedAt: time.Now(),
		SeatUpdatedAt: time.Now(),
		SeatDeletedAt: nil,
	}
}
