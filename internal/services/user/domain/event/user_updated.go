package userdomainevent

import "time"

type UserUpdated struct {
	UserID    string                 `json:"user_id"`
	Changes   map[string]interface{} `json:"changes"`
	UpdatedAt time.Time              `json:"updated_at"`
}

func (e UserUpdated) EventName() string {
	return "user.updated"
}

func (e UserUpdated) AggregateID() string {
	return e.UserID
}

func (e UserUpdated) OccurredOn() time.Time {
	return e.UpdatedAt
}
