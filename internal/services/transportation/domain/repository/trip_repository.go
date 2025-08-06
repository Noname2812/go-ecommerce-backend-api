package transportationrepository

import (
	"context"
	"time"

	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
)

type TripRepository interface {
	CreateTrip(ctx context.Context, model *transportationmodel.Trip) (uint64, error)
	UpdateTrip(ctx context.Context, model *transportationmodel.Trip) error
	DeleteForceTrip(ctx context.Context, id uint64) error
	DeleleTrip(ctx context.Context, id uint64) error
	GetTripById(ctx context.Context, id uint64) (*transportationmodel.Trip, error)
	GetListTrips(ctx context.Context, departureDate time.Time, fromLocation string, toLocation string, page int) ([]transportationmodel.Trip, error)
	GetListTripsCount(ctx context.Context, departureDate time.Time, fromLocation string, toLocation string) (int, error)
	GetTripDetail(ctx context.Context, id uint64) (*transportationmodel.Trip, error)
}
