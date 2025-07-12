package notificationmessaging

import (
	notificationmessaginghandler "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/application/messaging/handler"
)

type NotificationPublisher interface {
	Register() // register all producers of service
	// handlers
	notificationmessaginghandler.NotificationPublisherHandler
}
