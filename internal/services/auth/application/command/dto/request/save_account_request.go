package authcommandrequest

type SaveAccountRequest struct {
	Token       string `json:"token" binding:"required"`                         // token
	Email       string `json:"email" binding:"required,email"`                   // email
	Password    string `json:"password" binding:"required,min=8"`                // password
	ConfirmPass string `json:"confirm_pass" binding:"required,eqfield=Password"` // confirm password
	Name        string `json:"name" binding:"required"`                          // name
	Phone       string `json:"phone" binding:"omitempty,e164"`                   // phone
	Gender      int8   `json:"gender" binding:"omitempty,oneof=0 1 2"`           // gender
	Birthday    string `json:"birthday" binding:"omitempty"`                     // birthday
}

func (s *SaveAccountRequest) Validate() map[string]string {
	return nil
}
