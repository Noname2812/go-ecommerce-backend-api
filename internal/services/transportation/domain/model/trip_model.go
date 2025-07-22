package transportationmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

type Trip struct {
	TripId            uint64
	RouteId           uint64
	BusId             uint64
	TripDepartureTime time.Time
	TripArrivalTime   time.Time
	TripBasePrice     decimal.Decimal
	TripCreatedAt     time.Time
	TripUpdatedAt     time.Time
	TripDeletedAt     *time.Time
}

func NewTrip(routeId, busId uint64, departure, arrival time.Time, price decimal.Decimal) *Trip {
	return &Trip{
		RouteId:           routeId,
		BusId:             busId,
		TripDepartureTime: departure,
		TripArrivalTime:   arrival,
		TripBasePrice:     price,
		TripCreatedAt:     time.Now(),
		TripUpdatedAt:     time.Now(),
		TripDeletedAt:     nil,
	}
}
