package authservice

import (
	"context"
	"time"
)

type AuthCacheService interface {
	// OTP
	SetOTP(ctx context.Context, email string, otp string) (*string, error)
	IsCheckOTP(ctx context.Context, email string, otp string) (bool, error)
	ClearOTP(ctx context.Context, email string) error
	// Block/Throttle
	IsUserBlocked(ctx context.Context, email string) (bool, error)
	BlockUser(ctx context.Context, email string, duration time.Duration) error
	UnblockUser(ctx context.Context, email string) error
}
