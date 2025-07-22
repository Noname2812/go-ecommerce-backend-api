package utils

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

func NullInt64ToUint64Ptr(n sql.NullInt64) *uint64 {
	if !n.Valid {
		return nil
	}
	u := uint64(n.Int64)
	return &u
}

func NullStringToStringPtr(ns sql.NullString) *string {
	if !ns.Valid {
		return nil
	}
	return &ns.String
}

func LoadLuaScript(name string) (string, error) {
	path := filepath.Join("scripts", "lua", name)
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read Lua script: %w", err)
	}
	return string(content), nil
}
