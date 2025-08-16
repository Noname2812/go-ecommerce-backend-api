//go:build wireinject

package userwire

import (
	"database/sql"

	userpb "github.com/Noname2812/go-ecommerce-backend-api/internal/common/protogen/user"
	usercommandhandler "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/command/handler"
	usermessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/messaging"
	userqueryhandler "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/query/handler"
	useremessagingimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/infrastructure/messaging"
	userrepositoryimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/infrastructure/persistence"
	userserviceimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/infrastructure/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/google/wire"
)

var UserRepositorySet = wire.NewSet(
	userrepositoryimpl.NewUserInfoRepository,
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
		useremessagingimpl.NewUserPublisher,
		userserviceimpl.NewUserCommandService,
		usercommandhandler.NewUserCommandHandler,
	)
	return nil
}

func InitUserServiceServer(db *sql.DB, logger *zap.Logger) userpb.UserServiceServer {
	wire.Build(
		UserRepositorySet,
		userserviceimpl.NewUserServiceServer,
	)
	return nil
}

func InitUserConsumer(db *sql.DB, logger *zap.Logger, manager *kafka.Manager, redisClient *redis.Client) usermessaging.UserConsumer {
	wire.Build(
		UserRepositorySet,
		useremessagingimpl.NewUserPublisher,
		userserviceimpl.NewUserCommandService,
		useremessagingimpl.NewUserConsumer,
	)
	return nil
}
