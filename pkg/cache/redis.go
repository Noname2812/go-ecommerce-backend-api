package cache

import (
	"context"
	"fmt"
	"runtime/debug"
	"sync"
	"time"

	"github.com/Noname2812/go-ecommerce-backend-api/pkg/logger"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

var (
	redisRetryCount = 0        // count retry connect
	maxRetries      = 3        // max count retry connect
	redisMutex      sync.Mutex // Mutex lock
)

// NewRedis creates a new Redis client
func NewRedis(config setting.RedisSetting, logger *logger.LoggerZap) *redis.Client {
	var rdb *redis.Client
	for redisRetryCount = 0; redisRetryCount <= maxRetries; redisRetryCount++ {
		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Error("Recovered from Redis panic",
						zap.Any("error", r),
						zap.Int("retry_count", redisRetryCount),
						zap.Int("maxRetries", maxRetries),
						zap.String("stack", string(debug.Stack())),
					)
				}
			}()

			rdb = redis.NewClient(&redis.Options{
				Addr:     fmt.Sprintf("%s:%v", config.Host, config.Port),
				Password: config.Password,
				DB:       config.Database,
				PoolSize: 50, // Increased from 10 to handle high concurrency
			})

			_, err := rdb.Ping(ctx).Result()
			if err != nil {
				panic(err)
			}
			logger.Info("Initializing Redis Successfully")
			redisRetryCount = 0
		}()

		if rdb != nil {
			break
		}

		if redisRetryCount < maxRetries {
			backoff := time.Duration((redisRetryCount+1)*(redisRetryCount+1)) * time.Second
			fmt.Println(">>>>>>backoff: ", backoff)
			logger.Warn("Retrying Redis connection...",
				zap.Int("attempt", redisRetryCount+1),
				zap.Duration("backoff", backoff),
			)
			time.Sleep(backoff)
		} else {
			panic("Redis connection failed after max retries")
		}
	}
	return rdb
}

// advanced redis sentinel
func InitRedisSentinel(config setting.RedisSentinelSetting, logger *logger.LoggerZap) *redis.Client {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    config.MasterName, // Tên master do Sentinel quản lý
		SentinelAddrs: config.SentinelAddrs,
		DB:            config.Database, // Sử dụng database mặc định
		Password:      config.Password, // Nếu Redis có mật khẩu, điền vào đây
	})

	// Check the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logger.Error("Failed to connect to Redis Sentinel", zap.Error(err))
	}

	logger.Info("Connected to Redis Sentinel successfully!")
	return rdb
}
