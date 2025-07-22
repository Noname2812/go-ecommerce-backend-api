package bookingmessaingimpl

import (
	"context"

	commonenum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
	bookingmessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/application/messaging"
	bookingmessagingdto "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/application/messaging/dto"
	kafkapkg "github.com/Noname2812/go-ecommerce-backend-api/pkg/kafka"
	"go.uber.org/zap"
)

type bookingConsumer struct {
	logger       *zap.Logger
	kafkaManager *kafkapkg.Manager
}

// HandlePaymentFailed implements bookingmessaging.BookingConsumer.
func (b *bookingConsumer) HandlePaymentFailed(ctx context.Context, key []byte, value []byte) error {
	event, err := bookingmessagingdto.UnmarshalPaymentFailed(value)
	if err != nil {
		b.logger.Error("Failed to unmarshal payment failed event", zap.Error(err))
		return err
	}

	if event.Status == commonenum.EVENTFAILED {

	}

	return nil
}

// Subscribe implements bookingmessaging.BookingConsumer.
func (b *bookingConsumer) Subscribe() error {
	panic("unimplemented")
}

func NewBookingConsumer(logger *zap.Logger, kafkaManager *kafkapkg.Manager) bookingmessaging.BookingConsumer {
	return &bookingConsumer{
		logger:       logger,
		kafkaManager: kafkaManager,
	}
}
