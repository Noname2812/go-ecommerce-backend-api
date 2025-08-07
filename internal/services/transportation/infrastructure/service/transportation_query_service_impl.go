package transportationserviceimpl

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils"
	cacheservice "github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/cache"
	transportationqueryrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/request"
	transportationqueryresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/response"
	transportationservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/service"
	transportationconsts "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/consts"
	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
	transportationrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/repository"
	transportationmapper "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/infrastructure/mapper"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"github.com/panjf2000/ants"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
)

type transportationQueryService struct {
	logger            *zap.Logger
	tripRepo          transportationrepository.TripRepository
	seatRepo          transportationrepository.SeatRepository
	seatLockRepo      transportationrepository.TripSeatLockRepository
	redisCacheService cacheservice.RedisCache
	localCacheService cacheservice.LocalCache
	sfGroup           singleflight.Group
	goroutinePool     *ants.Pool
}

// GetTripDetail implements transportationservice.TransportationQueryService.
func (t *transportationQueryService) GetTripDetail(ctx context.Context, id uint64) (code int, data *transportationqueryresponse.TripDetailResponse, err error) {
	if id == 0 {
		return response.ErrCodeParamInvalid, nil, fmt.Errorf("invalid trip id")
	}
	key := fmt.Sprintf(transportationconsts.TRIP_DETAIL_KEY, id)
	result, err, _ := t.sfGroup.Do(key, func() (interface{}, error) {
		// 1. get trip detail from local cache
		detail, err := t.getTripDetailFromLocalCache(ctx, key)
		if err != nil {
			return nil, err
		}
		if detail != nil {
			return detail, nil
		}
		// 2. get trip detail from redis cache or db
		lockKey := fmt.Sprintf("lock:%s", key)
		return t.redisCacheService.WithDistributedLock(ctx, lockKey, transportationconsts.TRIPS_LOCK_TTL_SECONDS, func(ctx context.Context) (interface{}, error) {
			// Double-check cache after acquiring lock
			cache, err := t.getTripDetailFromLocalCache(ctx, key)
			if err != nil {
				return nil, err
			}
			if cache != nil {
				return cache, nil
			}

			var (
				detail *transportationqueryresponse.TripDetailResponse
				seats  []transportationqueryresponse.Seat
			)
			g, ctx := errgroup.WithContext(ctx)
			// Goroutine 1: Get trip detail
			g.Go(func() error {
				d, err := t.getTripDetail(ctx, id)
				if err != nil {
					return fmt.Errorf("getTripDetail error: %w", err)
				}
				detail = d
				return nil
			})

			// Goroutine 2: Get seat status
			g.Go(func() error {
				s, err := t.getSeatsStatus(ctx, id)
				if err != nil {
					return fmt.Errorf("getSeatsStatus error: %w", err)
				}
				seats = s
				return nil
			})

			// Wait for all goroutines to complete
			if err := g.Wait(); err != nil {
				return nil, err
			}

			// Merge data and save to local cache
			detail.Seats = seats
			isSuccess := t.localCacheService.SetWithTTL(ctx, key, detail, transportationconsts.TRIP_DETAIL_KEY_LOCAL_TTL)
			if !isSuccess {
				t.logger.Error("save trip detail to local cache failed", zap.String("key", key))
			}
			t.logger.Info("save trip detail to local cache success", zap.String("key", key))
			return detail, nil
		})
	})

	if err != nil {
		t.logger.Error("get trip detail failed", zap.Error(err), zap.Uint64("trip_id", id))
		return response.ErrServerError, nil, err
	}

	detailRes, ok := result.(*transportationqueryresponse.TripDetailResponse)
	if !ok {
		return response.ErrServerError, nil, fmt.Errorf("unexpected result type from singleflight")
	}

	return response.ErrCodeSuccess, detailRes, nil
}

