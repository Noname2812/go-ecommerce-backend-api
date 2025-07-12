package authserviceimpl

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/consts"
	cacheservice "github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/cache"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/crypto"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/random"
	authcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto/request"
	authcommandresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto/response"
	authservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/service"
	authdomainevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/event"
	authrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/repository"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	OTP_KEY               = "auth:otp:%s"               // auth:otp:<email>
	OTP_COUNT_SEND_KEY    = "auth:otp_count_send:%s"    // auth:max_count_send:<email>
	EMAIL_BLOCKED_KEY     = "auth:blocked:%s"           // auth:blocked:<email>
	VERIFY_OTP_FAILED_KEY = "auth:verify_otp_failed:%s" // auth:verify_otp_failed:<email>
	TOKEN_UPDATE_INFO_KEY = "auth:token_update_info:%s" // auth:token_update_info:<email>

	OTP_KEY_TTL            = 5 * time.Minute  // OTP expiry
	OTP_COUNT_SEND_KEY_TTL = 30 * time.Minute // Count OTP send expiry
	EMAIL_BLOCKED_TTL      = 60 * time.Minute // Blocked user expiry
	VERIFY_OTP_FAILED_TTL  = 10 * time.Minute // Count verify OTP failed expiry
	TOKEN_UPDATE_INFO_TTL  = 15 * time.Minute // Token update info expiry

	OTP_MAX_COUNT_SEND = 10 // Max send 10 OTP
	VERIFY_OTP_MAX_TRY = 5  // Max 5 tries to verify OTP
	MAX_LENGHT_TOKEN   = 24 // Max length token generated
)

type authCommandService struct {
	logger             *zap.Logger
	userBaseRepo       authrepository.UserBaseRepository
	redisCacheService  cacheservice.RedisCache
	authEventPublisher authservice.AuthEventPublisher
}

// VerifyOTP implements authservice.AuthCommandService.
func (a *authCommandService) VerifyOTP(ctx context.Context, input *authcommandrequest.VerifyOTPRequest) (code int, res *authcommandresponse.VerifyOTPResponse, err error) {
	hashKey := crypto.GetHash(strings.ToLower(input.Email))

	// 1. Check if user is blocked
	isBlocked, err := a.redisCacheService.Exists(ctx, fmt.Sprintf(EMAIL_BLOCKED_KEY, hashKey))
	if err != nil {
		return response.ErrServerError, nil, err
	}
	if isBlocked {
		return response.ErrCodeUserBlocked, nil, fmt.Errorf("user is temporarily blocked")
	}

	// 2. check OTP in cache
	otpCache, err := a.redisCacheService.Get(ctx, fmt.Sprintf(OTP_KEY, hashKey))
	if err != nil {
		return response.ErrServerError, nil, err
	}

	if otpCache != input.OTP {
		countVerifyOTPFailed, err := a.redisCacheService.Get(ctx, fmt.Sprintf(VERIFY_OTP_FAILED_KEY, hashKey))
		if err != nil {
			return response.ErrServerError, nil, err
		}
		if countVerifyOTPFailed == "" {
			_, err = a.redisCacheService.Incr(ctx, fmt.Sprintf(VERIFY_OTP_FAILED_KEY, hashKey))
			if err != nil {
				return response.ErrServerError, nil, err
			}
			err = a.redisCacheService.Expire(ctx, fmt.Sprintf(VERIFY_OTP_FAILED_KEY, hashKey), VERIFY_OTP_FAILED_TTL)
			if err != nil {
				return response.ErrServerError, nil, err
			}
			return response.ErrInvalidToken, nil, fmt.Errorf("invalid OTP")
		} else {
			// check count verify OTP failed for block user
			if countVerifyOTPFailedInt, err := strconv.Atoi(countVerifyOTPFailed); err == nil {
				if countVerifyOTPFailedInt >= VERIFY_OTP_MAX_TRY {
					// user is blocked
					err = a.redisCacheService.Set(ctx, fmt.Sprintf(EMAIL_BLOCKED_KEY, hashKey), "1", EMAIL_BLOCKED_TTL)
					if err != nil {
						return response.ErrServerError, nil, err
					}
					return response.ErrCodeUserBlocked, nil, fmt.Errorf("user is temporarily blocked")
				}

				// increase count verify OTP failed
				_, err = a.redisCacheService.Incr(ctx, fmt.Sprintf(VERIFY_OTP_FAILED_KEY, hashKey))
				if err != nil {
					return response.ErrServerError, nil, err
				}
			}
		}
		return response.ErrInvalidToken, nil, fmt.Errorf("invalid OTP")
	}

	// 3. create token for update information user
	token, _ := random.GenarateToken(MAX_LENGHT_TOKEN)
	err = a.redisCacheService.Set(ctx, fmt.Sprintf(TOKEN_UPDATE_INFO_KEY, hashKey), token, OTP_KEY_TTL)
	if err != nil {
		return response.ErrServerError, nil, err
	}
	res = &authcommandresponse.VerifyOTPResponse{
		Token:   token,
		Expried: time.Now().Add(OTP_KEY_TTL).Unix(),
	}
	return response.ErrCodeSuccess, res, nil
}

