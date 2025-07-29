package usermodel

import (
	"time"

	enum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
	vo "github.com/Noname2812/go-ecommerce-backend-api/internal/common/vo"
)

// User information model
type UserInfo struct {
	UserID                  uint64                   // primary key
	UserAccount             string                   // account (ex: "john.doe@example.com")
	UserNickname            string                   // nickname (ex: "John Doe")
	UserAvatar              *string                  // avatar (ex: "https://example.com/avatar.jpg")
	UserState               enum.UserState           // state (ex: 0-Locked, 1-Activated, 2-Not Activated)
	UserPhone               *vo.Phone                // phone (ex: "0909090909")
	UserGender              enum.Gender              // gender (ex: 0-Secret, 1-Male, 2-Female)
	UserBirthday            *time.Time               // birthday (ex: 2000-01-01)
	UserAuthenticationState enum.AuthenticationState // authentication state (ex: 0-Not Authenticated, 1-Pending, 2-Authenticated, 3-Failed)
	UserCreatedAt           time.Time                // created at (ex: 2021-01-01 00:00:00)
	UserUpdatedAt           time.Time                // updated at (ex: 2021-01-01 00:00:00)
	UserDeletedAt           *time.Time               // deleted at (ex: 2021-01-01 00:00:00)
}

// NewUserInfo creates a new UserInfo domain entity with validation.
func NewUserInfo(
	account string,
	nickname string,
	avatar *string,
	state enum.UserState,
	phone *vo.Phone,
	gender enum.Gender,
	birthday *time.Time,
	authState enum.AuthenticationState,
) *UserInfo {
	return &UserInfo{
		UserAccount:             account,
		UserNickname:            nickname,
		UserAvatar:              avatar,
		UserState:               state,
		UserPhone:               phone,
		UserGender:              gender,
		UserBirthday:            birthday,
		UserAuthenticationState: authState,
		UserCreatedAt:           time.Now(),
		UserUpdatedAt:           time.Now(),
		UserDeletedAt:           nil,
	}
}

// Check active
func (u *UserInfo) IsActive() bool {
	return u.UserState == enum.Activated
}

// Check authenticated
func (u *UserInfo) IsAuthenticated() bool {
	return u.UserAuthenticationState == enum.Authenticated
}

// Check locked
func (u *UserInfo) IsLocked() bool {
	return u.UserState == enum.Locked
}

// Check not activated
func (u *UserInfo) IsNotActivated() bool {
	return u.UserState == enum.NotActivated
}
