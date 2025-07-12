package userinforepositoryimpl

import (
	"context"
	"database/sql"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/database"

	usermodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/model"
	repository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/repository"
)

type userInfoRepository struct {
	sqlc *database.Queries
}

// Create implements userinforepository.IUserInfoRepository.
func (u *userInfoRepository) CreateUserInfo(ctx context.Context, user *usermodel.UserInfo) error {
	panic("unimplemented")
}

// Delete implements userinforepository.IUserInfoRepository.
func (u *userInfoRepository) DeleteUserInfo(ctx context.Context, id uint64) error {
	panic("unimplemented")
}

// FindByID implements userinforepository.IUserInfoRepository.
func (u *userInfoRepository) FindUserInfoByID(ctx context.Context, id uint64) (*usermodel.UserInfo, error) {
	user, err := u.sqlc.GetUserByUserId(ctx, id)
	if err != nil {
		return nil, err
	}
	rs := ToUserInfo(user)
	return &rs, nil
}

// Update implements userinforepository.IUserInfoRepository.
func (u *userInfoRepository) UpdateUserInfo(ctx context.Context, user *usermodel.UserInfo) error {
	panic("unimplemented")
}

func NewUserInfoRepository(db *sql.DB) repository.UserInfoRepository {
	return &userInfoRepository{
		sqlc: database.New(db),
	}
}
