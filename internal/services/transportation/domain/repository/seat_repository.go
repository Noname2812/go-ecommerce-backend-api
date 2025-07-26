package transportationrepository

import (
	"context"

	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
)

type SeatRepository interface {
	CreateSeat(ctx context.Context, model *transportationmodel.Seat) (uint64, error)
	UpdateSeat(ctx context.Context, model *transportationmodel.Seat) error
	DeleteForceSeat(ctx context.Context, id uint64) error
	DeleleSeat(ctx context.Context, id uint64) error
	GetSeatById(ctx context.Context, id uint32) (*transportationmodel.Seat, error)
}
