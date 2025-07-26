package commonenum

// BookingStatus indicates the status of a booking
type BookingStatus uint8

const (
	BOOKINGPENDING BookingStatus = 1 // PENDING: User has booked the trip but not yet confirmed

	BOOKINGBOOKED BookingStatus = 2 // BOOKED: User has confirmed the booking

	BOOKINGCANCELLED BookingStatus = 3 // CANCELLED: User has cancelled the booking

	BOOKINGCOMPLETED BookingStatus = 4 // COMPLETED: User has completed the trip

	BOOKINGREFUNDED BookingStatus = 5 // REFUNDED: User has been refunded

	BOOKINGEXPIRED BookingStatus = 6 // EXPIRED: Booking has expired
)
