package response

const (
	// Success
	ErrCodeSuccess = 20001 // Success

	// Bad Request
	ErrCodeParamInvalid = 40001 // param is invalid
	ErrCodeEmailInvalid = 40002 // Email is invalid
	ErrInvalidToken     = 40003 // token is invalid
	ErrInvalidJson      = 40004 // invalid json
	ErrInvalidOTP       = 40005 // Otp error
	ErrSendEmailOtp     = 40006 // Failed to send email OTP
	ErrLockSeatFailed   = 40007 // Lock seat failed

	// Unauthorized
	ErrCodeAuthFailed = 40100 // Authentication failed

	// Not Found
	ErrCodeUserNotFound     = 40401 // User not found
	ErrCodeOtpNotExists     = 40402 // OTP not exists
	ErrCodeUserOtpNotExists = 40403 // User OTP not exists

	// Request Timeout
	ErrRequestTimeout = 40801 // Request timeout

	// Conflict
	ErrCodeEmailExistsUserBase = 40900 // Email is exists
	ErrCodeUserHasExists       = 40901 // User has already registered

	// Forbidden

	// Unprocessable Entity
	ErrCodeTwoFactorAuthSetupFailed  = 42201 // Two Factor Authentication setup failed
	ErrCodeTwoFactorAuthVerifyFailed = 42202 // Two Factor Authentication verify failed

	// Too Many Requests
	ErrCodeUserBlocked = 42901 // User is blocked

	// Internal Server Error
	ErrServerError = 50000 // server error
)

// message
var msg = map[int]string{
	ErrCodeSuccess:                   "Success",
	ErrServerError:                   "Server error",
	ErrCodeParamInvalid:              "Param is invalid",
	ErrCodeEmailInvalid:              "Email is invalid",
	ErrCodeEmailExistsUserBase:       "Email is exists",
	ErrInvalidToken:                  "token is invalid",
	ErrInvalidOTP:                    "Otp error",
	ErrSendEmailOtp:                  "Failed to send email OTP",
	ErrCodeUserHasExists:             "user has already registered",
	ErrCodeOtpNotExists:              "OTP exists but not registered",
	ErrCodeUserOtpNotExists:          "User OTP not exists",
	ErrCodeAuthFailed:                "Authentication failed",
	ErrCodeUserBlocked:               "You are blocked by server. Please try again after 60 minutes",
	ErrCodeTwoFactorAuthSetupFailed:  "Two Factor Authentication setup failed",
	ErrCodeTwoFactorAuthVerifyFailed: "Two Factor Authentication verify failed",
	ErrCodeUserNotFound:              "User not found",
	ErrRequestTimeout:                "Request timeout",
	ErrLockSeatFailed:                "You can't book this seat because it's already booked",
}
