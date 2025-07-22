package authserviceimpl

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/consts"
	userpb "github.com/Noname2812/go-ecommerce-backend-api/internal/common/protogen/user"
	cacheservice "github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/cache"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/crypto"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/random"
	commonvo "github.com/Noname2812/go-ecommerce-backend-api/internal/common/vo"
	authcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto/request"
	authcommandresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto/response"
	authmessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/messaging"
	authservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/service"
	authconsts "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/consts"
	authdomainevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/event"
	authmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/model"
	authrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/repository"
	authclientgrpc "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/infrastructure/client"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

type authCommandService struct {
	logger             *zap.Logger
	userBaseRepo       authrepository.UserBaseRepository
	redisCacheService  cacheservice.RedisCache
	authEventPublisher authmessaging.AuthPublisher
	userClient         *authclientgrpc.UserGRPCClient
	transactionManager authrepository.TransactionManager
}

// SaveAccount implements authservice.AuthCommandService.
func (a *authCommandService) SaveAccount(ctx context.Context, input *authcommandrequest.SaveAccountRequest) (code int, err error) {
	hashKey := crypto.GetHash(strings.ToLower(input.Email))
	// 1. check user exists in user base
	userFound, err := a.userBaseRepo.CheckUserBaseExists(ctx, input.Email)
	if err != nil {
		return response.ErrServerError, err
	}
	if userFound {
		return response.ErrCodeEmailExistsUserBase, fmt.Errorf("email already exists")
	}
	// 2. check token
	raw, err := a.redisCacheService.Get(ctx, fmt.Sprintf(authconsts.TOKEN_UPDATE_INFO_KEY, hashKey))
	if err != nil {
		return response.ErrServerError, err
	}
	var token string
	if err := json.Unmarshal([]byte(raw), &token); err != nil {
		return response.ErrServerError, err
	}
	if token != input.Token {
		return response.ErrInvalidToken, fmt.Errorf("invalid token")
	}

	email, err := commonvo.NewEmail(input.Email)
	if err != nil {
		a.logger.Warn("Email is invalid", zap.String("email", email.String()), zap.Error(err))
		return response.ErrCodeEmailInvalid, err
	}

	err = a.transactionManager.WithTransaction(ctx, func(txCtx context.Context) error {
		// 3. save user base
		salt, _ := crypto.GenerateSalt(authconsts.MAX_LENGHT_SALT)
		password := crypto.HashPassword(input.Password, salt)
		userBase := authmodel.NewUserBase(email, password, salt)
		_, err := a.userBaseRepo.CreateUserBase(txCtx, userBase)
		if err != nil {
			return err
		}

		// 4. save user info
		req := &userpb.CreateUserRequest{
			UserAccount:  input.Email,
			UserNickName: input.Name,
			UserPhone:    input.Phone,
			UserGender:   int32(input.Gender),
			UserBirthday: input.Birthday,
		}
		_, err = a.userClient.CreateUser(txCtx, req)
		if err != nil {
			return err
		}
		// clear token
		return a.redisCacheService.Del(ctx, fmt.Sprintf(authconsts.TOKEN_UPDATE_INFO_KEY, hashKey))
	})

	if err != nil {
		// 5. send event user base inserted fail
		event := &authdomainevent.UserBaseInsertedFail{Email: email.String(), Success: false}
		_ = a.authEventPublisher.PublishUserBaseInsertedFail(ctx, event)
		return response.ErrServerError, err
	}
	return response.ErrCodeSuccess, nil
}

