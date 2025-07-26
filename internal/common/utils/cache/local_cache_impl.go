package cacheservice

import (
	"context"
	"time"

	"github.com/dgraph-io/ristretto"
)

type localCache struct {
	cache *ristretto.Cache
}

// implementation localcache

func NewLocalCache(cache *ristretto.Cache) LocalCache {
	return &localCache{cache: cache}
}

func (rc *localCache) Get(ctx context.Context, key string) (interface{}, bool) {
	return rc.cache.Get(key)
}

func (rc *localCache) Set(ctx context.Context, key string, value interface{}) bool {
	return rc.cache.Set(key, value, 1) // Cost mặc định = 1
}

func (rc *localCache) SetWithTTL(ctx context.Context, key string, value interface{}, ttl time.Duration) bool {
	return rc.cache.SetWithTTL(key, value, 1, ttl) // Cost mặc định = 1
}

func (rc *localCache) Del(ctx context.Context, key string) {
	rc.cache.Del(key)
}
