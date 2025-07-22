package transportationrepositoryimpl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/database"
	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
	transportationrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/repository"
)

type routeRepository struct {
	sqlc *database.Queries
	db   *sql.DB
}

// CreateRoute implements transportationrepository.RouteRepository.
func (r *routeRepository) CreateRoute(ctx context.Context, model *transportationmodel.Route) (uint64, error) {
	txQueries := r.getRouteQueries(ctx)
	data := &database.AddRouteParams{
		RouteStartLocation:     model.RouteStartLocation,
		RouteEndLocation:       model.RouteEndLocation,
		RouteEstimatedDuration: int32(model.RouteEstimatedDuration),
		RouteCreatedAt:         sql.NullTime{Time: model.RouteCreatedAt, Valid: true},
		RouteUpdatedAt:         sql.NullTime{Time: model.RouteUpdatedAt, Valid: true},
	}

	result, err := txQueries.AddRoute(ctx, *data)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// DeleleRoute implements transportationrepository.RouteRepository.
func (r *routeRepository) DeleleRoute(ctx context.Context, RouteId uint64) error {
	txQueries := r.getRouteQueries(ctx)
	return txQueries.DeleteBus(ctx, int32(RouteId))
}

// DeleteForceRoute implements transportationrepository.RouteRepository.
func (r *routeRepository) DeleteForceRoute(ctx context.Context, RouteId uint64) error {
	txQueries := r.getRouteQueries(ctx)
	return txQueries.DeleteForceBus(ctx, int32(RouteId))
}

// GetById implements transportationrepository.RouteRepository.
func (r *routeRepository) GetRouteById(ctx context.Context, RouteId uint32) (*transportationmodel.Route, error) {
	result, err := r.sqlc.GetRouteById(ctx, int32(RouteId))
	if err != nil {
		return nil, err
	}
	return &transportationmodel.Route{
		RouteId:                uint64(result.RouteID),
		RouteStartLocation:     result.RouteStartLocation,
		RouteEndLocation:       result.RouteEndLocation,
		RouteEstimatedDuration: uint64(result.RouteEstimatedDuration),
		RouteCreatedAt:         result.RouteCreatedAt.Time,
		RouteUpdatedAt:         result.RouteUpdatedAt.Time,
	}, nil
}

// UpdateRoute implements transportationrepository.RouteRepository.
func (r *routeRepository) UpdateRoute(ctx context.Context, model *transportationmodel.Route) error {
	txQueries := r.getRouteQueries(ctx)
	params := &database.UpdateRouteParams{
		RouteStartLocation:     model.RouteStartLocation,
		RouteEndLocation:       model.RouteEndLocation,
		RouteEstimatedDuration: int32(model.RouteEstimatedDuration),
		RouteID:                int32(model.RouteId),
	}

	rowsAffected, err := txQueries.UpdateRoute(ctx, *params)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("update failed: data was modified by another process")
	}

	return nil
}

func NewRouteRepository(db *sql.DB) transportationrepository.RouteRepository {
	return &routeRepository{
		sqlc: database.New(db),
		db:   db,
	}
}

func (r *routeRepository) getRouteQueries(ctx context.Context) *database.Queries {
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		return r.sqlc.WithTx(tx)
	}
	return r.sqlc
}
