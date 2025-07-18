// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package authwire

import (
	"database/sql"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/handler"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/persistence/userbase"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/service"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/service/messaging"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// Injectors from auth.wire.go:

func InitAuthHttpCommandHandler(db *sql.DB, rdb *redis.Client, logger *zap.Logger, kafkaManager *kafka.Manager) authcommandhandler.AuthCommandHttpHandler {
	userBaseRepository := userbaserepositoryimpl.NewUserBaseRepository(db)
	authCacheService := authserviceimpl.NewAuthCacheService(rdb)
	authEventPublisher := authmessagingserviceimpl.NewAuthEventPublisher(kafkaManager, logger)
	authCommandService := authserviceimpl.NewAuthCommandService(logger, userBaseRepository, authCacheService, authEventPublisher)
	authCommandHttpHandler := authcommandhandler.NewAuthCommandHttpHandler(authCommandService, logger)
	return authCommandHttpHandler
}

// auth.wire.go:

var authRepositorySet = wire.NewSet(userbaserepositoryimpl.NewUserBaseRepository)
