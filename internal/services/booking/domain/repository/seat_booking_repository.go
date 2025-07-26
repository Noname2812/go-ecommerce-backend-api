package bookingrepository

import (
	"context"

	bookingmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/domain/model"
)

type SeatBookingRepository interface {
	CreateSeatBooking(ctx context.Context, model *bookingmodel.SeatBooking) (uint64, error)
	UpdateSeatBooking(ctx context.Context, model *bookingmodel.SeatBooking) error
	DeleteForceSeatBooking(ctx context.Context, id uint64) error
	DeleleSeatBooking(ctx context.Context, id uint64) error
	GetSeatBookingById(ctx context.Context, id uint32) (*bookingmodel.SeatBooking, error)
}
