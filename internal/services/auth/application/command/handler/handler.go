package authcommandhandler

import "github.com/gin-gonic/gin"

// AuthCommandHttpHandler is the interface for the HTTP handler
type AuthCommandHttpHandler interface {
	// Register user
	Register(ctx *gin.Context)
	// Verify OTP
	VerifyOTP(ctx *gin.Context)
	// Save account
	SaveAccount(ctx *gin.Context)
}
