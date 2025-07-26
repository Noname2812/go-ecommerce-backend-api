package transportationqueryhandler

import (
	"strconv"

	transportationqueryrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/request"
	transportationservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type transportationQueryHandler struct {
	transportationQueryService transportationservice.TransportationQueryService
}

// GetListTrips implements TransportationQueryHandler.
// @Summary Get list trips
// @Description Get list trips
// @Tags Transportation
// @Accept json
// @Produce json
// @Param departure_date query string true "Departure date"
// @Param from_location  query string true "From location"
// @Param to_location    query string true "To location"
// @Param page           query int    false "Page number (default 1)"
// @Param page_size      query int    false "Page size (default 10)"
// @Success 200 {object} response.ResponseData{data=transportationqueryresponse.GetListTripsResponse}
// @Failure 400 {object} response.ErrorResponseData
// @Failure 408 {object} response.ErrorResponseData
// @Failure 500 {object} response.ErrorResponseData
// @Router /transportation/search-trips [get]
func (t *transportationQueryHandler) GetListTrips(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}

	request := transportationqueryrequest.GetListTripsRequest{
		DepartureDate: ctx.Query("departure_date"),
		FromLocation:  ctx.Query("from_location"),
		ToLocation:    ctx.Query("to_location"),
		Page:          page,
	}

	// Validate payload
	if err := request.Validate(ctx); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "", err)
		return
	}

	// call service
	code, data, err := t.transportationQueryService.GetListTrips(ctx, &request)
	if err != nil {
		response.ErrorResponse(ctx, code, err.Error(), nil)
		return
	}

	// response
	response.SuccessResponse(ctx, code, data)
}

func NewTransportationQueryHandler(transportationQueryService transportationservice.TransportationQueryService) TransportationQueryHandler {
	return &transportationQueryHandler{
		transportationQueryService: transportationQueryService,
	}
}
