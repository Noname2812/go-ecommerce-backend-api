package authrepository

import (
	"context"

	authmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/model"
)

type UserBaseRepository interface {
	CreateUserBase(ctx context.Context, user *authmodel.UserBase) (uint64, error)
	FindUserBaseByID(ctx context.Context, id uint64) (*authmodel.UserBase, error)
	UpdateUserBase(ctx context.Context, user *authmodel.UserBase) error
	DeleteUserBase(ctx context.Context, id uint64) error
	CheckUserBaseExists(ctx context.Context, email string) (bool, error)
}
