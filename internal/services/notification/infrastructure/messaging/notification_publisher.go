package notificationmessagingimpl

import (
	"context"
	"encoding/json"

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

// Publisher implements notificationmessaging.NotificationPublisher.
func (p *notificationPublisher) publishEvent(ctx context.Context, topic, key string, event interface{}) error {
	eventData, err := json.Marshal(event)
	if err != nil {
		p.logger.Error("failed to marshal event",
			zap.String("topic", topic),
			zap.String("key", key),
			zap.Error(err),
		)
		return err
	}

	if err := p.kafkaManager.SendMessage(ctx, topic, []byte(key), eventData); err != nil {
		p.logger.Error("failed to publish event",
			zap.String("topic", topic),
			zap.String("key", key),
			zap.Error(err),
		)
		return err
	}

	p.logger.Info("event published successfully",
		zap.String("topic", topic),
		zap.String("key", key),
	)

	return nil
}

func NewNotificationPublisher(kafkaManager *kafka.Manager, logger *zap.Logger) notificationmessaging.NotificationPublisher {
	return &notificationPublisher{kafkaManager: kafkaManager, logger: logger}
}
