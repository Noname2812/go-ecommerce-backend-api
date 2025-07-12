package userinforepositoryimpl

import (
	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils"
	"github.com/Noname2812/go-ecommerce-backend-api/internal/database"
	usermodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/model"
)

func ToUserInfo(user database.GetUserByUserIdRow) usermodel.UserInfo {
	return usermodel.UserInfo{
		UserID:                  user.UserID,
		UserAccount:             user.UserAccount,
		UserNickname:            user.UserNickname.String,
		UserAvatar:              utils.NullStringToPtr(user.UserAvatar),
		UserState:               utils.Uint8ToUserState(user.UserState),
		UserPhone:               utils.PhoneFromNullString(user.UserPhone),
		UserGender:              utils.NullInt16ToGender(user.UserGender),
		UserBirthday:            utils.NullTimeToPtr(user.UserBirthday),
		UserAuthenticationState: utils.Uint8ToAuthState(user.UserIsAuthentication),
		CreatedAt:               user.CreatedAt.Time,
		UpdatedAt:               user.UpdatedAt.Time,
		DeletedAt:               nil,
	}
}
