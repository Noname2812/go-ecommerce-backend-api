package bookingrepositoryimpl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/database"
	bookingmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/domain/model"
	bookingrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/domain/repository"
	"github.com/shopspring/decimal"
)

type seatBookingRepository struct {
	sqlc *database.Queries
	db   *sql.DB
}

// Createbooking implements transportationrepository.bookingRepository.
func (s *seatBookingRepository) CreateSeatBooking(ctx context.Context, model *bookingmodel.SeatBooking) (uint64, error) {
	txQueries := s.getSeatBookingQueries(ctx)
	data := &database.AddSeatBookingParams{
		BookingID:             int64(model.BookingId),
		SeatBookingSeatNumber: model.SeatBookingSeatNumber,
		SeatBookingPrice:      model.SeatBookingPrice.String(),
		PassengerName:         model.PassengerName,
		PassengerPhone:        model.PassengerPhone,
		SeatBookingCreatedAt:  sql.NullTime{Time: model.SeatBookingCreatedAt, Valid: true},
		SeatBookingUpdatedAt:  sql.NullTime{Time: model.SeatBookingUpdatedAt, Valid: true},
	}

	result, err := txQueries.AddSeatBooking(ctx, *data)
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
func (s *seatBookingRepository) DeleleSeatBooking(ctx context.Context, id uint64) error {
	txQueries := s.getSeatBookingQueries(ctx)
	return txQueries.DeleteBooking(ctx, int64(id))
}

// DeleteForcebooking implements transportationrepository.bookingRepository.
func (s *seatBookingRepository) DeleteForceSeatBooking(ctx context.Context, id uint64) error {
	txQueries := s.getSeatBookingQueries(ctx)
	return txQueries.DeleteForceBooking(ctx, int64(id))
}

// GetById implements transportationrepository.bookingRepository.
func (s *seatBookingRepository) GetSeatBookingById(ctx context.Context, bookingId uint32) (*bookingmodel.SeatBooking, error) {
	result, err := s.sqlc.GetSeatBookingById(ctx, int64(bookingId))
	if err != nil {
		return nil, err
	}
	return &bookingmodel.SeatBooking{
		SeatBookingId:         uint64(result.SeatBookingID),
		BookingId:             uint64(result.BookingID),
		SeatBookingSeatNumber: result.SeatBookingSeatNumber,
		SeatBookingPrice:      decimal.RequireFromString(result.SeatBookingPrice),
		PassengerName:         result.PassengerName,
		PassengerPhone:        result.PassengerPhone,
		SeatBookingCreatedAt:  result.SeatBookingCreatedAt.Time,
		SeatBookingUpdatedAt:  result.SeatBookingUpdatedAt.Time,
	}, nil
}

// Updatebooking implements transportationrepository.bookingRepository.
func (s *seatBookingRepository) UpdateSeatBooking(ctx context.Context, model *bookingmodel.SeatBooking) error {
	txQueries := s.getSeatBookingQueries(ctx)
	params := &database.UpdateSeatBookingParams{}

	rowsAffected, err := txQueries.UpdateSeatBooking(ctx, *params)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("update failed: data was modified by another process")
	}

	return nil
}

func NewSeatBookingRepository(db *sql.DB) bookingrepository.SeatBookingRepository {
	return &seatBookingRepository{
		sqlc: database.New(db),
		db:   db,
	}
}

func (s *seatBookingRepository) getSeatBookingQueries(ctx context.Context) *database.Queries {
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		return s.sqlc.WithTx(tx)
	}
	return s.sqlc
}
