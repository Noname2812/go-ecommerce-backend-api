package authservice

import (
	"context"

	authcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto"
)

type AuthQueryService interface {
}

type AuthCommandService interface {
	Register(ctx context.Context, input *authcommandrequest.UserRegistratorRequest) (int, error)
}
