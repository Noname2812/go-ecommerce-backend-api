package initialize

import (
	"context"

	authmessagingserviceimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/service/messaging"
	notificationwire "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/wire"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"go.uber.org/zap"
)

func InitKafka(ctx context.Context, config *setting.Config, container *AppContainer) {
	// ---------------- Producers ---------------- //

	// auth
	authPublisher := authmessagingserviceimpl.NewAuthEventPublisher(container.KafkaManager, container.Logger)
	authPublisher.Register()

	// ---------------- Consumers ---------------- //
	// notification
	notificationConsumer := notificationwire.InitNotificationConsumer(container.KafkaManager, container.Logger, &config.Email)
	if err := notificationConsumer.Subscribe(); err != nil {
		container.Logger.Fatal("failed to subscribe notification consumers", zap.Error(err))
	}

	// start all
	container.KafkaManager.StartAllConsumers(ctx)
}
