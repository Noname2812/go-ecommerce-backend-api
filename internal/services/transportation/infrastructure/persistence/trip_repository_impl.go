package transportationrepositoryimpl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/database"
	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
	transportationrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/repository"
	"github.com/shopspring/decimal"
)

type tripRepository struct {
	sqlc *database.Queries
	db   *sql.DB
}

// CreateTrip implements transportationrepository.TripRepository.
func (t *tripRepository) CreateTrip(ctx context.Context, model *transportationmodel.Trip) (uint64, error) {
	txQueries := t.getTripQueries(ctx)
	data := &database.AddTripParams{
		TripDepartureTime: model.TripDepartureTime,
		TripArrivalTime:   model.TripArrivalTime,
		TripBasePrice:     model.TripBasePrice.String(),
		TripCreatedAt:     sql.NullTime{Time: model.TripCreatedAt, Valid: true},
		TripUpdatedAt:     sql.NullTime{Time: model.TripUpdatedAt, Valid: true},
	}

	result, err := txQueries.AddTrip(ctx, *data)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// DeleleTrip implements transportationrepository.TripRepository.
func (t *tripRepository) DeleleTrip(ctx context.Context, tripId uint64) error {
	txQueries := t.getTripQueries(ctx)
	return txQueries.DeleteTrip(ctx, int64(tripId))
}

// DeleteForceTrip implements transportationrepository.TripRepository.
func (t *tripRepository) DeleteForceTrip(ctx context.Context, tripId uint64) error {
	txQueries := t.getTripQueries(ctx)
	return txQueries.DeleteForceTrip(ctx, int64(tripId))
}

// GetById implements transportationrepository.TripRepository.
func (t *tripRepository) GetTripById(ctx context.Context, tripId uint32) (*transportationmodel.Trip, error) {
	result, err := t.sqlc.GetTripById(ctx, int64(tripId))
	if err != nil {
		return nil, err
	}
	return &transportationmodel.Trip{
		TripId:            uint64(result.TripID),
		RouteId:           uint64(result.RouteID),
		BusId:             uint64(result.BusID),
		TripDepartureTime: result.TripDepartureTime,
		TripArrivalTime:   result.TripArrivalTime,
		TripBasePrice:     decimal.RequireFromString(result.TripBasePrice),
		TripCreatedAt:     result.TripCreatedAt.Time,
		TripUpdatedAt:     result.TripUpdatedAt.Time,
	}, nil
}

// UpdateTrip implements transportationrepository.TripRepository.
func (t *tripRepository) UpdateTrip(ctx context.Context, model *transportationmodel.Trip) error {
	txQueries := t.getTripQueries(ctx)
	params := &database.UpdateTripParams{
		TripDepartureTime: model.TripDepartureTime,
		TripArrivalTime:   model.TripArrivalTime,
		TripBasePrice:     model.TripBasePrice.String(),
		TripID:            int64(model.TripId),
		TripUpdatedAt:     sql.NullTime{Time: model.TripUpdatedAt, Valid: true},
	}

	rowsAffected, err := txQueries.UpdateTrip(ctx, *params)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("update failed: data was modified by another process")
	}

	return nil
}

func NewTripRepository(db *sql.DB) transportationrepository.TripRepository {
	return &tripRepository{
		sqlc: database.New(db),
		db:   db,
	}
}

func (b *tripRepository) getTripQueries(ctx context.Context) *database.Queries {
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		return b.sqlc.WithTx(tx)
	}
	return b.sqlc
}
