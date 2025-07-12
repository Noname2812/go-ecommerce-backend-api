package authcommandhandler

import "github.com/gin-gonic/gin"

type AuthCommandHttpHandler interface {
	Register(ctx *gin.Context)
}

type AuthCommandGrpcHandler interface{}
