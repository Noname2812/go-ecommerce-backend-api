package authmodel

import (
	"time"

	vo "github.com/Noname2812/go-ecommerce-backend-api/internal/common/vo"
)

// UserBase represents the basic information of a user.
type UserBase struct {
	// User ID
	UserID uint64

	// User account
	UserAccount vo.Email

	// User Password
	UserPassword string

	// User Salt
	UserSalt string

	// Last login time
	UserLoginTime *time.Time

	// Last logout time
	UserLogoutTime *time.Time

	// Login IP address
	UserLoginIp string

	// Is 2FA enabled
	IsTwoFactorEnabled bool

	// Record creation time
	CreatedAt time.Time

	// Record update time
	UpdatedAt time.Time

	// Record deletion time
	DeletedAt *time.Time
}

// NewUserBase creates a new UserBase domain entity.
func NewUserBase(account vo.Email, password string, salt string) *UserBase {
	return &UserBase{
		UserAccount:        account,
		UserPassword:       password,
		UserSalt:           salt,
		UserLoginTime:      nil,
		UserLogoutTime:     nil,
		UserLoginIp:        "",
		IsTwoFactorEnabled: false,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		DeletedAt:          nil,
	}
}

func (u *UserBase) UpdateLoginTime(ip string) {
	now := time.Now()
	u.UserLoginTime = &now
	u.UserLogoutTime = nil
	u.UserLoginIp = ip
}

func (u *UserBase) UpdateLogoutTime() {
	now := time.Now()
	u.UserLogoutTime = &now
}

func (u *UserBase) EnableTwoFactor() {
	u.IsTwoFactorEnabled = true
	u.UpdatedAt = time.Now()
}

func (u *UserBase) DisableTwoFactor() {
	u.IsTwoFactorEnabled = false
	u.UpdatedAt = time.Now()
}
