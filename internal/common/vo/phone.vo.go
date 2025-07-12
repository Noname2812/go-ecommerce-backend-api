package commonvo

import (
	"errors"
	"regexp"
)

// Phone
type Phone struct {
	value string
}

// Phone regex
var phoneRegex = regexp.MustCompile(`^\+?[0-9]{9,15}$`)

// NewPhone
func NewPhone(val string) (*Phone, error) {
	if !phoneRegex.MatchString(val) {
		return nil, errors.New("invalid phone number format")
	}
	return &Phone{value: val}, nil
}

// String value
func (p *Phone) String() string {
	return p.value
}
