package useremessagingimpl

import (
	"context"
	"encoding/json"

	usermessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/messaging"
	userdomainevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/event"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"go.uber.org/zap"
)

type userEventPublisher struct {
	kafkaManager *kafka.Manager
	logger       *zap.Logger
}

func NewuserEventPublisher(manager *kafka.Manager, logger *zap.Logger) usermessaging.UserPublisherHandler {
	return &userEventPublisher{
		kafkaManager: manager,
		logger:       logger,
	}
}

func (p *userEventPublisher) PublishUserRegistered(ctx context.Context, event *userdomainevent.UserRegistered) error {
	return p.publishEvent(ctx, event.EventName(), event.AggregateID(), event)
}

func (p *userEventPublisher) PublishUserUpdated(ctx context.Context, event *userdomainevent.UserUpdated) error {
	return p.publishEvent(ctx, event.EventName(), event.AggregateID(), event)
}

func (p *userEventPublisher) PublishUserDeleted(ctx context.Context, event *userdomainevent.UserDeleted) error {
	return p.publishEvent(ctx, event.EventName(), event.AggregateID(), event)
}

func (p *userEventPublisher) publishEvent(ctx context.Context, topic, key string, event interface{}) error {
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
