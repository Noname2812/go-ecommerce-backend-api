package transportationqueryhandler

import "github.com/gin-gonic/gin"

type TransportationQueryHandler interface {
	GetListTrips(ctx *gin.Context)
}
