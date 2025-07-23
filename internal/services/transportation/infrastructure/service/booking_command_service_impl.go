package bookingserviceimpl

import (
	"context"

	bookingcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/application/command/dto/request"
	bookingrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/domain/repository"
	bookingservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/service"
)

type bookingCommandService struct {
	bookingRepo     bookingrepository.BookingRepository
	seatBookingRepo bookingrepository.SeatBookingRepository
}

// CreateBooking implements bookingservice.BookingCommandService.
func (b *bookingCommandService) CreateBooking(ctx context.Context, booking *bookingcommandrequest.CreateBookingRequest) (int, error) {
	panic("unimplemented")
}

func NewBookingCommandService(bookingRepo bookingrepository.BookingRepository, seatBookingRepo bookingrepository.SeatBookingRepository) bookingservice.BookingCommandService {
	return &bookingCommandService{
		bookingRepo:     bookingRepo,
		seatBookingRepo: seatBookingRepo,
	}
}
