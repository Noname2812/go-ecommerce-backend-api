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
	db   *sql.DB
}

// DeleteForceUserInfo implements userrepository.UserInfoRepository.
func (u *userInfoRepository) DeleteForceUserInfo(ctx context.Context, email string) error {
	err := u.sqlc.DeleteForceUser(ctx, email)
	return err
}

// Create implements userinforepository.IUserInfoRepository.
func (u *userInfoRepository) CreateUserInfo(ctx context.Context, user *usermodel.UserInfo) (uint64, error) {
	var avatar sql.NullString
	if user.UserAvatar != nil {
		avatar = sql.NullString{String: *user.UserAvatar, Valid: true}
	} else {
		avatar = sql.NullString{Valid: false}
	}
	data := &database.AddUserAutoUserIdParams{
		UserAccount:          user.UserAccount,
		UserNickname:         sql.NullString{String: user.UserNickname, Valid: true},
		UserAvatar:           avatar,
		UserState:            uint8(user.UserState),
		UserPhone:            sql.NullString{String: user.UserPhone.String(), Valid: user.UserPhone != nil},
		UserGender:           sql.NullInt16{Int16: int16(user.UserGender), Valid: true},
		UserBirthday:         sql.NullTime{Time: *user.UserBirthday, Valid: user.UserBirthday != nil},
		UserIsAuthentication: uint8(user.UserAuthenticationState),
	}
	result, err := u.sqlc.AddUserAutoUserId(ctx, *data)
	if err != nil {
		return 0, err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(userId), nil
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

func (u *userInfoRepository) GetDb() *sql.DB { return u.db }

func NewUserInfoRepository(db *sql.DB) repository.UserInfoRepository {
	return &userInfoRepository{
		sqlc: database.New(db),
		db:   db,
	}
}
