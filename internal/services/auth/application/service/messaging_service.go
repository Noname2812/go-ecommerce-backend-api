package authservice

import (
	"context"

	authdomainevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/event"
)

type AuthEventPublisher interface {
	Register()
	PublishOtpVertifyUserRegisterCreated(ctx context.Context, event *authdomainevent.OtpVertifyUserRegisterCreated) error
	PublishUserBaseInsertedFail(ctx context.Context, event *authdomainevent.UserBaseInsertedFail) error
}
