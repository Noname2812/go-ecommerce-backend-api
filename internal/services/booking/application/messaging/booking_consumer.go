package bookingmessaging

type BookingConsumer interface {
	Subscribe() error // Subscribe all topic of service
	// handler

}
