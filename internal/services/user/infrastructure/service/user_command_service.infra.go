package userserviceimpl

import (
	usermessaging "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/messaging"
	userservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/service"
	userrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/repository"
)

type userCommandService struct {
	userInfoRepo     userrepository.UserInfoRepository
	userCacheService userservice.UserCacheService
	userPublisher    usermessaging.UserPublisherHandler
}

func NewUserCommandService(ucs userservice.UserCacheService, up usermessaging.UserPublisherHandler, uir userrepository.UserInfoRepository) userservice.UserCommandService {
	return &userCommandService{
		userPublisher:    up,
		userCacheService: ucs,
		userInfoRepo:     uir,
	}
}
