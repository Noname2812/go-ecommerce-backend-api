package transportationrepositoryimpl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/database"
	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
	transportationrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/repository"
)

type seatRepository struct {
	sqlc *database.Queries
	db   *sql.DB
}

// CreateSeat implements transportationrepository.SeatRepository.
func (t *seatRepository) CreateSeat(ctx context.Context, model *transportationmodel.Seat) (uint64, error) {
	txQueries := t.getSeatQueries(ctx)
	data := &database.AddSeatParams{
		BusID:         int32(model.BusId),
		SeatNumber:    model.SeatNumber,
		IsAvailable:   sql.NullBool{Bool: true, Valid: true},
		SeatCreatedAt: sql.NullTime{Time: model.SeatCreatedAt, Valid: true},
		SeatUpdatedAt: sql.NullTime{Time: model.SeatUpdatedAt, Valid: true},
	}

	result, err := txQueries.AddSeat(ctx, *data)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// DeleleSeat implements transportationrepository.SeatRepository.
func (t *seatRepository) DeleleSeat(ctx context.Context, SeatId uint64) error {
	txQueries := t.getSeatQueries(ctx)
	return txQueries.DeleteSeat(ctx, int32(SeatId))
}

// DeleteForceSeat implements transportationrepository.SeatRepository.
func (t *seatRepository) DeleteForceSeat(ctx context.Context, SeatId uint64) error {
	txQueries := t.getSeatQueries(ctx)
	return txQueries.DeleteForceSeat(ctx, int32(SeatId))
}

// GetById implements transportationrepository.SeatRepository.
func (t *seatRepository) GetSeatById(ctx context.Context, SeatId uint32) (*transportationmodel.Seat, error) {
	result, err := t.sqlc.GetSeatById(ctx, int32(SeatId))
	if err != nil {
		return nil, err
	}
	return &transportationmodel.Seat{
		SeatId:        uint64(result.SeatID),
		BusId:         uint64(result.BusID),
		SeatNumber:    result.SeatNumber,
		SeatAvailable: result.IsAvailable.Bool,
		SeatCreatedAt: result.SeatCreatedAt.Time,
		SeatUpdatedAt: result.SeatUpdatedAt.Time,
	}, nil
}

// UpdateSeat implements transportationrepository.SeatRepository.
func (t *seatRepository) UpdateSeat(ctx context.Context, model *transportationmodel.Seat) error {
	txQueries := t.getSeatQueries(ctx)
	params := &database.UpdateSeatParams{
		SeatNumber:    model.SeatNumber,
		IsAvailable:   sql.NullBool{Bool: model.SeatAvailable, Valid: true},
		SeatID:        int32(model.SeatId),
		BusID:         int32(model.BusId),
		SeatUpdatedAt: sql.NullTime{Time: model.SeatUpdatedAt, Valid: true},
	}

	rowsAffected, err := txQueries.UpdateSeat(ctx, *params)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("update failed: data was modified by another process")
	}

	return nil
}

func NewSeatRepository(db *sql.DB) transportationrepository.SeatRepository {
	return &seatRepository{
		sqlc: database.New(db),
		db:   db,
	}
}

func (t *seatRepository) getSeatQueries(ctx context.Context) *database.Queries {
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		return t.sqlc.WithTx(tx)
	}
	return t.sqlc
}
