package usermessagingevent

// UserBaseInserted is the event for user base inserted
type UserBaseInserted struct {
	Email   string `json:"email"`   // Email of user register
	Success bool   `json:"success"` // Register status
}
