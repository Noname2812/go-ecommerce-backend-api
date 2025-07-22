package authcommandrequest

type UserRegistratorRequest struct {
	Email   string `json:"email" binding:"required,email"`                            // email
	Purpose string `json:"purpose" binding:"required,oneof=TEST_USER CUSTOMER ADMIN"` // purpose (TEST_USER, CUSTOMER, ADMIN, etc.)
}

func (u *UserRegistratorRequest) Validate() map[string]string {
	return nil
}
