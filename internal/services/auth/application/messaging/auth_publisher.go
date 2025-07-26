package authmessaging

import (
	"context"

	authdomainevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/event"
)

type AuthPublisher interface {
	Register()
	// handler
	PublishOtpVertifyUserRegisterCreated(ctx context.Context, event *authdomainevent.OtpVertifyUserRegisterCreated) error
	PublishUserBaseInsertedFail(ctx context.Context, event *authdomainevent.UserBaseInsertedFail) error
}
