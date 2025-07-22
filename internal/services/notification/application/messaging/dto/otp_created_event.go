package notificationmessagingevent

import (
	"encoding/json"
	"time"
)

type OtpCreatedEvent struct {
	Email     string    `json:"email"`      // Email of user register
	Value     string    `json:"value"`      // UUID OTP
	OtpType   int       `json:"otp_type"`   // 1: register, 2: forgot, ...
	ExpiresAt int64     `json:"expires_at"` // UNIX timestamp OTP expires
	CreatedAt time.Time `json:"created_at"`
}

func UnmarshalOtpCreatedEvent(data []byte) (*OtpCreatedEvent, error) {
	var otpCreatedEvent OtpCreatedEvent
	if err := json.Unmarshal(data, &otpCreatedEvent); err != nil {
		return nil, err
	}
	return &otpCreatedEvent, nil
}
