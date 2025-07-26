package bookingrepositoryimpl

import (
	"context"
	"database/sql"
	"fmt"

	commonenum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/database"
	bookingmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/domain/model"
	bookingrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/domain/repository"
)

type tripSeatLockRepository struct {
	sqlc *database.Queries
	db   *sql.DB
}

// Createbooking implements transportationrepository.bookingRepository.
func (t *tripSeatLockRepository) CreateTripSeatLock(ctx context.Context, model *bookingmodel.TripSeatLock) (uint64, error) {
	txQueries := t.getTripSeatLockQueries(ctx)
	data := &database.AddTripSeatLockParams{
		TripID:                 int64(model.TripId),
		TripSeatLockSeatNumber: model.SeatNumber,
		LockedByBookingID:      sql.NullInt64{Int64: int64(*model.LockedByBookingId), Valid: model.LockedByBookingId != nil},
		TripSeatLockStatus:     int8(model.LockStatus),
		TripSeatLockExpiresAt:  sql.NullTime{Time: *model.LockExpiresAt, Valid: model.LockExpiresAt != nil},
		TripSeatLockCreatedAt:  sql.NullTime{Time: model.BookingCreatedAt, Valid: true},
		TripSeatLockUpdatedAt:  sql.NullTime{Time: model.BookingUpdatedAt, Valid: true},
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

// Delelebooking implements transportationrepository.bookingRepository.
func (t *tripSeatLockRepository) DeleleTripSeatLock(ctx context.Context, id uint64) error {
	txQueries := t.getTripSeatLockQueries(ctx)
	return txQueries.DeleteBooking(ctx, int64(id))
}

// DeleteForcebooking implements transportationrepository.bookingRepository.
func (t *tripSeatLockRepository) DeleteForceTripSeatLock(ctx context.Context, id uint64) error {
	txQueries := t.getTripSeatLockQueries(ctx)
	return txQueries.DeleteForceBooking(ctx, int64(id))
}

// GetById implements transportationrepository.bookingRepository.
func (t *tripSeatLockRepository) GetTripSeatLockById(ctx context.Context, bookingId uint32) (*bookingmodel.TripSeatLock, error) {
	result, err := t.sqlc.GetTripSeatLockById(ctx, int64(bookingId))
	if err != nil {
		return nil, err
	}
	return &bookingmodel.TripSeatLock{
		LockId:            uint64(result.TripSeatLockID),
		TripId:            uint64(result.TripID),
		SeatNumber:        result.TripSeatLockSeatNumber,
		LockedByBookingId: utils.NullInt64ToUint64Ptr(result.LockedByBookingID),
		LockStatus:        commonenum.SeatLockStatus(result.TripSeatLockStatus),
		LockExpiresAt:     &result.TripSeatLockExpiresAt.Time,
		BookingCreatedAt:  result.TripSeatLockCreatedAt.Time,
		BookingUpdatedAt:  result.TripSeatLockUpdatedAt.Time,
	}, nil
}

// Updatebooking implements transportationrepository.bookingRepository.
func (t *tripSeatLockRepository) UpdateTripSeatLock(ctx context.Context, model *bookingmodel.TripSeatLock) error {
	txQueries := t.getTripSeatLockQueries(ctx)
	params := &database.UpdateTripSeatLockParams{
		TripSeatLockSeatNumber: model.SeatNumber,
		LockedByBookingID:      sql.NullInt64{Int64: int64(*model.LockedByBookingId)},
		TripSeatLockStatus:     int8(model.LockStatus),
		TripSeatLockExpiresAt:  sql.NullTime{Time: *model.LockExpiresAt, Valid: model.LockExpiresAt != nil},
		TripSeatLockID:         int64(model.TripId),
		TripSeatLockUpdatedAt:  sql.NullTime{Time: model.BookingUpdatedAt, Valid: true},
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

func NewtripSeatLockRepository(db *sql.DB) bookingrepository.TripSeatLockRepository {
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
