package notificationmessagingimpl

import (
	"context"
	"encoding/json"

	commonkafka "github.com/Noname2812/go-ecommerce-backend-api/internal/common/kafka"
	notificationmessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/application/messaging"
	notificationmessagingevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/application/messaging/dto"
	notificationservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/notification/application/service"
	kafkaManager "github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"go.uber.org/zap"
)

const (
	COUNT_WORKER_HANDLER_OTP = 2
)

type notificationConsumer struct {
	kafkaManager *kafkaManager.Manager
	logger       *zap.Logger
	emailService notificationservice.EmailService
}

// HandleOtpVerifyUserRegisterCreatedEvent implements notificationmessaging.NotificationConsumer.
func (c *notificationConsumer) HandleOtpVerifyUserRegisterCreatedEvent(ctx context.Context, key, value []byte) error {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error("Panic recovered in HandleOtpVerifyUserRegisterCreatedEvent",
				zap.Any("recover", r),
				zap.Stack("stacktrace"))
		}
	}()
	var event notificationmessagingevent.OtpCreatedEvent
	if err := json.Unmarshal(value, &event); err != nil {
		c.logger.Error("Failed to unmarshal OTP event", zap.Error(err))
		return err
	}

	err := c.emailService.SendRegisterOTP(event.Email, event.Value)
	if err != nil {
		c.logger.Error("Failed to send OTP notification", zap.Error(err))
		return err
	}

	return nil
}

func (c *notificationConsumer) Subscribe() error {
	// Register event handlers
	if err := c.kafkaManager.AddConsumer(commonkafka.TOPIC_OTP_CREATED, commonkafka.GROUP_OTP, c.HandleOtpVerifyUserRegisterCreatedEvent, COUNT_WORKER_HANDLER_OTP, nil); err != nil {
		return err
	}
	return nil
}

func NewNotificationConsumer(
	kafkaManager *kafkaManager.Manager,
	logger *zap.Logger,
	emailService notificationservice.EmailService,
) notificationmessaging.NotificationConsumer {
	return &notificationConsumer{
		kafkaManager: kafkaManager,
		logger:       logger,
		emailService: emailService,
	}
}
