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

type busRepository struct {
	sqlc *database.Queries
	db   *sql.DB
}

// CreateBus implements transportationrepository.BusRepository.
func (b *busRepository) CreateBus(ctx context.Context, model *transportationmodel.Bus) (uint64, error) {
	txQueries := b.getBusQueries(ctx)
	data := &database.AddBusParams{
		BusLicensePlate: model.BusLicensePlate,
		BusCompany:      model.BusCompany,
		BusCapacity:     int32(model.BusCapacity),
		BusPrice:        model.BusPrice.String(),
		BusCreatedAt:    sql.NullTime{Time: model.BusCreatedAt, Valid: true},
		BusUpdatedAt:    sql.NullTime{Time: model.BusUpdatedAt, Valid: true},
	}
	result, err := txQueries.AddBus(ctx, *data)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}

// DeleleBus implements transportationrepository.BusRepository.
func (b *busRepository) DeleleBus(ctx context.Context, id uint64) error {
	txQueries := b.getBusQueries(ctx)
	return txQueries.DeleteBus(ctx, int32(id))
}

// DeleteForceBus implements transportationrepository.BusRepository.
func (b *busRepository) DeleteForceBus(ctx context.Context, id uint64) error {
	txQueries := b.getBusQueries(ctx)
	return txQueries.DeleteForceBus(ctx, int32(id))
}

// GetById implements transportationrepository.BusRepository.
func (b *busRepository) GetBusById(ctx context.Context, id uint32) (*transportationmodel.Bus, error) {
	result, err := b.sqlc.GetBusById(ctx, int32(id))
	if err != nil {
		return nil, err
	}
	return &transportationmodel.Bus{
		BusId:           uint64(result.BusID),
		BusLicensePlate: result.BusLicensePlate,
		BusCompany:      result.BusCompany,
		BusCapacity:     uint8(result.BusCapacity),
		BusCreatedAt:    result.BusCreatedAt.Time,
		BusUpdatedAt:    result.BusUpdatedAt.Time,
		BusPrice:        decimal.RequireFromString(result.BusPrice),
	}, nil
}

// UpdateBus implements transportationrepository.BusRepository.
func (b *busRepository) UpdateBus(ctx context.Context, model *transportationmodel.Bus) error {
	txQueries := b.getBusQueries(ctx)
	params := &database.UpdateBusParams{
		BusLicensePlate: model.BusLicensePlate,
		BusCompany:      model.BusCompany,
		BusPrice:        model.BusPrice.String(),
		BusCapacity:     int32(model.BusCapacity),
		BusID:           int32(model.BusId),
		BusUpdatedAt:    sql.NullTime{Time: model.BusUpdatedAt, Valid: true},
	}

	rowsAffected, err := txQueries.UpdateBus(ctx, *params)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("update failed: data was modified by another process")
	}

	return nil
}

func NewbusRepository(db *sql.DB) transportationrepository.BusRepository {
	return &busRepository{
		sqlc: database.New(db),
		db:   db,
	}
}

func (b *busRepository) getBusQueries(ctx context.Context) *database.Queries {
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		return b.sqlc.WithTx(tx)
	}
	return b.sqlc
}
