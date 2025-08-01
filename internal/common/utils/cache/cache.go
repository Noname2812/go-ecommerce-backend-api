package cacheservice

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Noname2812/go-ecommerce-backend-api/global"
	"github.com/redis/go-redis/v9"
)

func GetCache(ctx context.Context, key string, obj interface{}) error {
	rs, err := global.Rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("key %s not found", key)
	} else if err != nil {
		return err
	}
	// convert rs json to object
	if err := json.Unmarshal([]byte(rs), obj); err != nil {
		return fmt.Errorf("failed to unmarshal")
	}
	return nil
}
