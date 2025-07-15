package authcommandhandler

import "github.com/gin-gonic/gin"

type AuthCommandHttpHandler interface {
	Register(ctx *gin.Context)
	VerifyOTP(ctx *gin.Context)
	SaveAccount(ctx *gin.Context)
}
