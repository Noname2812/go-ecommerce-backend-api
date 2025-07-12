package userbaserepositoryimpl

import (
	"context"
	"database/sql"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/database"

	authmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/model"
	authrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/domain/repository"
)

type userBaseRepository struct {
	sqlc *database.Queries
}

// CheckUserBaseExists implements repository.UserBaseRepository.
func (u *userBaseRepository) CheckUserBaseExists(ctx context.Context, email string) (bool, error) {
	isExist, err := u.sqlc.CheckUserBaseExists(ctx, email)
	if err != nil {
		return false, err
	}
	if isExist > 0 {
		return true, nil
	}
	return false, nil
}

// CreateUserBase implements repository.UserBaseRepository.
func (u *userBaseRepository) CreateUserBase(ctx context.Context, user *authmodel.UserBase) error {
	panic("unimplemented")
}

// DeleteUserBase implements repository.UserBaseRepository.
func (u *userBaseRepository) DeleteUserBase(ctx context.Context, id uint64) error {
	panic("unimplemented")
}

// FindUserBaseByID implements repository.UserBaseRepository.
func (u *userBaseRepository) FindUserBaseByID(ctx context.Context, id uint64) (*authmodel.UserBase, error) {
	panic("unimplemented")
}

// UpdateUserBase implements repository.UserBaseRepository.
func (u *userBaseRepository) UpdateUserBase(ctx context.Context, user *authmodel.UserBase) error {
	panic("unimplemented")
}

func NewUserBaseRepository(db *sql.DB) authrepository.UserBaseRepository {
	return &userBaseRepository{
		sqlc: database.New(db),
	}
}
