package authcommandresponse

type VerifyOTPResponse struct {
	Token   string `json:"token"`
	Expried int64  `json:"expried"`
}
