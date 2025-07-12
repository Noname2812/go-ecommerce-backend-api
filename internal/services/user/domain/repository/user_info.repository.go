package userrepository

import (
	"context"

	model "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/model"
)

type UserInfoRepository interface {
	CreateUserInfo(ctx context.Context, user *model.UserInfo) error
	FindUserInfoByID(ctx context.Context, id uint64) (*model.UserInfo, error)
	UpdateUserInfo(ctx context.Context, user *model.UserInfo) error
	DeleteUserInfo(ctx context.Context, id uint64) error
}
