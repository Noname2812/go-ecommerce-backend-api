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

type seatRepository struct {
	sqlc *database.Queries
	db   *sql.DB
}

// GetListSeatsByBusId implements transportationrepository.SeatRepository.
func (t *seatRepository) GetListSeatsByBusId(ctx context.Context, busId uint64) ([]transportationmodel.Seat, error) {
	txQueries := t.getSeatQueries(ctx)
	seats, err := txQueries.GetListSeatsByBusId(ctx, int32(busId))
	if err != nil {
		return nil, err
	}

	// map response
	response := make([]transportationmodel.Seat, len(seats))
	for i, seat := range seats {
		response[i] = transportationmodel.Seat{
			SeatId:       uint64(seat.SeatID),
			BusId:        uint64(seat.BusID),
			SeatNumber:   seat.SeatNumber,
			SeatRowNo:    uint8(seat.SeatRowNo),
			SeatColumnNo: uint8(seat.SeatColumnNo),
			SeatFloorNo:  uint8(seat.SeatFloorNo),
			SeatType:     commonenum.SeatType(seat.SeatType),
		}
	}
	return response, nil
}

// CreateSeat implements transportationrepository.SeatRepository.
func (t *seatRepository) CreateSeat(ctx context.Context, model *transportationmodel.Seat) (uint64, error) {
	txQueries := t.getSeatQueries(ctx)
	data := &database.AddSeatParams{
		BusID:         int32(model.BusId),
		SeatNumber:    model.SeatNumber,
		SeatRowNo:     uint8(model.SeatRowNo),
		SeatColumnNo:  uint8(model.SeatColumnNo),
		SeatFloorNo:   uint8(model.SeatFloorNo),
		SeatType:      uint8(model.SeatType),
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
		SeatRowNo:     uint8(result.SeatRowNo),
		SeatColumnNo:  uint8(result.SeatColumnNo),
		SeatFloorNo:   uint8(result.SeatFloorNo),
		SeatType:      commonenum.SeatType(result.SeatType),
		SeatCreatedAt: result.SeatCreatedAt.Time,
		SeatUpdatedAt: result.SeatUpdatedAt.Time,
	}, nil
}

// UpdateSeat implements transportationrepository.SeatRepository.
func (t *seatRepository) UpdateSeat(ctx context.Context, model *transportationmodel.Seat) error {
	txQueries := t.getSeatQueries(ctx)
	params := &database.UpdateSeatParams{
		SeatNumber:    model.SeatNumber,
		SeatRowNo:     uint8(model.SeatRowNo),
		SeatColumnNo:  uint8(model.SeatColumnNo),
		SeatFloorNo:   uint8(model.SeatFloorNo),
		SeatType:      uint8(model.SeatType),
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