// GetListTrips implements transportationservice.TransportationQueryService.
func (t *transportationQueryService) GetListTrips(ctx context.Context, request *transportationqueryrequest.GetListTripsRequest) (int, *transportationqueryresponse.GetListTripsResponse, error) {
	key := fmt.Sprintf(transportationconsts.TRIPS_KEY, request.FromLocation, request.ToLocation, request.DepartureDate, request.Page)
	result, err, _ := t.sfGroup.Do(key, func() (interface{}, error) {
		// 1. get data from local cache first
		tripsLocal, err := t.getListTripsFromLocalCache(ctx, key)
		if err != nil {
			return nil, err
		}
		if tripsLocal != nil {
			return tripsLocal, nil
		}

		// 2. check if data exists in redis cache
		trips, err := t.getListTripsFromRedisCache(ctx, key)
		if err != nil {
			t.logger.Error("get list trips from redis cache failed", zap.Error(err), zap.String("key", key))
			return nil, err
		}
		if trips != nil {
			// Cache warming: save to local cache for next requests
			go func() {
				_ = t.localCacheService.SetWithTTL(ctx, key, trips, transportationconsts.TRIPS_KEY_LOCAL_TTL)
			}()
			return trips, nil
		}

		// 3. Data not in cache, try to acquire distributed lock
		lockKey := fmt.Sprintf("lock:%s", key)
		return t.redisCacheService.WithDistributedLock(ctx, lockKey, transportationconsts.TRIPS_LOCK_TTL_SECONDS, func(ctx context.Context) (interface{}, error) {
			// Double-check cache after acquiring lock
			tripsLocal, err := t.getListTripsFromLocalCache(ctx, key)
			if err != nil {
				return nil, err
			}
			if tripsLocal != nil {
				return tripsLocal, nil
			}

			// Query database
			response, err := t.getTripsAndCountFromDB(ctx, request)
			if err != nil {
				t.logger.Error("get trips and count from db failed", zap.Error(err))
				return nil, err
			}

			// save the data to the cache
			err = t.saveListTripsToRedisCache(ctx, key, response)
			if err != nil {
				t.logger.Error("save cache failed", zap.Error(err), zap.String("key", key))
				return nil, err
			}
			t.logger.Info("save cache success", zap.String("key", key))
			return response, nil
		})
	})

	if err != nil {
		t.logger.Error("get list trips failed", zap.Error(err), zap.String("key", key))
		return response.ErrServerError, nil, err
	}

	tripsRes, ok := result.(*transportationqueryresponse.GetListTripsResponse)
	if !ok {
		return response.ErrServerError, nil, fmt.Errorf("unexpected result type from singleflight")
	}

	return response.ErrCodeSuccess, tripsRes, nil
}

// Cleanup method for graceful shutdown
func (t *transportationQueryService) Close() error {
	if t.goroutinePool != nil {
		t.goroutinePool.Release()
	}
	return nil
}

func NewTransportationQueryService(tripRepo transportationrepository.TripRepository,
	seatRepo transportationrepository.SeatRepository,
	seatLockRepo transportationrepository.TripSeatLockRepository,
	cacheService cacheservice.RedisCache,
	localCacheService cacheservice.LocalCache,
	logger *zap.Logger,
) transportationservice.TransportationQueryService {
	// Create goroutine pool with 100 workers
	pool, err := ants.NewPool(100, ants.WithOptions(ants.Options{
		ExpiryDuration:   10 * time.Second,
		PreAlloc:         true,
		MaxBlockingTasks: 2000,
		Nonblocking:      false,
	}))
	if err != nil {
		logger.Fatal("failed to create goroutine pool", zap.Error(err))
	}
	return &transportationQueryService{
		tripRepo:          tripRepo,
		seatRepo:          seatRepo,
		seatLockRepo:      seatLockRepo,
		redisCacheService: cacheService,
		localCacheService: localCacheService,
		logger:            logger,
		sfGroup:           singleflight.Group{},
		goroutinePool:     pool,
	}
}

// get list trips from redis cache
func (t *transportationQueryService) getListTripsFromRedisCache(ctx context.Context, key string) (*transportationqueryresponse.GetListTripsResponse, error) {
	dataJson, isFound, err := t.redisCacheService.Get(ctx, key) // Get all trips
	if err != nil {
		t.logger.Error("get list trips from cache failed", zap.Error(err), zap.String("key", key))
		return nil, err
	}
	if !isFound {
		return nil, nil
	}
	trips := &transportationqueryresponse.GetListTripsResponse{}
	err = json.Unmarshal([]byte(dataJson), &trips)
	if err != nil {
		t.logger.Error("unmarshal trips failed", zap.Error(err), zap.String("key", key))
		return nil, err
	}
	return trips, nil
}

