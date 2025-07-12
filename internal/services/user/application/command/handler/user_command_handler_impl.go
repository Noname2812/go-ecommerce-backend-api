package usercommandhandler

import (
	userservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/service"
	"go.uber.org/zap"
)

type userCommandHttpHandler struct {
	service userservice.UserCommandService
	logger  *zap.Logger
}

func NewUserCommandHandler(logger *zap.Logger, service userservice.UserCommandService) UserCommandHandler {
	return &userCommandHttpHandler{
		logger:  logger,
		service: service,
	}
}
