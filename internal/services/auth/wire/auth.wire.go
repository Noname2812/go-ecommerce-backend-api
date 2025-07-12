//go:build wireinject

package authwire

import (
	"database/sql"

	cacheservice "github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/cache"
	authcommandhandler "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/handler"
	userbaserepositoryimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/persistence/userbase"
	authserviceimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/service"
	authmessagingserviceimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/service/messaging"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/google/wire"
)

var authRepositorySet = wire.NewSet(
	userbaserepositoryimpl.NewUserBaseRepository,
)

func InitAuthHttpCommandHandler(db *sql.DB, rdb *redis.Client, logger *zap.Logger, kafkaManager *kafka.Manager) authcommandhandler.AuthCommandHttpHandler {
	wire.Build(
		authRepositorySet,
		authserviceimpl.NewAuthCommandService,
		cacheservice.NewRedisCache,
		authmessagingserviceimpl.NewAuthEventPublisher,
		authcommandhandler.NewAuthCommandHttpHandler,
	)
	return nil
}
