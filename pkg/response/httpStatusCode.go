package response

import "errors"

const (
	// common
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20002 // param is invalid

	// Email
	ErrCodeEmailInvalid        = 20010 // Email is invalid
	ErrCodeEmailExistsUserBase = 20011 // Email is exists

	ErrInvalidToken = 30001 // token is invalid
	ErrInvalidOTP   = 30002
	ErrSendEmailOtp = 30003

	// User Authentication
	ErrCodeAuthFailed  = 40005
	ErrCodeUserBlocked = 40000 // User is blocked
	// Register Code
	ErrCodeUserHasExists = 50001 // user has already registered

	// Err Login
	ErrCodeOtpNotExists     = 60009
	ErrCodeUserOtpNotExists = 60008

	// Two Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed  = 80001
	ErrCodeTwoFactorAuthVerifyFailed = 80002

	// User
	ErrCodeUserNotFound = 40001 // User not found
)

// message
var msg = map[int]string{
	ErrCodeSuccess:             "Success",
	ErrCodeParamInvalid:        "Param is invalid",
	ErrCodeEmailInvalid:        "Email is invalid",
	ErrCodeEmailExistsUserBase: "Email is exists",
	ErrInvalidToken:            "token is invalid",
	ErrInvalidOTP:              "Otp error",
	ErrSendEmailOtp:            "Failed to send email OTP",

	ErrCodeUserHasExists: "user has already registered",

	ErrCodeOtpNotExists:     "OTP exists but not registered",
	ErrCodeUserOtpNotExists: "User OTP not exists",
	ErrCodeAuthFailed:       "Authentication failed",
	ErrCodeUserBlocked:      "You are blocked by server. Please try again after 60 minutes",

	// Two Factor Authentication
	ErrCodeTwoFactorAuthSetupFailed:  "Two Factor Authentication setup failed",
	ErrCodeTwoFactorAuthVerifyFailed: "Two Factor Authentication verify failed",

	// User
	ErrCodeUserNotFound: "User not found",
}

// These are error sentinels
var (
	CouldNotGetTicketErr = errors.New("Could not get Ticket from MYSQL") //Type of Internal Error
)
