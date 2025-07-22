package bookingrepository

import (
	"context"

	bookingmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/domain/model"
)

type BookingRepository interface {
	CreateBooking(ctx context.Context, model *bookingmodel.Booking) (uint64, error)
	UpdateBooking(ctx context.Context, model *bookingmodel.Booking) error
	DeleteForceBooking(ctx context.Context, id uint64) error
	DeleleBooking(ctx context.Context, id uint64) error
	GetBookingById(ctx context.Context, id uint32) (*bookingmodel.Booking, error)
}
