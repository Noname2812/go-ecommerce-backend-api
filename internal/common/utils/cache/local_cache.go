package cacheservice

import "context"

type LocalCache interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, value interface{}) bool
	SetWithTTL(ctx context.Context, key string, value interface{}) bool
	Del(ctx context.Context, key string)
}
