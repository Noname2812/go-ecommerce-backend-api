package useremessagingimpl

import (
	"context"
	"encoding/json"

	userdomainevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/event"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"go.uber.org/zap"
)

type userEventConsumer struct {
	kafkaManager *kafka.Manager
	logger       *zap.Logger
}

func NewUserEventConsumer(
	kafkaManager *kafka.Manager,
	logger *zap.Logger,
) *userEventConsumer {
	return &userEventConsumer{
		kafkaManager: kafkaManager,
		logger:       logger,
	}
}

func (c *userEventConsumer) SubscribeAllChannels(ctx context.Context) error {
	// Register event handlers
	if err := c.kafkaManager.AddConsumer("user.registered", "user-service", c.handleUserRegistered, 1, nil); err != nil {
		return err
	}

	if err := c.kafkaManager.AddConsumer("user.updated", "user-service", c.handleUserUpdated, 1, nil); err != nil {
		return err
	}

	if err := c.kafkaManager.AddConsumer("user.deleted", "user-service", c.handleUserDeleted, 1, nil); err != nil {
		return err
	}
	return nil
}

func (c *userEventConsumer) handleUserRegistered(ctx context.Context, key []byte, value []byte) error {
	var event userdomainevent.UserRegistered
	if err := json.Unmarshal(value, &event); err != nil {
		c.logger.Error("failed to unmarshal user registered event", zap.Error(err))
		return err
	}

	// Handle notification/cache logic here
	return nil
}

func (c *userEventConsumer) handleUserUpdated(ctx context.Context, key []byte, value []byte) error {
	var event userdomainevent.UserUpdated
	if err := json.Unmarshal(value, &event); err != nil {
		c.logger.Error("failed to unmarshal user updated event", zap.Error(err))
		return err
	}
	// Handle notification/cache logic here
	return nil
}

func (c *userEventConsumer) handleUserDeleted(ctx context.Context, key []byte, value []byte) error {
	var event userdomainevent.UserDeleted
	if err := json.Unmarshal(value, &event); err != nil {
		c.logger.Error("failed to unmarshal user deleted event", zap.Error(err))
		return err
	}
	// Handle notification/cache logic here
	return nil
}
