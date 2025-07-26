package transportationserviceimpl

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	cacheservice "github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/cache"
	transportationqueryrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/request"
	transportationqueryresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/response"
	transportationservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/service"
	transportationconsts "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/consts"
	transportationrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/repository"
	transportationmapper "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/infrastructure/mapper"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

type transportationQueryService struct {
	logger             *zap.Logger
	transportationRepo transportationrepository.TripRepository
	redisCacheService  cacheservice.RedisCache
	localCacheService  cacheservice.LocalCache
}

// GetListTrips implements transportationservice.TransportationQueryService.
func (t *transportationQueryService) GetListTrips(ctx context.Context, request *transportationqueryrequest.GetListTripsRequest) (int, *transportationqueryresponse.GetListTripsResponse, error) {
	key := fmt.Sprintf("%s:%s-%s:%s:%d", transportationconsts.TRIPS_KEY_PREFIX, request.FromLocation, request.ToLocation, request.DepartureDate, request.Page)

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
			_ = t.localCacheService.SetWithTTL(context.Background(), key, trips, transportationconsts.TRIPS_KEY_LOCAL_TTL)
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
		departureDate, err := time.Parse(time.DateOnly, request.DepartureDate)
		if err != nil {
			t.logger.Error("parse departure date failed", zap.Error(err), zap.String("departureDate", request.DepartureDate))
			return nil, err
		}

		data, err := t.transportationRepo.GetListTrips(ctx, departureDate, request.FromLocation, request.ToLocation, request.Page)
		if err != nil {
			t.logger.Error("get list trips failed", zap.Error(err), zap.String("departureDate", request.DepartureDate), zap.String("fromLocation", request.FromLocation), zap.String("toLocation", request.ToLocation))
			return nil, err
		}
		count, err := t.transportationRepo.GetListTripsCount(ctx, departureDate, request.FromLocation, request.ToLocation)
		if err != nil {
			t.logger.Error("get list trips count failed", zap.Error(err), zap.String("departureDate", request.DepartureDate), zap.String("fromLocation", request.FromLocation), zap.String("toLocation", request.ToLocation))
			return nil, err
		}

		response := &transportationqueryresponse.GetListTripsResponse{
			Trips: make([]transportationqueryresponse.Trip, len(data)),
			Total: count,
			Page:  request.Page,
		}
		for i, trip := range data {
			response.Trips[i] = transportationmapper.TripToResponse(trip)
		}

		// save the data to the cache
		t.saveCache(key, response)
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
	cache, err := t.retryWithExponentialBackoff(ctx, transportationconsts.MAX_RETRY_GET_LIST_TRIPS, transportationconsts.RETRY_GET_LIST_TRIPS_BACKOFF, func() (*transportationqueryresponse.GetListTripsResponse, error) {
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

	return response.ErrCodeSuccess, cache, nil
}

func NewTransportationQueryService(transportationRepo transportationrepository.TripRepository,
	cacheService cacheservice.RedisCache,
	localCacheService cacheservice.LocalCache,
	logger *zap.Logger,
) transportationservice.TransportationQueryService {
	return &transportationQueryService{
		transportationRepo: transportationRepo,
		redisCacheService:  cacheService,
		localCacheService:  localCacheService,
		logger:             logger,
	}
}

// retry with exponential backoff
func (t *transportationQueryService) retryWithExponentialBackoff(ctx context.Context, retries int, baseDelay time.Duration, getCache func() (*transportationqueryresponse.GetListTripsResponse, error)) (*transportationqueryresponse.GetListTripsResponse, error) {
	for i := 0; i < retries; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		cache, err := getCache()
		if err != nil {
			return nil, err
		}
		if cache != nil {
			return cache, nil
		}

		if i == retries-1 {
			break
		}

		// exponential backoff with jitter
		jitter := time.Duration(rand.Int63n(int64(baseDelay)))
		backoffDelay := baseDelay*time.Duration(1<<i) + jitter
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(backoffDelay):
		}
	}
	return nil, fmt.Errorf("exceeded max retries: %d", retries)
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
func (t *transportationQueryService) saveCache(key string, trips *transportationqueryresponse.GetListTripsResponse) {
	go func() {
		saveCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = t.localCacheService.SetWithTTL(saveCtx, key, trips, transportationconsts.TRIPS_KEY_LOCAL_TTL)
	}()
	go func() {
		saveCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = t.redisCacheService.Set(saveCtx, key, trips, transportationconsts.TRIPS_KEY_REDIS_TTL)
	}()
}
