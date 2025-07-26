package cacheservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

type sRedisCache struct {
	// client *redis.Client // Chỉ cần client, không cần redsync.
	client *redis.Client
	locker *redislock.Client
}

// Eval implements RedisCache.
// Run a Lua script in Redis for atomic operations
func (s *sRedisCache) Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error) {
	result, err := s.client.Eval(ctx, script, keys, args...).Result()
	if err != nil {
		return nil, fmt.Errorf("redis eval error: %w", err)
	}
	return result, nil
}

// HDel implements RedisCache.
// Delete a field in a hash
func (s *sRedisCache) HDel(ctx context.Context, key string, field string) error {
	if err := s.client.HDel(ctx, key, field).Err(); err != nil {
		return fmt.Errorf("redis hdel error: %w", err)
	}
	return nil
}

// HGetAll implements RedisCache.
// Get all fields in a hash
func (s *sRedisCache) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	result, err := s.client.HGetAll(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return map[string]string{}, nil // return empty map if not exist
		}
		return nil, fmt.Errorf("redis hgetall error: %w", err)
	}
	return result, nil
}

// HSet implements RedisCache.
// Set a field in a hash
func (s *sRedisCache) HSet(ctx context.Context, key string, field string, value interface{}) error {
	val, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json marshal error: %w", err)
	}
	if err := s.client.HSet(ctx, key, field, val).Err(); err != nil {
		return fmt.Errorf("redis hset error: %w", err)
	}
	return nil
}

// HSetNX implements RedisCache.
// Set a field in a hash only if it does not exist
func (s *sRedisCache) HSetNX(ctx context.Context, key string, field string, value interface{}) (bool, error) {
	val, err := json.Marshal(value)
	if err != nil {
		return false, fmt.Errorf("json marshal error: %w", err)
	}
	success, err := s.client.HSetNX(ctx, key, field, val).Result()
	if err != nil {
		return false, fmt.Errorf("redis hsetnx error: %w", err)
	}
	return success, nil
}

// Expire implements RedisCache.
func (s *sRedisCache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	if err := s.client.Expire(ctx, key, expiration).Err(); err != nil {
		return fmt.Errorf("redis expire error: %w", err)
	}
	return nil
}

func NewRedisCache(client *redis.Client) RedisCache {
	return &sRedisCache{
		client: client,
		locker: redislock.New(client),
	}
}

func (s *sRedisCache) Get(ctx context.Context, key string) (string, bool, error) {
	val, err := s.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return val, false, nil //val = ""
		}
		return val, false, fmt.Errorf("redis get error: %w", err)
	}

	return val, true, nil // string json
}

func (s *sRedisCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	b, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("json marshal error: %w", err)
	}
	if err := s.client.Set(ctx, key, b, expiration).Err(); err != nil {
		return fmt.Errorf("redis set error: %w", err)
	}
	return nil
}

func (s *sRedisCache) Del(ctx context.Context, key string) error {
	if err := s.client.Del(ctx, key).Err(); err != nil {
		return fmt.Errorf("redis del error: %w", err)
	}
	return nil
}

func (s *sRedisCache) Incr(ctx context.Context, key string) (int64, error) {
	val, err := s.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("redis incr error: %w", err)
	}
	return val, nil
}

func (s *sRedisCache) Decr(ctx context.Context, key string) (int64, error) {
	val, err := s.client.Decr(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("redis decr error: %w", err)
	}
	return val, nil
}

func (s *sRedisCache) Exists(ctx context.Context, key string) (bool, error) {
	val, err := s.client.Exists(ctx, key).Result()
	if err != nil {
		return false, fmt.Errorf("redis exists error: %w", err)
	}
	return val == 1, nil
}

func (s *sRedisCache) WithDistributedLock(ctx context.Context, key string, ttl time.Duration, fn func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	lock, err := s.locker.Obtain(ctx, key, ttl, nil)
	if err == redislock.ErrNotObtained {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to obtain lock: %w", err)
	}
	defer lock.Release(ctx)

	result, err := fn(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
