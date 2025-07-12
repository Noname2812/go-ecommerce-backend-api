//go:build wireinject

package notificationwire

import (
	notificationmessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/application/messaging"
	notificationmessagingimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/infrastructure/messaging"
	notificationserviceimpl "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/infrastructure/service"
	kafkaManager "github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/setting"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func InitNotificationConsumer(kafkaManager *kafkaManager.Manager, logger *zap.Logger, config *setting.EmailSetting) notificationmessaging.NotificationConsumer {
	wire.Build(
		notificationserviceimpl.NewNotificationService,
		notificationmessagingimpl.NewNotificationConsumer,
	)
	return nil
}
