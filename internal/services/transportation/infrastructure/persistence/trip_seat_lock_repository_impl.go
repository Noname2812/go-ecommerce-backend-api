package transportationrepositoryimpl

import (
	"context"
	"database/sql"
	"fmt"

	commonenum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/database"
	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
	transportationrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/repository"
)

type tripSeatLockRepository struct {
	sqlc *database.Queries
	db   *sql.DB
}

// GetListTripSeatLockByTripId implements transportationrepository.TripSeatLockRepository.
func (t *tripSeatLockRepository) GetListTripSeatLockByTripId(ctx context.Context, tripId uint64) ([]*transportationmodel.TripSeatLock, error) {
	txQueries := t.getTripSeatLockQueries(ctx)
	results, err := txQueries.GetMapSeatLockByTripId(ctx, int64(tripId))
	if err != nil {
		return nil, err
	}

	response := make([]*transportationmodel.TripSeatLock, 0)
	for _, result := range results {
		response = append(response, &transportationmodel.TripSeatLock{
			TripSeatLockId:        uint64(result.SeatID),
			LockedByBookingId:     result.LockedByBookingID.String,
			TripSeatLockStatus:    commonenum.SeatLockStatus(result.TripSeatLockStatus),
			TripSeatLockExpiresAt: &result.TripSeatLockExpiresAt.Time,
			Seat: &transportationmodel.Seat{
				SeatId:       uint64(result.SeatID),
				SeatNumber:   result.SeatNumber,
				SeatRowNo:    result.SeatRowNo,
				SeatColumnNo: result.SeatColumnNo,
				SeatFloorNo:  result.SeatFloorNo,
				SeatType:     commonenum.SeatType(result.SeatType),
			},
		})
	}

	return response, nil
}

// CreateOrUpdateSeatLock implements transportationrepository.TripSeatLockRepository.
func (t *tripSeatLockRepository) CreateOrUpdateSeatLock(ctx context.Context, model *transportationmodel.TripSeatLock) error {
	txQueries := t.getTripSeatLockQueries(ctx)
	return txQueries.CreateOrUpdateSeatLock(ctx, database.CreateOrUpdateSeatLockParams{
		TripID:                int64(model.TripId),
		SeatID:                int64(model.SeatId),
		LockedByBookingID:     model.LockedByBookingId,
		TripSeatLockStatus:    uint8(model.TripSeatLockStatus),
		TripSeatLockExpiresAt: sql.NullTime{Time: *model.TripSeatLockExpiresAt, Valid: model.TripSeatLockExpiresAt != nil},
	})
}

// CreateTripSeatLock implements transportationrepository.TripSeatLockRepository.
func (t *tripSeatLockRepository) CreateTripSeatLock(ctx context.Context, model *transportationmodel.TripSeatLock) (uint64, error) {
	txQueries := t.getTripSeatLockQueries(ctx)
	data := &database.AddTripSeatLockParams{
		TripID:                int64(model.TripId),
		SeatID:                int64(model.SeatId),
		LockedByBookingID:     model.LockedByBookingId,
		TripSeatLockStatus:    uint8(model.TripSeatLockStatus),
		TripSeatLockExpiresAt: sql.NullTime{Time: *model.TripSeatLockExpiresAt, Valid: model.TripSeatLockExpiresAt != nil},
	}

	result, err := txQueries.AddTripSeatLock(ctx, *data)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// DeleleTripSeatLock implements transportationrepository.TripSeatLockRepository.
func (t *tripSeatLockRepository) DeleleTripSeatLock(ctx context.Context, id uint64) error {
	txQueries := t.getTripSeatLockQueries(ctx)
	return txQueries.DeleteTripSeatLock(ctx, int64(id))
}

// DeleteForceTripSeatLock implements transportationrepository.TripSeatLockRepository.
func (t *tripSeatLockRepository) DeleteForceTripSeatLock(ctx context.Context, id uint64) error {
	txQueries := t.getTripSeatLockQueries(ctx)
	return txQueries.DeleteForceTripSeatLock(ctx, int64(id))
}

// GetTripSeatLockById implements transportationrepository.TripSeatLockRepository.
func (t *tripSeatLockRepository) GetTripSeatLockById(ctx context.Context, bookingId uint32) (*transportationmodel.TripSeatLock, error) {
	result, err := t.sqlc.GetTripSeatLockById(ctx, int64(bookingId))
	if err != nil {
		return nil, err
	}
	return &transportationmodel.TripSeatLock{
		TripSeatLockId:        uint64(result.TripSeatLockID),
		TripId:                uint64(result.TripID),
		SeatId:                uint64(result.SeatID),
		LockedByBookingId:     result.LockedByBookingID,
		TripSeatLockStatus:    commonenum.SeatLockStatus(result.TripSeatLockStatus),
		TripSeatLockExpiresAt: &result.TripSeatLockExpiresAt.Time,
		TripSeatLockCreatedAt: result.TripSeatLockCreatedAt.Time,
		TripSeatLockUpdatedAt: result.TripSeatLockUpdatedAt.Time,
	}, nil
}

// UpdateTripSeatLock implements transportationrepository.TripSeatLockRepository.
func (t *tripSeatLockRepository) UpdateTripSeatLock(ctx context.Context, model *transportationmodel.TripSeatLock) error {
	txQueries := t.getTripSeatLockQueries(ctx)
	params := &database.UpdateTripSeatLockParams{
		SeatID:                int64(model.SeatId),
		LockedByBookingID:     model.LockedByBookingId,
		TripSeatLockStatus:    uint8(model.TripSeatLockStatus),
		TripSeatLockExpiresAt: sql.NullTime{Time: *model.TripSeatLockExpiresAt, Valid: model.TripSeatLockExpiresAt != nil},
		TripSeatLockID:        int64(model.TripId),
		TripSeatLockUpdatedAt: sql.NullTime{Time: model.TripSeatLockUpdatedAt, Valid: true},
	}

	rowsAffected, err := txQueries.UpdateTripSeatLock(ctx, *params)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("update failed: data was modified by another process")
	}

	return nil
}

func NewTripSeatLockRepository(db *sql.DB) transportationrepository.TripSeatLockRepository {
	return &tripSeatLockRepository{
		sqlc: database.New(db),
		db:   db,
	}
}

func (b *tripSeatLockRepository) getTripSeatLockQueries(ctx context.Context) *database.Queries {
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		return b.sqlc.WithTx(tx)
	}
	return b.sqlc
}
