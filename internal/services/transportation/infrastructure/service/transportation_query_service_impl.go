package transportationserviceimpl

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
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
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type transportationQueryService struct {
	logger            *zap.Logger
	tripRepo          transportationrepository.TripRepository
	seatRepo          transportationrepository.SeatRepository
	seatLockRepo      transportationrepository.TripSeatLockRepository
	redisCacheService cacheservice.RedisCache
	localCacheService cacheservice.LocalCache
}

// GetTripDetail implements transportationservice.TransportationQueryService.
func (t *transportationQueryService) GetTripDetail(ctx context.Context, id uint64) (code int, data *transportationqueryresponse.TripDetailResponse, err error) {
	if id == 0 {
		return response.ErrCodeParamInvalid, nil, fmt.Errorf("invalid trip id")
	}
	key := fmt.Sprintf(transportationconsts.TRIP_DETAIL_KEY, id)
	// 1. get trip detail from local cache
	value, isFound := t.localCacheService.Get(ctx, key)
	if isFound {
		res, ok := value.(*transportationqueryresponse.TripDetailResponse)
		if !ok {
			t.logger.Error("local cache item with key is not TripDetailResponse", zap.String("key", key))
			t.localCacheService.Del(ctx, key)
			return response.ErrServerError, nil, fmt.Errorf("local cache item with key %s is not TripDetailResponse", key)
		}
		return response.ErrCodeSuccess, res, nil
	}
	// 2. get trip detail from redis cache or db
	lockKey := fmt.Sprintf("lock:%s", key)
	result, err := t.redisCacheService.WithDistributedLock(ctx, lockKey, transportationconsts.TRIPS_LOCK_TTL_SECONDS, func(ctx context.Context) (interface{}, error) {
		// Double-check cache after acquiring lock
		value, isFound := t.localCacheService.Get(ctx, key)
		if isFound {
			return value, nil // Data already cached by another process
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
	// If we couldn't acquire the lock (another process is handling it), retry to get from cache
	if err != nil {
		return response.ErrServerError, nil, err
	}
	// If we got the lock and processed the request
	if result != nil {
		return response.ErrCodeSuccess, result.(*transportationqueryresponse.TripDetailResponse), nil
	}
	// 4. Retry to get data from cache with exponential backoff
	cache, err := utils.RetryWithExponentialBackoff(ctx, transportationconsts.MAX_RETRY_GET_TRIP_DETAIL, transportationconsts.RETRY_GET_TRIP_DETAIL_BACKOFF, func() (interface{}, error) {
		// Check local cache
		value, isFound := t.localCacheService.Get(ctx, key)
		if isFound {
			res, ok := value.(*transportationqueryresponse.TripDetailResponse)
			if ok {
				return res, nil
			}
			return nil, fmt.Errorf("local cache item with key %s is not TripDetailResponse", key)
		}
		return nil, nil
	})

	if err != nil {
		t.logger.Error("retry to get from cache failed", zap.Error(err), zap.String("key", key))
		return response.ErrServerError, nil, err
	}

	return response.ErrCodeSuccess, cache.(*transportationqueryresponse.TripDetailResponse), nil
}

// GetListTrips implements transportationservice.TransportationQueryService.
func (t *transportationQueryService) GetListTrips(ctx context.Context, request *transportationqueryrequest.GetListTripsRequest) (int, *transportationqueryresponse.GetListTripsResponse, error) {
	key := fmt.Sprintf(transportationconsts.TRIPS_KEY, request.FromLocation, request.ToLocation, request.DepartureDate, request.Page)

	// 1. get data from local cache first
	value, isFound := t.localCacheService.Get(ctx, key)
	if isFound {
		tripsLocal, ok := value.(*transportationqueryresponse.GetListTripsResponse)
		if !ok {
			t.logger.Error("local cache item with key is not GetListTripsResponse", zap.String("key", key))
			t.localCacheService.Del(ctx, key)
			return response.ErrServerError, nil, fmt.Errorf("local cache item with key %s is not GetListTripsResponse", key)
		}
		return response.ErrCodeSuccess, tripsLocal, nil
	}

	// 2. check if data exists in redis cache
	trips, err := t.getListTripsFromRedisCache(ctx, key)
	if err != nil {
		t.logger.Error("get list trips from redis cache failed", zap.Error(err), zap.String("key", key))
		return response.ErrServerError, nil, err
	}
	if trips != nil {
		// Cache warming: save to local cache for next requests
		go func() {
			_ = t.localCacheService.SetWithTTL(ctx, key, trips, transportationconsts.TRIPS_KEY_LOCAL_TTL)
		}()
		return response.ErrCodeSuccess, trips, nil
	}

	// 3. Data not in cache, try to acquire distributed lock
	lockKey := fmt.Sprintf("lock:%s", key)
	result, err := t.redisCacheService.WithDistributedLock(ctx, lockKey, transportationconsts.TRIPS_LOCK_TTL_SECONDS, func(ctx context.Context) (interface{}, error) {
		// Double-check cache after acquiring lock
		exists, err := t.redisCacheService.Exists(ctx, key)
		if err == nil && exists {
			return nil, nil // Data already cached by another process
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
	// If we couldn't acquire the lock (another process is handling it), retry to get from cache
	if err != nil {
		return response.ErrServerError, nil, err
	}
	// If we got the lock and processed the request
	if result != nil {
		return response.ErrCodeSuccess, result.(*transportationqueryresponse.GetListTripsResponse), nil
	}
	// 4. Retry to get data from cache with exponential backoff
	cache, err := utils.RetryWithExponentialBackoff(ctx, transportationconsts.MAX_RETRY_GET_LIST_TRIPS, transportationconsts.RETRY_GET_LIST_TRIPS_BACKOFF, func() (interface{}, error) {
		// Check local cache first
		value, isFound := t.localCacheService.Get(ctx, key)
		if isFound {
			trips, ok := value.(*transportationqueryresponse.GetListTripsResponse)
			if ok {
				return trips, nil
			}
			t.logger.Error("local cache item with key is not GetListTripsResponse", zap.String("key", key))
			t.localCacheService.Del(ctx, key)
			return nil, fmt.Errorf("local cache item with key %s is not GetListTripsResponse", key)
		}

		// Check redis cache
		return t.getListTripsFromRedisCache(ctx, key)
	})

	if err != nil {
		t.logger.Error("retry to get from cache failed", zap.Error(err), zap.String("key", key))
		return response.ErrServerError, nil, err
	}

	return response.ErrCodeSuccess, cache.(*transportationqueryresponse.GetListTripsResponse), nil
}

func NewTransportationQueryService(tripRepo transportationrepository.TripRepository,
	seatRepo transportationrepository.SeatRepository,
	seatLockRepo transportationrepository.TripSeatLockRepository,
	cacheService cacheservice.RedisCache,
	localCacheService cacheservice.LocalCache,
	logger *zap.Logger,
) transportationservice.TransportationQueryService {
	return &transportationQueryService{
		tripRepo:          tripRepo,
		seatRepo:          seatRepo,
		seatLockRepo:      seatLockRepo,
		redisCacheService: cacheService,
		localCacheService: localCacheService,
		logger:            logger,
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

// save cache to local cache and redis cache
func (t *transportationQueryService) saveListTripsToRedisCache(ctx context.Context, key string, trips *transportationqueryresponse.GetListTripsResponse) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 2)

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := t.localCacheService.SetWithTTL(ctx, key, trips, transportationconsts.TRIPS_KEY_LOCAL_TTL); !err {
			errCh <- fmt.Errorf("local cache set failed")
		}
	}()

	go func() {
		defer wg.Done()
		if err := t.redisCacheService.Set(ctx, key, trips, transportationconsts.TRIPS_KEY_REDIS_TTL); err != nil {
			errCh <- fmt.Errorf("redis cache set failed: %w", err)
		}
	}()

	wg.Wait()
	close(errCh)

	for err := range errCh {
		t.logger.Error("save cache failed", zap.Error(err))
	}

	if len(errCh) > 0 {
		return fmt.Errorf("save cache failed with errors: %v", errCh)
	}
	return nil
}

// get trips and count from db
func (t *transportationQueryService) getTripsAndCountFromDB(ctx context.Context, request *transportationqueryrequest.GetListTripsRequest) (*transportationqueryresponse.GetListTripsResponse, error) {
	departureDate, err := time.Parse(time.DateOnly, request.DepartureDate)
	if err != nil {
		t.logger.Error("parse departure date failed", zap.Error(err), zap.String("departureDate", request.DepartureDate))
		return nil, err
	}

	var (
		tripsData []transportationmodel.Trip
		count     int
		queryErr  error
		countErr  error
		wg        sync.WaitGroup
	)

	wg.Add(2)

	// Query trips
	go func() {
		defer wg.Done()
		tripsData, queryErr = t.tripRepo.GetListTrips(ctx, departureDate, request.FromLocation, request.ToLocation, request.Page)
	}()

	// Query count
	go func() {
		defer wg.Done()
		count, countErr = t.tripRepo.GetListTripsCount(ctx, departureDate, request.FromLocation, request.ToLocation)
	}()

	wg.Wait()

	if queryErr != nil {
		t.logger.Error("get list trips failed", zap.Error(queryErr))
		return nil, queryErr
	}
	if countErr != nil {
		t.logger.Error("get list trips count failed", zap.Error(countErr))
		return nil, countErr
	}

	// map response
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
