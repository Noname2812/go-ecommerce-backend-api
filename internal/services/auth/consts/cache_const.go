package authconsts

import "time"

const (
	OTP_KEY               = "auth:otp:%s"               // auth:otp:<email>
	OTP_COUNT_SEND_KEY    = "auth:otp_count_send:%s"    // auth:max_count_send:<email>
	EMAIL_BLOCKED_KEY     = "auth:blocked:%s"           // auth:blocked:<email>
	VERIFY_OTP_FAILED_KEY = "auth:verify_otp_failed:%s" // auth:verify_otp_failed:<email>
	TOKEN_UPDATE_INFO_KEY = "auth:token_update_info:%s" // auth:token_update_info:<email>

	OTP_KEY_TTL            = 5 * time.Minute  // OTP expiry
	OTP_COUNT_SEND_KEY_TTL = 30 * time.Minute // Count OTP send expiry
	EMAIL_BLOCKED_TTL      = 60 * time.Minute // Blocked user expiry
	VERIFY_OTP_FAILED_TTL  = 10 * time.Minute // Count verify OTP failed expiry
	TOKEN_UPDATE_INFO_TTL  = 15 * time.Minute // Token update info expiry
)
