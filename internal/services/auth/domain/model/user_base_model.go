package authmodel

import (
	"time"

	vo "github.com/Noname2812/go-ecommerce-backend-api/internal/common/vo"
)

// UserBase model
type UserBase struct {
	UserID uint64 // primary key

	UserAccount vo.Email // account (ex: "john.doe@example.com")

	UserPassword string // password (ex: "123456")

	UserSalt string // salt (ex: "xzx312sa123")

	UserLoginTime *time.Time // login time (ex: 2021-01-01 00:00:00)

	UserLogoutTime *time.Time // logout time (ex: 2021-01-01 00:00:00)

	UserLoginIp string // login IP address (ex: "127.0.0.1")

	IsTwoFactorEnabled bool // is 2FA enabled (ex: true, false)

	CreatedAt time.Time // created at (ex: 2021-01-01 00:00:00)

	UpdatedAt time.Time // updated at (ex: 2021-01-01 00:00:00)

	DeletedAt *time.Time // deleted at (ex: 2021-01-01 00:00:00)
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
