package bookingrepository

import (
	"context"

	bookingmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/domain/model"
)

type TripSeatLockRepository interface {
	CreateTripSeatLock(ctx context.Context, model *bookingmodel.TripSeatLock) (uint64, error)
	UpdateTripSeatLock(ctx context.Context, model *bookingmodel.TripSeatLock) error
	DeleteForceTripSeatLock(ctx context.Context, id uint64) error
	DeleleTripSeatLock(ctx context.Context, id uint64) error
	GetTripSeatLockById(ctx context.Context, id uint32) (*bookingmodel.TripSeatLock, error)
}
