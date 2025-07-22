package commonenum

// OTPType indicates the type of OTP.
type OTPType int16

const (
	OtpTypeRegister OTPType = 0 // Register: OTP for register

	OtpTypeForgot OTPType = 1 // Forgot: OTP for forgot password

	OtpType2FA OTPType = 2 // 2FA: OTP for 2FA
)