// Register implements authservice.AuthCommandService.
func (a *authCommandService) Register(ctx context.Context, input *authcommandrequest.UserRegistratorRequest) (int, error) {

	hashKey := crypto.GetHash(strings.ToLower(input.Email))

	// 1. Check if user is blocked
	isBlocked, err := a.redisCacheService.Exists(ctx, fmt.Sprintf(EMAIL_BLOCKED_KEY, hashKey))
	if err != nil {
		return response.ErrServerError, err
	}
	if isBlocked {
		return response.ErrCodeUserBlocked, fmt.Errorf("user is temporarily blocked")
	}

	// 2. check user exists in user base
	userFound, err := a.userBaseRepo.CheckUserBaseExists(ctx, hashKey)
	if err != nil {
		return response.ErrServerError, err
	}
	if userFound {
		return response.ErrCodeUserHasExists, nil
	}
	// 3. Generate OTP
	otpNew := random.GenerateSixDigitOtp()
	if input.Purpose == consts.PurposeTestUser {
		otpNew = 123456
	}

	// 4. save OTP in Redis with expiration time
	countSend, err := a.redisCacheService.Get(ctx, fmt.Sprintf(OTP_COUNT_SEND_KEY, hashKey))
	if err != nil {
		if err == redis.Nil {
			// set first time send OTP
			_, err = a.redisCacheService.Incr(ctx, fmt.Sprintf(OTP_COUNT_SEND_KEY, hashKey))
			if err != nil {
				return response.ErrServerError, err
			}
			err = a.redisCacheService.Expire(ctx, fmt.Sprintf(OTP_COUNT_SEND_KEY, hashKey), OTP_COUNT_SEND_KEY_TTL)
			if err != nil {
				return response.ErrServerError, err
			}
		} else {
			return response.ErrServerError, err
		}
	} else {
		// check count send OTP
		if countSendInt, err := strconv.Atoi(countSend); err == nil {
			if countSendInt >= OTP_MAX_COUNT_SEND {
				// user is blocked
				err = a.redisCacheService.Set(ctx, fmt.Sprintf(EMAIL_BLOCKED_KEY, hashKey), "1", EMAIL_BLOCKED_TTL)
				if err != nil {
					return response.ErrServerError, err
				}
				return response.ErrCodeUserBlocked, fmt.Errorf("user is temporarily blocked")
			}
			// increase count send OTP
			_, err = a.redisCacheService.Incr(ctx, fmt.Sprintf(OTP_COUNT_SEND_KEY, hashKey))
			if err != nil {
				return response.ErrServerError, err
			}
		}
	}

	err = a.redisCacheService.Set(ctx, fmt.Sprintf(OTP_KEY, hashKey), otpNew, OTP_KEY_TTL)
	if err != nil {
		return response.ErrServerError, err
	}

	// 5. send event created OTP to kafka
	payload := authdomainevent.NewOtpVerifyRegisterEvent(input.Email, fmt.Sprint(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Minute)
	if err := a.authEventPublisher.PublishOtpVertifyUserRegisterCreated(ctx, payload); err != nil {
		return response.ErrServerError, err
	}

	return response.ErrCodeSuccess, nil
}

func NewAuthCommandService(logger *zap.Logger,
	userBaseRepo authrepository.UserBaseRepository,
	rdb cacheservice.RedisCache,
	authEventPublisher authservice.AuthEventPublisher) authservice.AuthCommandService {
	return &authCommandService{
		logger:             logger,
		userBaseRepo:       userBaseRepo,
		redisCacheService:  rdb,
		authEventPublisher: authEventPublisher,
	}
}
