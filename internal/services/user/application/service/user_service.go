package userservice

import (
	"context"

	usermodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/model"
)

type UserQueryService interface {
	GetUserProfile(ctx context.Context, id uint64) (*usermodel.UserInfo, error)
}

type UserCommandService interface{}
