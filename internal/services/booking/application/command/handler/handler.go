package bookingcommandhandler

import "github.com/gin-gonic/gin"

type BookingCommandHttpHandler interface {
	CreateBooking(ctx *gin.Context)
}
