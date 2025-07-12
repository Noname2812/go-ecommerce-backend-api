package cacheservice

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dgraph-io/ristretto"
)

type localCache struct {
	cache *ristretto.Cache
}

// implementation localcache

func NewlocalCache() LocalCache {
	// ref here ANH EM: https://github.com/hypermodeinc/ristretto?tab=readme-ov-file#usage
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic("failed to create ristretto cache")
	}
	return &localCache{cache: cache}
}

func (rc *localCache) Get(ctx context.Context, key string) (interface{}, bool) {
	return rc.cache.Get(key)
}

func (rc *localCache) Set(ctx context.Context, key string, value interface{}) bool {
	return rc.cache.Set(key, value, 1) // Cost mặc định = 1
}

func (rc *localCache) SetWithTTL(ctx context.Context, key string, value interface{}) bool {
	dataJson, _ := json.Marshal(value)
	return rc.cache.SetWithTTL(key, string(dataJson), 1, 5*time.Minute) // Cost mặc định = 1
}

func (rc *localCache) Del(ctx context.Context, key string) {
	rc.cache.Del(key)
}
