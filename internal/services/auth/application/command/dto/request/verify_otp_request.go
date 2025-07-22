package authcommandrequest

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"` // email
	OTP   string `json:"otp" binding:"required"`         // otp
}

func (v *VerifyOTPRequest) Validate() map[string]string {
	return nil
}
