package userqueryhandler

import (
	"github.com/gin-gonic/gin"
)

type UserQueryHandler interface {
	GetUserDetails(ctx *gin.Context)
}
