package transportationrepository

import (
	"context"

	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
)

type TripSeatLockRepository interface {
	CreateTripSeatLock(ctx context.Context, model *transportationmodel.TripSeatLock) (uint64, error)
	UpdateTripSeatLock(ctx context.Context, model *transportationmodel.TripSeatLock) error
	DeleteForceTripSeatLock(ctx context.Context, id uint64) error
	DeleleTripSeatLock(ctx context.Context, id uint64) error
	GetTripSeatLockById(ctx context.Context, id uint32) (*transportationmodel.TripSeatLock, error)
	CreateOrUpdateSeatLock(ctx context.Context, model *transportationmodel.TripSeatLock) error
	GetListTripSeatLockByTripId(ctx context.Context, tripId uint64) ([]*transportationmodel.TripSeatLock, error)
}
