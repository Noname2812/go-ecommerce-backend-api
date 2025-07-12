package userquerydto

import (
	"time"

	usermodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/model"
)

type GetUserByIDRequest struct {
	UserID string `json:"-" uri:"user_id" binding:"required"`
}

type GetUserByEmailRequest struct {
	Email string `json:"-" form:"email" binding:"required,email"`
}

type UserInfoQueryResponse struct {
	UserID       uint64     `json:"user_id"`
	UserAccount  string     `json:"user_account"`
	UserNickname string     `json:"user_nickname"`
	UserAvatar   *string    `json:"user_avatar,omitempty"`
	UserPhone    *string    `json:"user_phone,omitempty"`
	UserGender   int        `json:"user_gender"` // "Male", "Female", "Secret"
	UserBirthday *time.Time `json:"user_birthday,omitempty"`
}

func ToUserInfoResponse(user usermodel.UserInfo) UserInfoQueryResponse {
	var phone *string
	if user.UserPhone != nil {
		str := user.UserPhone.String()
		phone = &str
	}

	return UserInfoQueryResponse{
		UserID:       user.UserID,
		UserAccount:  user.UserAccount,
		UserNickname: user.UserNickname,
		UserAvatar:   user.UserAvatar,
		UserPhone:    phone,
		UserGender:   int(user.UserGender),
		UserBirthday: user.UserBirthday,
	}
}
