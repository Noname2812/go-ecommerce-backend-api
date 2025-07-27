package notificationmessagingimpl

import (
	notificationmessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/application/messaging"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"go.uber.org/zap"
)

type notificationPublisher struct {
	kafkaManager *kafka.Manager
	logger       *zap.Logger
}

// Register implements notificationmessaging.NotificationPublisher.
func (p *notificationPublisher) Register() {

}

func NewNotificationPublisher(kafkaManager *kafka.Manager, logger *zap.Logger) notificationmessaging.NotificationPublisher {
	return &notificationPublisher{kafkaManager: kafkaManager, logger: logger}
}