// get list trips from local cache
func (t *transportationQueryService) getListTripsFromLocalCache(ctx context.Context, key string) (*transportationqueryresponse.GetListTripsResponse, error) {
	value, isFound := t.localCacheService.Get(ctx, key)
	if !isFound {
		return nil, nil
	}
	trips, ok := value.(*transportationqueryresponse.GetListTripsResponse)
	if !ok {
		t.logger.Error("local cache item with key is not GetListTripsResponse", zap.String("key", key))
		t.localCacheService.Del(ctx, key)
		return nil, fmt.Errorf("local cache item with key %s is not GetListTripsResponse", key)
	}
	return trips, nil
}

// save cache to local cache and redis cache
func (t *transportationQueryService) saveListTripsToRedisCache(ctx context.Context, key string, trips *transportationqueryresponse.GetListTripsResponse) error {
	g, ctx := errgroup.WithContext(ctx)

	// Save to local cache
	g.Go(func() error {
		if ok := t.localCacheService.SetWithTTL(ctx, key, trips, transportationconsts.TRIPS_KEY_LOCAL_TTL); !ok {
			err := fmt.Errorf("local cache set failed")
			t.logger.Error("save local cache failed", zap.Error(err))
			return err
		}
		return nil
	})

	// Save to Redis cache
	g.Go(func() error {
		if err := t.redisCacheService.Set(ctx, key, trips, transportationconsts.TRIPS_KEY_REDIS_TTL); err != nil {
			errWrapped := fmt.Errorf("redis cache set failed: %w", err)
			t.logger.Error("save redis cache failed", zap.Error(errWrapped))
			return errWrapped
		}
		return nil
	})

	// Wait for both
	if err := g.Wait(); err != nil {
		return fmt.Errorf("save cache failed: %w", err)
	}

	return nil
}

// get trips and count from db
func (t *transportationQueryService) getTripsAndCountFromDB(ctx context.Context, request *transportationqueryrequest.GetListTripsRequest) (*transportationqueryresponse.GetListTripsResponse, error) {
	// Parse departure date
	departureDate, err := time.Parse(time.DateOnly, request.DepartureDate)
	if err != nil {
		t.logger.Error("parse departure date failed", zap.Error(err), zap.String("departureDate", request.DepartureDate))
		return nil, err
	}

	var (
		tripsData []transportationmodel.Trip
		count     int
	)

	// Create errgroup with context
	g, ctx := errgroup.WithContext(ctx)

	// Query trips
	g.Go(func() error {
		var err error
		tripsData, err = t.tripRepo.GetListTrips(ctx, departureDate, request.FromLocation, request.ToLocation, request.Page)
		if err != nil {
			t.logger.Error("get list trips failed", zap.Error(err))
		}
		return err
	})

	// Query count
	g.Go(func() error {
		var err error
		count, err = t.tripRepo.GetListTripsCount(ctx, departureDate, request.FromLocation, request.ToLocation)
		if err != nil {
			t.logger.Error("get list trips count failed", zap.Error(err))
		}
		return err
	})

	// Wait for both goroutines
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Map to response
	response := &transportationqueryresponse.GetListTripsResponse{
		Trips: make([]transportationqueryresponse.Trip, len(tripsData)),
		Total: count,
		Page:  request.Page,
	}
	for i, trip := range tripsData {
		response.Trips[i] = *transportationmapper.TripItemToResponse(&trip)
	}

	return response, nil
}

