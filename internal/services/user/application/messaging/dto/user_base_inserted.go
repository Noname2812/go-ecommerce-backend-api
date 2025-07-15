package usermessagingevent

type UserBaseInserted struct {
	Email   string `json:"email"`   // Email of user register
	Success bool   `json:"success"` // Register status
}
