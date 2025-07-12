package authmessagingserviceimpl

import (
	"context"

	commonkafka "github.com/Noname2812/go-ecommerce-backend-api/internal/common/kafka"
	authservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/service"
	authdomainevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/event"
	kafkaCustom "github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

type authPublisher struct {
	kafkaManager *kafkaCustom.Manager
	logger       *zap.Logger
}

// Register implements authservice.AuthEventPublisher.
func (a *authPublisher) Register() {
	a.kafkaManager.RegisterTopic(commonkafka.TOPIC_OTP_CREATED, kafkaCustom.TopicConfig{
		Async:        false,
		RequiredAcks: kafka.RequireOne,
		Balancer:     &kafka.Hash{},
	})
}

// OtpVertifyUserRegisterCreated implements authservice.AuthEventPublisher.
func (a *authPublisher) PublishOtpVertifyUserRegisterCreated(ctx context.Context, event *authdomainevent.OtpVertifyUserRegisterCreated) error {
	return a.publishEvent(ctx, event.EventName(), event.Email, event)
}

func NewAuthEventPublisher(manager *kafkaCustom.Manager, logger *zap.Logger) authservice.AuthEventPublisher {
	return &authPublisher{
		kafkaManager: manager,
		logger:       logger,
	}
}

func (p *authPublisher) publishEvent(ctx context.Context, topic, key string, event interface{}) error {
	if err := p.kafkaManager.SendMessage(ctx, topic, []byte(key), event); err != nil {
		p.logger.Error("failed to publish event ",
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
