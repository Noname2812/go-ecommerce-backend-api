package notificationmessaging

import notificationmessaginghandler "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/application/messaging/handler"

type NotificationConsumer interface {
	Subscribe() error // Subscribe all topic of service
	// handler
	notificationmessaginghandler.NotificationConsumerHandler
}
