package transportationserviceimpl

import (
	"context"
	"fmt"
	"time"

	commonenum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
	transportationpb "github.com/Noname2812/go-ecommerce-backend-api/internal/common/protogen/transportation"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils"
	cacheservice "github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/cache"
	transportationconsts "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/consts"
	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
	transportationrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/repository"
)

type transportationServiceServer struct {
	transportationpb.UnimplementedTransportationServiceServer
	seatLockRepo       transportationrepository.TripSeatLockRepository
	transactionManager transportationrepository.TransactionManager
	cacheService       cacheservice.RedisCache
}

// GetSeatStatus implements transportation.TransportationServiceServer.
func (t *transportationServiceServer) GetSeatStatus(context.Context, *transportationpb.GetSeatStatusRequest) (*transportationpb.GetSeatStatusResponse, error) {
	panic("unimplemented")
}

// LockSeats implements transportation.TransportationServiceServer.
func (t *transportationServiceServer) LockSeats(ctx context.Context, request *transportationpb.LockSeatsRequest) (*transportationpb.LockSeatsResponse, error) {
	// 1. lock seats on redis
	keys := make([]string, 0)
	for _, seat := range request.Seats {
		keys = append(keys, fmt.Sprintf(transportationconsts.SEATS_LOCK_KEY, request.TripId, seat.SeatNumber))
	}
	script, err := utils.LoadLuaScript("lock_seats.lua")
	if err != nil {
		return nil, err
	}
	result, err := t.cacheService.Eval(ctx, script, keys, request.LockDurationSeconds, request.BookingId)
	if err != nil {
		return nil, err
	}
	if result != 1 {
		return nil, fmt.Errorf("failed to lock seats")
	}
	// 2. save seats to seat lock db
	expiresAt := time.Now().Add(time.Duration(request.LockDurationSeconds) * time.Second)
	err = t.transactionManager.WithTransaction(ctx, func(txCtx context.Context) error {
		for _, seat := range request.Seats {
			seatLock := transportationmodel.TripSeatLock{
				TripId:                request.TripId,
				SeatId:                seat.SeatId,
				LockedByBookingId:     request.BookingId,
				TripSeatLockStatus:    commonenum.LOCKED,
				TripSeatLockExpiresAt: &expiresAt,
			}
			_, err := t.seatLockRepo.CreateTripSeatLock(txCtx, &seatLock)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		// release seats on redis
		script, err := utils.LoadLuaScript("release_lock_seats.lua")
		if err != nil {
			return nil, err
		}
		_, err = t.cacheService.Eval(ctx, script, keys, request.BookingId)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	// 3. return response
	return &transportationpb.LockSeatsResponse{
		Success: true,
	}, nil
}

// UnlockSeats implements transportation.TransportationServiceServer.
func (t *transportationServiceServer) UnlockSeats(context.Context, *transportationpb.UnlockSeatsRequest) (*transportationpb.UnlockSeatsResponse, error) {
	panic("unimplemented")
}

func NewTransportationServiceServer(seatLockRepo transportationrepository.TripSeatLockRepository, transactionManager transportationrepository.TransactionManager, cacheService cacheservice.RedisCache) transportationpb.TransportationServiceServer {
	return &transportationServiceServer{
		seatLockRepo:       seatLockRepo,
		transactionManager: transactionManager,
		cacheService:       cacheService,
	}
}
