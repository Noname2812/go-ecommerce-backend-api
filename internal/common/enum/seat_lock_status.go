package commonenum

// SeatLockStatus indicates the status of a seat.
type SeatLockStatus uint8

const (
	AVAILABLE SeatLockStatus = 1 // Available: Seat is available

	LOCKED SeatLockStatus = 2 // Locked: Seat is locked

	BOOKED SeatLockStatus = 3 // Booked: Seat is booked
)
