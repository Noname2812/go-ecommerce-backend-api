package initialize

import (
	"context"

	authmessagingimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/messaging"
	notificationwire "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/wire"
	userwire "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/wire"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
)

func InitKafka(ctx context.Context, config *setting.Config, container *AppContainer) {
	// ---------------- Producers ---------------- //

	// auth
	authPublisher := authmessagingimpl.NewAuthPublisher(container.KafkaManager, container.Logger)
	authPublisher.Register()

	// ---------------- Consumers ---------------- //
	// notification
	notificationConsumer := notificationwire.InitNotificationConsumer(container.KafkaManager, container.Logger, &config.Email)
	if err := notificationConsumer.Subscribe(); err != nil {
		container.Logger.Fatal("failed to subscribe notification consumers", zap.Error(err))
	}
	// user
	userConsumer := userwire.InitUserConsumer(container.DB, container.Logger, container.KafkaManager, container.RedisClient)
	if err := userConsumer.Subscribe(); err != nil {
		container.Logger.Fatal("failed to subscribe notification consumers", zap.Error(err))
	}
	container.KafkaManager.StartAllConsumers(ctx)
}
