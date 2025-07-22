package useremessagingimpl

import (
	"context"
	"encoding/json"

	usermessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/messaging"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"go.uber.org/zap"
)

type userPublisher struct {
	kafkaManager *kafka.Manager
	logger       *zap.Logger
}

// Register implements usermessaging.UserPublisher.
func (p *userPublisher) Register() {
	panic("unimplemented")
}

func NewUserPublisher(manager *kafka.Manager, logger *zap.Logger) usermessaging.UserPublisher {
	return &userPublisher{
		kafkaManager: manager,
		logger:       logger,
	}
}

func (p *userPublisher) publishEvent(ctx context.Context, topic, key string, event interface{}) error {
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
