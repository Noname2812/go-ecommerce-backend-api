package authserviceimpl

import (
	"context"
	"fmt"
	"time"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/consts"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils/random"
	authcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto"
	authservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/service"
	authdomainevent "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/event"
	authrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/repository"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"go.uber.org/zap"
)

type authCommandService struct {
	logger             *zap.Logger
	userBaseRepo       authrepository.UserBaseRepository
	authCacheService   authservice.AuthCacheService
	authEventPublisher authservice.AuthEventPublisher
}

// Register implements authservice.AuthCommandService.
func (a *authCommandService) Register(ctx context.Context, input *authcommandrequest.UserRegistratorRequest) (int, error) {

	// 1. Check if user is blocked
	isBlocked, err := a.authCacheService.IsUserBlocked(ctx, input.Email)
	if err != nil {
		return response.ErrCodeParamInvalid, err
	}
	if isBlocked {
		return response.ErrCodeUserBlocked, fmt.Errorf("user is temporarily blocked")
	}

	// 2. check user exists in user base
	userFound, err := a.userBaseRepo.CheckUserBaseExists(ctx, input.Email)
	if err != nil {
		return response.ErrCodeUserHasExists, err
	}
	if userFound {
		return response.ErrCodeUserHasExists, nil
	}
	// 3. Generate OTP
	otpNew := random.GenerateSixDigitOtp()
	if input.Purpose == consts.PurposeTestUser {
		otpNew = 123456
	}

	// 4. save OTP in Db
	// 5. save OTP in Redis with expiration time
	_, err = a.authCacheService.SetOTP(ctx, input.Email, fmt.Sprint(otpNew))
	if err != nil {
		return response.ErrInvalidOTP, err
	}

	// 6. Create event and publish to kafka
	payload := authdomainevent.NewOtpVerifyRegisterEvent(input.Email, fmt.Sprint(otpNew), time.Duration(consts.TIME_OTP_REGISTER)*time.Minute)

	// 7. send to kafka
	if err := a.authEventPublisher.PublishOtpVertifyUserRegisterCreated(ctx, payload); err != nil {
		// revert redis
		a.authCacheService.ClearOTP(ctx, input.Email)
		return response.ErrCodeParamInvalid, err
	}

	return response.ErrCodeSuccess, nil
}

func NewAuthCommandService(logger *zap.Logger,
	userBaseRepo authrepository.UserBaseRepository,
	authCacheService authservice.AuthCacheService,
	authEventPublisher authservice.AuthEventPublisher) authservice.AuthCommandService {
	return &authCommandService{
		logger:             logger,
		userBaseRepo:       userBaseRepo,
		authCacheService:   authCacheService,
		authEventPublisher: authEventPublisher,
	}
}
