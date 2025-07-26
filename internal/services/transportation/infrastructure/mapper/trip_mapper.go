package transportationmapper

import (
	"time"

	transportationqueryresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/response"
	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
)

func TripToResponse(trip transportationmodel.Trip) transportationqueryresponse.Trip {
	return transportationqueryresponse.Trip{
		ID:            trip.TripId,
		FromLocation:  trip.Route.RouteStartLocation,
		ToLocation:    trip.Route.RouteEndLocation,
		DepartureDate: trip.TripDepartureTime.Format(time.DateOnly),
		ArrivalDate:   trip.TripArrivalTime.Format(time.DateOnly),
		Price:         trip.TripBasePrice,
	}
}
