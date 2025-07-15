package useremessagingimpl

import (
	"context"
	"encoding/json"

	commonkafka "github.com/Noname2812/go-ecommerce-backend-api/internal/common/kafka"
	usermessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/messaging"
	usermessagingevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/messaging/dto"
	userservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"go.uber.org/zap"
)

const (
	COUNT_WORKER_HANDLER_USER_BASE_INSERTED = 1
)

type userEventConsumer struct {
	kafkaManager *kafka.Manager
	logger       *zap.Logger
	userService  userservice.UserCommandService
}

// HandleUserBaseInserted implements usermessaging.UserConsumerHandler.
func (c *userEventConsumer) HandleUserBaseInserted(ctx context.Context, key []byte, value []byte) error {
	var event usermessagingevent.UserBaseInserted
	if err := json.Unmarshal(value, &event); err != nil {
		c.logger.Error("Failed to unmarshal OTP event", zap.Error(err))
		return err
	}
	if !event.Success {
		err := c.userService.DeleteForceUser(ctx, event.Email)
		return err
	}
	return nil
}

// Subscribe implements usermessaging.UserConsumerHandler.
func (c *userEventConsumer) Subscribe() error {
	// Register event handlers
	if err := c.kafkaManager.AddConsumer(commonkafka.TOPIC_USER_BASE_INSERTED, commonkafka.GROUP_USER_BASE_INSERTED, c.HandleUserBaseInserted, COUNT_WORKER_HANDLER_USER_BASE_INSERTED, nil); err != nil {
		return err
	}
	return nil
}

func NewUserConsumer(
	kafkaManager *kafka.Manager,
	logger *zap.Logger,
	userService userservice.UserCommandService,
) usermessaging.UserConsumerHandler {
	return &userEventConsumer{
		kafkaManager: kafkaManager,
		logger:       logger,
		userService:  userService,
	}
}
