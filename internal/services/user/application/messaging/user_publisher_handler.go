package usermessaging

import (
	"context"

	userdomainevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/event"
)

type UserPublisherHandler interface {
	PublishUserRegistered(ctx context.Context, event *userdomainevent.UserRegistered) error
	PublishUserUpdated(ctx context.Context, event *userdomainevent.UserUpdated) error
	PublishUserDeleted(ctx context.Context, event *userdomainevent.UserDeleted) error
}
