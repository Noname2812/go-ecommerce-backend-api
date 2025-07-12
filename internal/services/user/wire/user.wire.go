//go:build wireinject

package userwire

import (
	"database/sql"

	usercommandhandler "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/command/handler"
	userqueryhandler "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/query/handler"
	useremessagingimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/infrastructure/messaging"
	userinforepositoryimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/infrastructure/persistence/userinfo"
	userserviceimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/infrastructure/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/google/wire"
)

var UserRepositorySet = wire.NewSet(
	userinforepositoryimpl.NewUserInfoRepository,
)

func InitUserQueryHandler(db *sql.DB, logger *zap.Logger) userqueryhandler.UserQueryHandler {
	wire.Build(
		UserRepositorySet,
		userserviceimpl.NewUserQueryService,
		userqueryhandler.NewUserQueryHandler,
	)
	return nil
}

func InitUserCommandHandler(db *sql.DB, logger *zap.Logger, manager *kafka.Manager, redisClient *redis.Client) usercommandhandler.UserCommandHandler {
	wire.Build(
		UserRepositorySet,
		userserviceimpl.NewUserCacheService,
		useremessagingimpl.NewuserEventPublisher,
		userserviceimpl.NewUserCommandService,
		usercommandhandler.NewUserCommandHandler,
	)
	return nil
}
