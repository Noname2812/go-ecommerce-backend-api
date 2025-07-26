package authservice

import (
	"context"

	authcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto/request"
	authcommandresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto/response"
)

// AuthQueryService is the interface for the query service
type AuthQueryService interface {
}

// AuthCommandService is the interface for the command service
type AuthCommandService interface {
	// Save account
	SaveAccount(ctx context.Context, input *authcommandrequest.SaveAccountRequest) (int, error)
	// Register user
	Register(ctx context.Context, input *authcommandrequest.UserRegistratorRequest) (int, error)
	// Verify OTP
	VerifyOTP(ctx context.Context, input *authcommandrequest.VerifyOTPRequest) (code int, res *authcommandresponse.VerifyOTPResponse, err error)
}
