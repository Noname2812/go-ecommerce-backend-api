package authdomainevent

import (
	"time"

	commonenum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
	commonkafka "github.com/Noname2812/go-ecommerce-backend-api/internal/common/kafka"
)

type OtpVertifyUserRegisterCreated struct {
	Email     string    `json:"email"`      // Email of user register
	Value     string    `json:"value"`      // UUID OTP
	OtpType   int       `json:"otp_type"`   // 1: register, 2: forgot, ...
	ExpiresAt int64     `json:"expires_at"` // UNIX timestamp OTP expires
	CreatedAt time.Time `json:"created_at"`
}

func (o OtpVertifyUserRegisterCreated) EventName() string {
	return commonkafka.TOPIC_OTP_CREATED
}

func NewOtpVerifyRegisterEvent(email string, otp string, duration time.Duration) *OtpVertifyUserRegisterCreated {
	return &OtpVertifyUserRegisterCreated{
		Email:     email,
		Value:     otp,
		OtpType:   int(commonenum.OtpTypeRegister),
		ExpiresAt: time.Now().Add(duration).Unix(),
		CreatedAt: time.Now(),
	}
}
