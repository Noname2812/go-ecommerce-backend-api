package userdomainevent

import "time"

type UserRegistered struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Purpose   string    `json:"purpose"`
	CreatedAt time.Time `json:"created_at"`
}

func (e UserRegistered) EventName() string {
	return "user.registered"
}

func (e UserRegistered) AggregateID() string {
	return e.UserID
}

func (e UserRegistered) OccurredOn() time.Time {
	return e.CreatedAt
}
