package usercommandrequest

type UserRegistratorRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Purpose string `json:"purpose" binding:"required"` // TEST_USER, CUSTOMER, ADMIN, etc.
}
