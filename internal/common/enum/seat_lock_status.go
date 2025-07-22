package commonenum

// SeatLockStatus indicates the status of a seat.
type SeatLockStatus uint8

const (
	SEATAVAILABLE SeatLockStatus = 1 // Available: Seat is available

	SEATLOCKED SeatLockStatus = 2 // Locked: Seat is locked

	SEATBOOKED SeatLockStatus = 3 // Booked: Seat is booked
)
