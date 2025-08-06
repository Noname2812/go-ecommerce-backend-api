package transportationrepositoryimpl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/database"
	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
	transportationrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/repository"
	"github.com/shopspring/decimal"
)

type tripRepository struct {
	sqlc *database.Queries
	db   *sql.DB
}

// GetTripDetail implements transportationrepository.TripRepository.
func (t *tripRepository) GetTripDetail(ctx context.Context, id uint64) (*transportationmodel.Trip, error) {
	txQueries := t.getTripQueries(ctx)
	result, err := txQueries.GetTripDetail(ctx, int64(id))
	if err != nil {
		return nil, err
	}
	return &transportationmodel.Trip{
		TripId:            uint64(result.TripID),
		TripDepartureTime: result.TripDepartureTime,
		TripArrivalTime:   result.TripArrivalTime,
		TripBasePrice:     decimal.RequireFromString(result.TripBasePrice),
		Route: &transportationmodel.Route{
			RouteId:            uint64(result.RouteID),
			RouteStartLocation: result.RouteStartLocation,
			RouteEndLocation:   result.RouteEndLocation,
		},
		Bus: &transportationmodel.Bus{
			BusId:           uint64(result.BusID),
			BusLicensePlate: result.BusLicensePlate,
			BusCompany:      result.BusCompany,
			BusCapacity:     uint8(result.BusCapacity),
		},
	}, nil
}

// GetListTripsCount implements transportationrepository.TripRepository.
func (t *tripRepository) GetListTripsCount(ctx context.Context, departureDate time.Time, fromLocation string, toLocation string) (int, error) {
	txQueries := t.getTripQueries(ctx)
	count, err := txQueries.GetListTripsCount(ctx, database.GetListTripsCountParams{
		TripDepartureTime:  departureDate,
		RouteStartLocation: fromLocation,
		RouteEndLocation:   toLocation,
	})
	return int(count), err
}

// GetListTrips implements transportationrepository.TripRepository.
func (t *tripRepository) GetListTrips(ctx context.Context, departureDate time.Time, fromLocation string, toLocation string, page int) ([]transportationmodel.Trip, error) {
	txQueries := t.getTripQueries(ctx)
	trips, err := txQueries.GetListTrips(ctx, database.GetListTripsParams{
		TripDepartureTime:  departureDate,
		RouteStartLocation: fromLocation,
		RouteEndLocation:   toLocation,
		Offset:             int32((page - 1) * 10),
	})
	if err != nil {
		return nil, err
	}
	tripsModel := make([]transportationmodel.Trip, len(trips))
	for i, trip := range trips {
		tripsModel[i] = transportationmodel.Trip{
			TripId:            uint64(trip.TripID),
			TripDepartureTime: trip.TripDepartureTime,
			TripArrivalTime:   trip.TripArrivalTime,
			TripBasePrice:     decimal.RequireFromString(trip.TripBasePrice),
			Route: &transportationmodel.Route{
				RouteStartLocation: trip.RouteStartLocation,
				RouteEndLocation:   trip.RouteEndLocation,
			},
		}
	}
	return tripsModel, nil
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
func (t *tripRepository) GetTripById(ctx context.Context, tripId uint64) (*transportationmodel.Trip, error) {
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
