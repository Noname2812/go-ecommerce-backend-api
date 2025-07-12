package commonvo

import (
	"errors"
	"regexp"
	"strings"
)

// Email
type Email struct {
	value string
}

// NewEmail
func NewEmail(email string) (Email, error) {
	if email == "" {
		return Email{}, errors.New("email cannot be empty")
	}

	email = strings.ToLower(strings.TrimSpace(email))

	emailRegex := regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)
	if !emailRegex.MatchString(email) {
		return Email{}, errors.New("invalid email format")
	}

	return Email{value: email}, nil
}

// Value
func (e Email) Value() string {
	return e.value
}

// String
func (e Email) String() string {
	return e.value
}

// Equality comparison
func (e Email) Equals(other Email) bool {
	return e.value == other.value
}
