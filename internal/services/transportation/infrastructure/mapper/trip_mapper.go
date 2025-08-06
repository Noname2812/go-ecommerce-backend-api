package transportationmapper

import (
	"time"

	transportationqueryresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/response"
	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
)

// mapper trip model to trip item response
func TripItemToResponse(trip *transportationmodel.Trip) *transportationqueryresponse.Trip {
	return &transportationqueryresponse.Trip{
		ID:            trip.TripId,
		FromLocation:  trip.Route.RouteStartLocation,
		ToLocation:    trip.Route.RouteEndLocation,
		DepartureTime: trip.TripDepartureTime.Format(time.DateTime),
		ArrivalTime:   trip.TripArrivalTime.Format(time.DateTime),
		Price:         trip.TripBasePrice,
	}
}

// mapper trip model to trip detail response
func TripDetailToResponse(trip *transportationmodel.Trip) *transportationqueryresponse.TripDetailResponse {
	return &transportationqueryresponse.TripDetailResponse{
		TripID:        trip.TripId,
		DepartureTime: trip.TripDepartureTime.Format(time.DateTime),
		ArrivalTime:   trip.TripArrivalTime.Format(time.DateTime),
		BasePrice:     trip.TripBasePrice,
		Bus: transportationqueryresponse.Bus{
			BusID:           trip.Bus.BusId,
			BusLicensePlate: trip.Bus.BusLicensePlate,
			BusCompany:      trip.Bus.BusCompany,
			BusCapacity:     trip.Bus.BusCapacity,
		},
		Route: transportationqueryresponse.Route{
			RouteID:      trip.Route.RouteId,
			FromLocation: trip.Route.RouteStartLocation,
			ToLocation:   trip.Route.RouteEndLocation,
		},
	}
}
