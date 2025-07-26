package authdomainevent

import (
	commonkafka "github.com/Noname2812/go-ecommerce-backend-api/internal/common/kafka"
)

type UserBaseInsertedFail struct {
	Email   string `json:"email"`   // Email of user register
	Success bool   `json:"success"` // status of user register
}

func (u UserBaseInsertedFail) EventName() string {
	return commonkafka.TOPIC_USER_BASE_INSERTED
}

func NewUserBaseInsertedFailEvent(email string, success bool) *UserBaseInsertedFail {
	return &UserBaseInsertedFail{
		Email:   email,
		Success: success,
	}
}
