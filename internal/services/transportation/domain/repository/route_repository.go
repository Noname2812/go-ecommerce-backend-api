package transportationrepository

import (
	"context"

	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
)

type RouteRepository interface {
	CreateRoute(ctx context.Context, model *transportationmodel.Route) (uint64, error)
	UpdateRoute(ctx context.Context, model *transportationmodel.Route) error
	DeleteForceRoute(ctx context.Context, id uint64) error
	DeleleRoute(ctx context.Context, id uint64) error
	GetRouteById(ctx context.Context, id uint32) (*transportationmodel.Route, error)
}
