package commonenum

// Gender indicates the gender of a user.
type OTPType int16

const (
	// Secret indicates an undisclosed or unknown gender.
	OtpTypeRegister OTPType = 0

	// Male indicates male gender.
	OtpTypeForgot OTPType = 1

	// Female indicates female gender.
	OtpType2FA OTPType = 2
)

// IsValid checks if the Gender value is valid.
func (o OTPType) IsValid() bool {
	return o == OtpTypeRegister || o == OtpTypeForgot || o == OtpType2FA
}

// String converts the Gender value to a string.
func (g OTPType) String() string {
	switch g {
	case OtpTypeRegister:
		return "OtpTypeRegister"
	case OtpTypeForgot:
		return "OtpTypeForgot"
	case OtpType2FA:
		return "OtpType2FA"
	default:
		return "OtpTypeRegister"
	}
}
