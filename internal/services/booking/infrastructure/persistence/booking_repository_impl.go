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
	"github.com/shopspring/decimal"
)

type bookingRepository struct {
	sqlc *database.Queries
	db   *sql.DB
}

// Createbooking implements transportationrepository.bookingRepository.
func (t *bookingRepository) CreateBooking(ctx context.Context, model *bookingmodel.Booking) error {
	txQueries := t.getBookingQueries(ctx)
	data := &database.AddBookingParams{
		TripID:              int64(model.TripId),
		UserID:              sql.NullInt64{Int64: int64(*model.UserId), Valid: model.UserId != nil && int64(*model.UserId) > 0},
		BookingTotalPrice:   model.BookingTotalPrice.String(),
		BookingStatus:       int8(commonenum.BookingStatus(model.BookingStatus)),
		BookingContactName:  model.BookingContactName,
		BookingContactPhone: model.BookingContactPhone,
		BookingContactEmail: model.BookingContactEmail,
		BookingNote:         sql.NullString{String: *model.BookingNote, Valid: model.BookingNote != nil},
		BookingCreatedAt:    sql.NullTime{Time: model.BookingCreatedAt, Valid: true},
		BookingUpdatedAt:    sql.NullTime{Time: model.BookingUpdatedAt, Valid: true},
	}

	_, err := txQueries.AddBooking(ctx, *data)
	if err != nil {
		return err
	}
	return nil
}

// Delelebooking implements transportationrepository.bookingRepository.
func (t *bookingRepository) DeleleBooking(ctx context.Context, id string) error {
	txQueries := t.getBookingQueries(ctx)
	return txQueries.DeleteBooking(ctx, string(id))
}

// DeleteForcebooking implements transportationrepository.bookingRepository.
func (t *bookingRepository) DeleteForceBooking(ctx context.Context, bookingId string) error {
	txQueries := t.getBookingQueries(ctx)
	return txQueries.DeleteForceBooking(ctx, string(bookingId))
}

// GetById implements transportationrepository.bookingRepository.
func (t *bookingRepository) GetBookingById(ctx context.Context, bookingId string) (*bookingmodel.Booking, error) {
	result, err := t.sqlc.GetBookingById(ctx, string(bookingId))
	if err != nil {
		return nil, err
	}
	return &bookingmodel.Booking{
		BookingId:           result.BookingID,
		TripId:              uint64(result.TripID),
		UserId:              utils.NullInt64ToUint64Ptr(result.UserID),
		BookingTotalPrice:   decimal.RequireFromString(result.BookingTotalPrice),
		BookingStatus:       result.BookingStatus,
		BookingContactName:  result.BookingContactName,
		BookingContactPhone: result.BookingContactPhone,
		BookingContactEmail: result.BookingContactEmail,
		BookingNote:         utils.NullStringToStringPtr(result.BookingNote),
		BookingCreatedAt:    result.BookingCreatedAt.Time,
		BookingUpdatedAt:    result.BookingUpdatedAt.Time,
	}, nil
}

// Updatebooking implements transportationrepository.bookingRepository.
func (t *bookingRepository) UpdateBooking(ctx context.Context, model *bookingmodel.Booking) error {
	txQueries := t.getBookingQueries(ctx)
	params := &database.UpdateBookingParams{
		BookingTotalPrice:   model.BookingTotalPrice.String(),
		BookingStatus:       int8(commonenum.BookingStatus(model.BookingStatus)),
		BookingContactName:  model.BookingContactName,
		BookingContactPhone: model.BookingContactPhone,
		BookingContactEmail: model.BookingContactEmail,
		BookingNote:         sql.NullString{String: *model.BookingNote, Valid: model.BookingNote != nil},
		BookingID:           model.BookingId,
		BookingUpdatedAt:    sql.NullTime{Time: model.BookingUpdatedAt, Valid: true},
	}

	rowsAffected, err := txQueries.UpdateBooking(ctx, *params)
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("update failed: data was modified by another process")
	}

	return nil
}

func NewBookingRepository(db *sql.DB) bookingrepository.BookingRepository {
	return &bookingRepository{
		sqlc: database.New(db),
		db:   db,
	}
}

func (b *bookingRepository) getBookingQueries(ctx context.Context) *database.Queries {
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		return b.sqlc.WithTx(tx)
	}
	return b.sqlc
}
