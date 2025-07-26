package cache

import (
	"errors"

	"github.com/dgraph-io/ristretto"
)

// implementation localcache

func NewRistrettoCache() (*ristretto.Cache, error) {
	// ref here : https://github.com/hypermodeinc/ristretto?tab=readme-ov-file#usage
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		return nil, errors.New("failed to create ristretto cache")
	}
	return cache, nil
}
