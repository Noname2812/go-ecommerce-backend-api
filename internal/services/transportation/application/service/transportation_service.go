package transportationservice

import (
	"context"

	transportationqueryrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/request"
	transportationqueryresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/response"
)

type TransportationCommandService interface{}

type TransportationQueryService interface {
	GetListTrips(ctx context.Context, request *transportationqueryrequest.GetListTripsRequest) (code int, data *transportationqueryresponse.GetListTripsResponse, err error)
}
