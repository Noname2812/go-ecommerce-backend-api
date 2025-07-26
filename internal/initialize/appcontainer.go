package initialize

import (
	"database/sql"

	"github.com/Noname2812/go-ecommerce-backend-api/global"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/cache"
	redispkg "github.com/Noname2812/go-ecommerce-backend-api/pkg/cache"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/db"
	grpcserver "github.com/Noname2812/go-ecommerce-backend-api/pkg/grpc"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/logger"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"github.com/dgraph-io/ristretto"
	"github.com/redis/go-redis/v9"

	"go.uber.org/zap"
)

type AppContainer struct {
	DB                *sql.DB
	RedisClient       *redis.Client
	KafkaManager      *kafka.Manager
	Logger            *zap.Logger
	GRPCServerManager *grpcserver.GRPCServerManager
	LocalCache        *ristretto.Cache
}

func NewAppContainer(config *setting.Config) (*AppContainer, error) {

	log := logger.NewLogger(config.Logger)
	db := db.NewMySqlC(config.Mysql, log)
	redisClient := redispkg.NewRedis(config.Redis, log)
	kafkaManager := kafka.NewManager(config.Kafka.Brokers, log.Logger)
	GRPCServerManager := grpcserver.NewGRPCServerManager()
	localCache, err := cache.NewRistrettoCache()
	if err != nil {
		return nil, err
	}

	// save global
	global.KafkaManager = kafkaManager
	global.Logger = log
	global.Mdbc = db
	global.Rdb = redisClient

	return &AppContainer{
		DB:                db,
		RedisClient:       redisClient,
		KafkaManager:      kafkaManager,
		Logger:            log.Logger,
		GRPCServerManager: GRPCServerManager,
		LocalCache:        localCache,
	}, nil
}
