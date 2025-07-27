package bookingmessaingimpl

import (
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
