package userserviceimpl

import (
	"context"
	"fmt"

	userservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/service"
	usermodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/model"
	repository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/repository"
)

type userQueryService struct {
	userInfoRepo repository.UserInfoRepository
}

// GetUserProfile implements userservice.IUserService.
func (us *userQueryService) GetUserProfile(ctx context.Context, id uint64) (*usermodel.UserInfo, error) {
	user, err := us.userInfoRepo.FindUserInfoByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("find user info by id failed")
	}
	return user, nil
}

func NewUserQueryService(uir repository.UserInfoRepository) userservice.UserQueryService {
	return &userQueryService{userInfoRepo: uir}
}
