package notificationservice

type EmailService interface {
	SendRegisterOTP(email, otp string) error
}
