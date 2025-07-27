package useremessagingimpl

import (
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
