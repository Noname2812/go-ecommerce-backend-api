package redis

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
				PoolSize: 10,
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

// advanced
func InitRedisSentinel() {
	// rdb := redis.NewFailoverClient(&redis.FailoverOptions{
	// 	MasterName:    "mymaster", // Tên master do Sentinel quản lý
	// 	SentinelAddrs: []string{"127.0.0.1:26379", "127.0.0.1:26380", "127.0.0.1:26381"},
	// 	DB:            0,        // Sử dụng database mặc định
	// 	Password:      "123456", // Nếu Redis có mật khẩu, điền vào đây
	// })

	// // Check the connection
	// _, err := rdb.Ping(ctx).Result()
	// if err != nil {
	// 	log.Fatalf("Failed to connect to Redis Sentinel: %v", err)
	// }

	// fmt.Println("Connected to Redis Sentinel successfully!")

	// // Try setting and getting a value
	// err = rdb.Set(ctx, "test_key", "Hello Redis Sentinel!", 0).Err()
	// if err != nil {
	// 	log.Fatalf("Error setting key: %v", err)
	// }

	// val, err := rdb.Get(ctx, "test_key").Result()
	// if err != nil {
	// 	log.Fatalf("Error getting key: %v", err)
	// }

	// fmt.Println("Value of test_key:", val)

	// global.Logger.Info("Initializing RedisSentinel Successfully")
	// global.Rdb = rdb
	// redisExample()
}
