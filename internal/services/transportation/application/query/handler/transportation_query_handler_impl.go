package transportationqueryhandler

import (
	transportationqueryrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/request"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type transportationQueryHandler struct {
}

// GetListTrips implements TransportationQueryHandler.
func (t *transportationQueryHandler) GetListTrips(ctx *gin.Context) {
	var request transportationqueryrequest.GetListTripsRequest
	// Parse JSON payload
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	// Validate payload
	if err := request.Validate(); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "", err)
		return
	}

	panic("unimplemented")
}

func NewTransportationQueryHandler() TransportationQueryHandler {
	return &transportationQueryHandler{}
}
