package transportationrepository

import (
	"context"

	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
)

type BusRepository interface {
	CreateBus(ctx context.Context, model *transportationmodel.Bus) (uint64, error)
	UpdateBus(ctx context.Context, model *transportationmodel.Bus) error
	DeleteForceBus(ctx context.Context, id uint64) error
	DeleleBus(ctx context.Context, id uint64) error
	GetBusById(ctx context.Context, id uint32) (*transportationmodel.Bus, error)
}
