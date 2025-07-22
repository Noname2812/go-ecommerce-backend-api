package authcommandresponse

type VerifyOTPResponse struct {
	Token   string `json:"token"`   // token
	Expried int64  `json:"expried"` // expried
}
