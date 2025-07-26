//go:build wireinject

package authwire

import (
	"database/sql"

	cacheservice "github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/cache"
	authcommandhandler "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/handler"
	authclientgrpc "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/client"
	authmessagingimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/messaging"
	authrepositoryimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/persistence"
	authserviceimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/service"
	grpcserver "github.com/Noname2812/go-ecommerce-backend-api/pkg/grpc"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"github.com/google/wire"
)

var authRepositorySet = wire.NewSet(
	authrepositoryimpl.NewUserBaseRepository,
	authrepositoryimpl.NewTransactionManager,
)

func InitAuthHttpCommandHandler(db *sql.DB, rdb *redis.Client, logger *zap.Logger, kafkaManager *kafka.Manager, manager *grpcserver.GRPCServerManager) authcommandhandler.AuthCommandHttpHandler {
	wire.Build(
		authRepositorySet,
		authclientgrpc.NewUserGRPCClient,
		authserviceimpl.NewAuthCommandService,
		cacheservice.NewRedisCache,
		authmessagingimpl.NewAuthPublisher,
		authcommandhandler.NewAuthCommandHttpHandler,
	)
	return nil
}
