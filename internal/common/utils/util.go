package utils

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	enum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
	vo "github.com/Noname2812/go-ecommerce-backend-api/internal/common/vo"

	"github.com/google/uuid"
)

func GetUserKey(hashKey string) string {
	return fmt.Sprintf("u:%s:otp", hashKey)
}

func GenerateCliTokenUUID(userId int) string {
	newUUID := uuid.New()
	// convert UUID to string, remove -
	uuidString := strings.ReplaceAll((newUUID).String(), "", "")
	// 10clitokenijkasdmfasikdjfpomgasdfgl,masdl;gmsdfpgk
	return strconv.Itoa(userId) + "clitoken" + uuidString
}

func NullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func NullTimeToPtr(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

func NullInt16ToGender(n sql.NullInt16) enum.Gender {
	if n.Valid {
		return enum.Gender(n.Int16)
	}
	return enum.Secret
}

func PhoneFromNullString(ns sql.NullString) *vo.Phone {
	if !ns.Valid {
		return nil
	}
	phone, _ := vo.NewPhone(ns.String)
	return phone
}

func Uint8ToAuthState(val uint8) enum.AuthenticationState {
	return enum.AuthenticationState(val)
}

func Uint8ToUserState(val uint8) enum.UserState {
	return enum.UserState(val)
}
