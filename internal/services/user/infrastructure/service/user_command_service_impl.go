package userserviceimpl

import (
	"context"

	usermessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/messaging"
	userservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/service"
	userrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/repository"
	"github.com/redis/go-redis/v9"
)

type userCommandService struct {
	userInfoRepo  userrepository.UserInfoRepository
	userPublisher usermessaging.UserPublisher
	redisClient   *redis.Client
}

// DeleteForceUser implements userservice.UserCommandService.
func (u *userCommandService) DeleteForceUser(ctx context.Context, email string) error {
	err := u.userInfoRepo.DeleteForceUserInfo(ctx, email)
	return err
}

func NewUserCommandService(up usermessaging.UserPublisher, uir userrepository.UserInfoRepository, rdb *redis.Client) userservice.UserCommandService {
	return &userCommandService{
		userPublisher: up,
		userInfoRepo:  uir,
		redisClient:   rdb,
	}
}
