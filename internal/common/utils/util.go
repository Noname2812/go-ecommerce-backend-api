package utils

import (
	"context"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"sort"
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

func StringToDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

// GenCacheKeyFromStruct generates a cache key from any struct (e.g. parsed query struct).
func GenCacheKeyFromStruct(input interface{}, prefix string) string {
	val := reflect.ValueOf(input)
	typ := reflect.TypeOf(input)

	if val.Kind() != reflect.Struct {
		return ""
	}

	var parts []string
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fieldVal := val.Field(i)

		// Skip field unset or zero value
		if isZeroValue(fieldVal) {
			continue
		}

		// Get field name from `json` tag if exists, fallback to `form` tag, fallback to variable name
		key := field.Tag.Get("json")
		if key == "" {
			key = field.Tag.Get("form")
		}
		if key == "" {
			key = strings.ToLower(field.Name)
		}

		// Get value as string
		valStr := fmt.Sprintf("%v", fieldVal.Interface())
		parts = append(parts, fmt.Sprintf("%s=%s", key, valStr))
	}

	// Sort to ensure key consistency
	sort.Strings(parts)
	rawQuery := strings.Join(parts, "&")
	hash := sha1.Sum([]byte(rawQuery))

	return fmt.Sprintf("%s:%x", prefix, hash)
}

// isZeroValue checks if the field has a zero value
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Int, reflect.Int64, reflect.Uint, reflect.Uint64:
		return v.Int() == 0
	case reflect.Bool:
		return !v.Bool()
	default:
		return v.IsZero()
	}
}

// retry with exponential backoff
func RetryWithExponentialBackoff(ctx context.Context, retries int, baseDelay time.Duration, getCache func() (interface{}, error)) (interface{}, error) {
	for i := 0; i < retries; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		cache, err := getCache()
		if err != nil {
			return nil, err
		}
		if cache != nil {
			return cache, nil
		}

		if i == retries-1 {
			break
		}

		// exponential backoff with jitter
		jitter := time.Duration(rand.Int63n(int64(baseDelay)))
		backoffDelay := baseDelay*time.Duration(1<<i) + jitter
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(backoffDelay):
		}
	}
	return nil, fmt.Errorf("exceeded max retries: %d", retries)
}
