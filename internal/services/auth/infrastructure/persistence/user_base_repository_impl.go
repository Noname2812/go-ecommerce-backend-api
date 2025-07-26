package authrepositoryimpl

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
func (u *userBaseRepository) CreateUserBase(ctx context.Context, user *authmodel.UserBase) (uint64, error) {
	txQueries := u.getQueries(ctx)
	data := &database.AddUserBaseParams{
		UserAccount:  user.UserAccount.String(),
		UserPassword: user.UserPassword,
		UserSalt:     user.UserSalt,
	}
	result, err := txQueries.AddUserBase(ctx, *data)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
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

func (u *userBaseRepository) getQueries(ctx context.Context) *database.Queries {
	if tx, ok := ctx.Value(txKey).(*sql.Tx); ok {
		return u.sqlc.WithTx(tx)
	}
	return u.sqlc
}