// VerifyOTP implements authservice.AuthCommandService.
func (a *authCommandService) VerifyOTP(ctx context.Context, input *authcommandrequest.VerifyOTPRequest) (code int, res *authcommandresponse.VerifyOTPResponse, err error) {
	hashKey := crypto.GetHash(strings.ToLower(input.Email))

	// 1. Check if user is blocked
	isBlocked, err := a.redisCacheService.Exists(ctx, fmt.Sprintf(authconsts.EMAIL_BLOCKED_KEY, hashKey))
	if err != nil {
		return response.ErrServerError, nil, err
	}
	if isBlocked {
		return response.ErrCodeUserBlocked, nil, fmt.Errorf("user is temporarily blocked")
	}

	// 2. check OTP in cache
	otpCache, err := a.redisCacheService.Get(ctx, fmt.Sprintf(authconsts.OTP_KEY, hashKey))
	if err != nil {
		return response.ErrServerError, nil, err
	}

	if otpCache != input.OTP {
		countVerifyOTPFailed, err := a.redisCacheService.Get(ctx, fmt.Sprintf(authconsts.VERIFY_OTP_FAILED_KEY, hashKey))
		if err != nil {
			return response.ErrServerError, nil, err
		}
		if countVerifyOTPFailed == "" {
			_, err = a.redisCacheService.Incr(ctx, fmt.Sprintf(authconsts.VERIFY_OTP_FAILED_KEY, hashKey))
			if err != nil {
				return response.ErrServerError, nil, err
			}
			err = a.redisCacheService.Expire(ctx, fmt.Sprintf(authconsts.VERIFY_OTP_FAILED_KEY, hashKey), authconsts.VERIFY_OTP_FAILED_TTL)
			if err != nil {
				return response.ErrServerError, nil, err
			}
			return response.ErrInvalidToken, nil, fmt.Errorf("invalid OTP")
		} else {
			// check count verify OTP failed for block user
			if countVerifyOTPFailedInt, err := strconv.Atoi(countVerifyOTPFailed); err == nil {
				if countVerifyOTPFailedInt >= authconsts.VERIFY_OTP_MAX_TRY {
					// user is blocked
					err = a.redisCacheService.Set(ctx, fmt.Sprintf(authconsts.EMAIL_BLOCKED_KEY, hashKey), "1", authconsts.EMAIL_BLOCKED_TTL)
					if err != nil {
						return response.ErrServerError, nil, err
					}
					return response.ErrCodeUserBlocked, nil, fmt.Errorf("user is temporarily blocked")
				}

				// increase count verify OTP failed
				_, err = a.redisCacheService.Incr(ctx, fmt.Sprintf(authconsts.VERIFY_OTP_FAILED_KEY, hashKey))
				if err != nil {
					return response.ErrServerError, nil, err
				}
			}
		}
		return response.ErrInvalidToken, nil, fmt.Errorf("invalid OTP")
	}

	// 3. create token for update information user
	token, _ := random.GenarateToken(authconsts.MAX_LENGHT_TOKEN)
	err = a.redisCacheService.Set(ctx, fmt.Sprintf(authconsts.TOKEN_UPDATE_INFO_KEY, hashKey), token, authconsts.OTP_KEY_TTL)
	if err != nil {
		return response.ErrServerError, nil, err
	}

	// 4. clear otp
	err = a.redisCacheService.Del(ctx, fmt.Sprintf(authconsts.OTP_KEY, hashKey))
	if err != nil {
		return response.ErrServerError, nil, err
	}
	// 5. Response token register
	res = &authcommandresponse.VerifyOTPResponse{
		Token:   token,
		Expried: time.Now().Add(authconsts.OTP_KEY_TTL).Unix(),
	}
	return response.ErrCodeSuccess, res, nil
}

// Register implements authservice.AuthCommandService.
func (a *authCommandService) Register(ctx context.Context, input *authcommandrequest.UserRegistratorRequest) (int, error) {

	hashKey := crypto.GetHash(strings.ToLower(input.Email))

	// 1. Check if user is blocked
	isBlocked, err := a.redisCacheService.Exists(ctx, fmt.Sprintf(authconsts.EMAIL_BLOCKED_KEY, hashKey))
	if err != nil {
		return response.ErrServerError, err
	}
	if isBlocked {
		return response.ErrCodeUserBlocked, fmt.Errorf("user is temporarily blocked")
	}

	// 2. check user exists in user base
	userFound, err := a.userBaseRepo.CheckUserBaseExists(ctx, input.Email)
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
	countSend, err := a.redisCacheService.Get(ctx, fmt.Sprintf(authconsts.OTP_COUNT_SEND_KEY, hashKey))
	if err != nil {
		return response.ErrServerError, err
	}
	if countSend == "" {
		// set first time send OTP
		_, err = a.redisCacheService.Incr(ctx, fmt.Sprintf(authconsts.OTP_COUNT_SEND_KEY, hashKey))
		if err != nil {
			return response.ErrServerError, err
		}
		err = a.redisCacheService.Expire(ctx, fmt.Sprintf(authconsts.OTP_COUNT_SEND_KEY, hashKey), authconsts.OTP_COUNT_SEND_KEY_TTL)
		if err != nil {
			return response.ErrServerError, err
		}
	} else {
		// check count send OTP
		if countSendInt, err := strconv.Atoi(countSend); err == nil {
			if countSendInt >= authconsts.OTP_MAX_COUNT_SEND {
				// user is blocked
				err = a.redisCacheService.Set(ctx, fmt.Sprintf(authconsts.EMAIL_BLOCKED_KEY, hashKey), "1", authconsts.EMAIL_BLOCKED_TTL)
				if err != nil {
					return response.ErrServerError, err
				}
				return response.ErrCodeUserBlocked, fmt.Errorf("user is temporarily blocked")
			}
			// increase count send OTP
			_, err = a.redisCacheService.Incr(ctx, fmt.Sprintf(authconsts.OTP_COUNT_SEND_KEY, hashKey))
			if err != nil {
				return response.ErrServerError, err
			}
		}
	}

	err = a.redisCacheService.Set(ctx, fmt.Sprintf(authconsts.OTP_KEY, hashKey), otpNew, authconsts.OTP_KEY_TTL)
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
	authEventPublisher authmessaging.AuthPublisher,
	userClient *authclientgrpc.UserGRPCClient,
	transactionManager authrepository.TransactionManager,
) authservice.AuthCommandService {
	return &authCommandService{
		logger:             logger,
		userBaseRepo:       userBaseRepo,
		redisCacheService:  rdb,
		authEventPublisher: authEventPublisher,
		userClient:         userClient,
		transactionManager: transactionManager,
	}
}
