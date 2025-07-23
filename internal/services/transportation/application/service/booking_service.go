package bookingservice

import (
	"context"

	bookingcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/application/command/dto/request"
)

type BookingCommandService interface {
	CreateBooking(ctx context.Context, booking *bookingcommandrequest.CreateBookingRequest) (int, error)
}

type BookingQueryService interface{}
