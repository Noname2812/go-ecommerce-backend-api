package authservice

import (
	"context"

	authcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto/request"
	authcommandresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto/response"
)

type AuthQueryService interface {
}

type AuthCommandService interface {
	SaveAccount(ctx context.Context, input *authcommandrequest.SaveAccountRequest) (int, error)
	Register(ctx context.Context, input *authcommandrequest.UserRegistratorRequest) (int, error)
	VerifyOTP(ctx context.Context, input *authcommandrequest.VerifyOTPRequest) (code int, res *authcommandresponse.VerifyOTPResponse, err error)
}
