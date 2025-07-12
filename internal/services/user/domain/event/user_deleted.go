package userdomainevent

import "time"

type UserDeleted struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Reason    string    `json:"reason,omitempty"`
	DeletedAt time.Time `json:"deleted_at"`
}

func (e UserDeleted) EventName() string {
	return "user.deleted"
}

func (e UserDeleted) AggregateID() string {
	return e.UserID
}

func (e UserDeleted) OccurredOn() time.Time {
	return e.DeletedAt
}
