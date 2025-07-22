package transportationmodel

import "time"

type Seat struct {
	SeatId        uint64
	BusId         uint64
	SeatNumber    string
	SeatAvailable bool
	SeatCreatedAt time.Time
	SeatUpdatedAt time.Time
	SeatDeletedAt *time.Time
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
