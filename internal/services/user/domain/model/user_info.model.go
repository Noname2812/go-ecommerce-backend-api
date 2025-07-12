package usermodel

import (
	"time"

	enum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
	vo "github.com/Noname2812/go-ecommerce-backend-api/internal/common/vo"
)

// User information
type UserInfo struct {
	// User ID
	UserID uint64
	// User account
	UserAccount string
	// User nickname
	UserNickname string
	// User avatar
	UserAvatar *string
	// User state: 0-Locked, 1-Activated, 2-Not Activated
	UserState enum.UserState
	// Mobile phone number
	UserPhone *vo.Phone
	// User gender: 0-Secret, 1-Male, 2-Female
	UserGender enum.Gender
	// User birthday
	UserBirthday *time.Time
	// Authentication status: 0-Not Authenticated, 1-Pending, 2-Authenticated, 3-Failed
	UserAuthenticationState enum.AuthenticationState
	// Record creation time
	CreatedAt time.Time
	// Record update time
	UpdatedAt time.Time
	// Record deletion time
	DeletedAt *time.Time
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
) (*UserInfo, error) {
	return &UserInfo{
		UserAccount:             account,
		UserNickname:            nickname,
		UserAvatar:              avatar,
		UserState:               state,
		UserPhone:               phone,
		UserGender:              gender,
		UserBirthday:            birthday,
		UserAuthenticationState: authState,
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
		DeletedAt:               nil,
	}, nil
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
