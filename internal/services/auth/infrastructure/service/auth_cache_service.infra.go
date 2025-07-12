package authserviceimpl

import (
	"context"
	"fmt"
	"time"

	authservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/service"
	"github.com/redis/go-redis/v9"
)

type authCacheService struct {
	rdb *redis.Client
}

const (
	otpKeyFormat       = "user:otp:%s"       // user:otp:<email>
	otpCountKeyFormat  = "user:otp_count:%s" // user:otp_count:<email>
	blockedKeyFormat   = "user:blocked:%s"   // user:blocked:<email>
	otpTTL             = 5 * time.Minute     // OTP expiry
	otpRateLimitWindow = 30 * time.Minute    // 30p
	otpMaxSend         = 10                  // Max 10 OTP trong 30p
	blockedUserTTL     = 30 * time.Minute    // Blocked user expiry
)

// SetOTP saves the OTP and checks rate limit
func (u *authCacheService) SetOTP(ctx context.Context, email string, otp string) (*string, error) {

	// Increment OTP count
	countKey := fmt.Sprintf(otpCountKeyFormat, email)
	count, err := u.rdb.Incr(ctx, countKey).Result()
	if err != nil {
		return nil, err
	}

	if count == 1 {
		// First OTP in window → set TTL
		if err := u.rdb.Expire(ctx, countKey, otpRateLimitWindow).Err(); err != nil {
			return nil, err
		}
	}

	if count > otpMaxSend {
		// Block user
		if err := u.BlockUser(ctx, email, blockedUserTTL); err != nil {
			return nil, err
		}
		// clear OTP
		if err := u.ClearOTP(ctx, email); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("too many OTP requests, user %s is now blocked", email)
	}

	// Save OTP
	key := fmt.Sprintf(otpKeyFormat, email)
	if err := u.rdb.Set(ctx, key, otp, otpTTL).Err(); err != nil {
		return nil, err
	}

	return &key, nil
}

// IsCheckOTP checks if the provided OTP matches the one in Redis.
func (u *authCacheService) IsCheckOTP(ctx context.Context, email string, otp string) (bool, error) {
	key := fmt.Sprintf(otpKeyFormat, email)
	cachedOTP, err := u.rdb.Get(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return cachedOTP == otp, nil
}

// ClearOTP removes the OTP from Redis.
func (u *authCacheService) ClearOTP(ctx context.Context, email string) error {
	key := fmt.Sprintf(otpKeyFormat, email)
	err := u.rdb.Del(ctx, key).Err()
	return err
}

// BlockUser marks a user as blocked in Redis for the given duration.
func (u *authCacheService) BlockUser(ctx context.Context, email string, duration time.Duration) error {
	key := fmt.Sprintf(blockedKeyFormat, email)
	err := u.rdb.Set(ctx, key, true, duration).Err()
	return err
}

// IsUserBlocked checks if the user is blocked in Redis.
func (u *authCacheService) IsUserBlocked(ctx context.Context, email string) (bool, error) {
	key := fmt.Sprintf(blockedKeyFormat, email)
	exists, err := u.rdb.Exists(ctx, key).Result()
	return exists == 1, err
}

// UnblockUser removes the blocked user key from Redis.
func (u *authCacheService) UnblockUser(ctx context.Context, email string) error {
	key := fmt.Sprintf(blockedKeyFormat, email)
	err := u.rdb.Del(ctx, key).Err()
	return err
}

// NewAuthCacheService creates a new instance of AuthCacheService.
func NewAuthCacheService(rd *redis.Client) authservice.AuthCacheService {
	return &authCacheService{
		rdb: rd,
	}
}
