package global

import (
	"database/sql"

	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/logger"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var (
	Config       setting.Config
	Logger       *logger.LoggerZap
	Rdb          *redis.Client
	Mdbc         *sql.DB
	KafkaManager *kafka.Manager
)
