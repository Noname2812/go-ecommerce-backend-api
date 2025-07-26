package cacheservice

import (
	"context"
	"time"
)

// RedisCache is an interface for caching operations
type RedisCache interface {
	// Get a value from the cache
	Get(ctx context.Context, key string) (string, bool, error)
	// Set a value in the cache
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	// Delete a value from the cache
	Del(ctx context.Context, key string) error
	// Increment a value in the cache
	Incr(ctx context.Context, key string) (int64, error)
	// Decrement a value in the cache
	Decr(ctx context.Context, key string) (int64, error)
	// Check if a key exists in the cache
	Exists(ctx context.Context, key string) (bool, error)
	// Set the expiration time for a key
	Expire(ctx context.Context, key string, expiration time.Duration) error
	// WithDistributedLock is a helper function to execute a function with a distributed lock
	WithDistributedLock(ctx context.Context, key string, ttl time.Duration, fn func(ctx context.Context) (interface{}, error)) (interface{}, error)
	// Get all fields in a hash
	HGetAll(ctx context.Context, key string) (map[string]string, error)
	// Set a field in a hash
	HSet(ctx context.Context, key string, field string, value interface{}) error
	// Delete a field from a hash
	HDel(ctx context.Context, key string, field string) error
	// Atomically set field if not exists (useful for holding a seat)
	HSetNX(ctx context.Context, key string, field string, value interface{}) (bool, error)
	// Run Lua script for atomic operations (e.g., hold seat with TTL or cleanup expired holds)
	Eval(ctx context.Context, script string, keys []string, args ...interface{}) (interface{}, error)
}
