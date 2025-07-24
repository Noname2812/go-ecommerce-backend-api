package utils

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
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

func ValidateStructWithValidatorTags(validate *validator.Validate, req interface{}) map[string]string {
	errors := make(map[string]string)

	if err := validate.Struct(req); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fieldName := strings.ToLower(fieldError.Field())
			switch fieldError.Tag() {
			case "required":
				errors[fieldName] = fmt.Sprintf("%s is required", fieldError.Field())
			case "email":
				errors[fieldName] = "Email is not valid"
			case "e164":
				errors[fieldName] = "Phone number is not valid"
			case "eqfield":
				errors[fieldName] = "Password and confirm password do not match"
			case "min":
				errors[fieldName] = fmt.Sprintf("%s must be at least %s characters long", fieldError.Field(), fieldError.Param())
			case "oneof":
				errors[fieldName] = fmt.Sprintf("%s must be one of %s", fieldError.Field(), fieldError.Param())
			case "datetime":
				errors[fieldName] = "Invalid date format"
			default:
				errors[fieldName] = fmt.Sprintf("%s is not valid", fieldError.Field())
			}
		}
	}
	return errors
}

func DateValidator(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	_, err := time.Parse("2006-01-02", dateStr) // yyyy-MM-dd
	return err == nil
}