// get trip detail
func (t *transportationQueryService) getTripDetail(ctx context.Context, id uint64) (*transportationqueryresponse.TripDetailResponse, error) {

	// 1. get data from redis cache
	key := fmt.Sprintf(transportationconsts.TRIP_DETAIL_KEY, id)
	dataJson, isFound, err := t.redisCacheService.Get(ctx, key)
	if err != nil {
		t.logger.Error("get trip detail from cache failed", zap.Error(err), zap.String("key", key))
		return nil, err
	}
	if isFound {
		response := &transportationqueryresponse.TripDetailResponse{}
		err = json.Unmarshal([]byte(dataJson), &response)
		if err != nil {
			t.logger.Error("unmarshal trip detail failed", zap.Error(err), zap.String("key", key))
			return nil, err
		}
		return response, nil
	}

	// 2. get data from db
	tripData, err := t.tripRepo.GetTripDetail(ctx, id)
	if err != nil {
		t.logger.Error("get trip detail from db failed", zap.Error(err), zap.String("key", key))
		return nil, err
	}

	// 3. map response
	response := transportationmapper.TripDetailToResponse(tripData)

	// 4. save to redis cache
	err = t.redisCacheService.Set(ctx, key, response, transportationconsts.TRIP_DETAIL_KEY_REDIS_TTL)
	if err != nil {
		t.logger.Error("save trip detail to redis cache failed", zap.Error(err), zap.String("key", key))
	}

	return response, nil
}

// get seats status
func (t *transportationQueryService) getSeatsStatus(ctx context.Context, tripId uint64) ([]transportationqueryresponse.Seat, error) {
	key := fmt.Sprintf(transportationconsts.SEATS_KEY, tripId)
	dataJson, err := t.redisCacheService.HGetAll(ctx, key)
	if err != nil {
		t.logger.Error("get seats status from cache failed", zap.Error(err), zap.String("key", key))
		return nil, err
	}
	if len(dataJson) > 0 {
		response := make([]transportationqueryresponse.Seat, 0, len(dataJson))
		for _, seat := range dataJson {
			seatRes := &transportationqueryresponse.Seat{}
			err = json.Unmarshal([]byte(seat), seatRes)
			if err != nil {
				t.logger.Error("unmarshal seat failed", zap.Error(err), zap.String("key", key))
				return nil, err
			}
			response = append(response, *seatRes)
		}
		return response, nil
	}

	// 2. get data from db
	seatsData, err := t.seatLockRepo.GetListTripSeatLockByTripId(ctx, tripId)
	if err != nil {
		t.logger.Error("get seats by trip id failed", zap.Error(err), zap.String("key", key))
		return nil, err
	}

	// 3. map response
	response := make([]transportationqueryresponse.Seat, 0)
	for _, seat := range seatsData {
		response = append(response, transportationmapper.SeatLockModelToSeatStatusItem(seat))
	}

	// 4. save to redis cache
	err = t.saveSeatsToRedisCache(ctx, key, response)
	if err != nil {
		t.logger.Error("save seats status to redis cache failed", zap.Error(err), zap.String("key", key))
	}

	return response, nil
}

// save seats to redis cache
func (t *transportationQueryService) saveSeatsToRedisCache(ctx context.Context, key string, seats []transportationqueryresponse.Seat) error {
	dataJSON, err := json.Marshal(seats)
	if err != nil {
		return err
	}
	script, err := utils.LoadLuaScript("save_cache_map_seat.lua")
	if err != nil {
		return err
	}
	result, err := t.redisCacheService.Eval(ctx, script, []string{key}, dataJSON, transportationconsts.SEATS_KEY_REDIS_TTL_SECONDS)
	if err != nil {

		return err
	}
	if result != 1 {
		return fmt.Errorf("save seats status to redis cache failed")
	}

	return nil
}

// get seat detail from local cache
func (t *transportationQueryService) getTripDetailFromLocalCache(ctx context.Context, key string) (*transportationqueryresponse.TripDetailResponse, error) {
	value, isFound := t.localCacheService.Get(ctx, key)
	if !isFound {
		return nil, nil
	}
	detail, ok := value.(*transportationqueryresponse.TripDetailResponse)
	if !ok {
		t.logger.Error("local cache item with key is not TripDetailResponse", zap.String("key", key))
		t.localCacheService.Del(ctx, key)
		return nil, fmt.Errorf("local cache item with key %s is not TripDetailResponse", key)
	}
	return detail, nil
}
