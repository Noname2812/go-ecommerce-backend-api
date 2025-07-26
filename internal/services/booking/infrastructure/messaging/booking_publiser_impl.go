package bookingmessaingimpl

import (
	"context"
	"encoding/json"

	bookingmessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/application/messaging"
	kafkapkg "github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"go.uber.org/zap"
)

type bookingPublisher struct {
	logger       *zap.Logger
	kafkaManager *kafkapkg.Manager
}

// Register implements bookingmessaging.BookingPublisher.
func (b *bookingPublisher) Register() {
	panic("unimplemented")
}

func NewBookingPublisher(logger *zap.Logger, kafkaManager *kafkapkg.Manager) bookingmessaging.BookingPublisher {
	return &bookingPublisher{
		logger:       logger,
		kafkaManager: kafkaManager,
	}
}

func (p *bookingPublisher) publishEvent(ctx context.Context, topic, key string, event interface{}) error {
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
