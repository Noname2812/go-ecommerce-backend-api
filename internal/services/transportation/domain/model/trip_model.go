package transportationmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

// Trip model
type Trip struct {
	TripId            uint64          // primary key
	RouteId           uint64          // route id
	BusId             uint64          // bus id
	TripDepartureTime time.Time       // departure time (ex: 2021-01-01 00:00:00)
	TripArrivalTime   time.Time       // arrival time (ex: 2021-01-01 00:00:00)
	TripBasePrice     decimal.Decimal // base price (ex: 1000000 VND)
	TripCreatedAt     time.Time       // created at (ex: 2021-01-01 00:00:00)
	TripUpdatedAt     time.Time       // updated at (ex: 2021-01-01 00:00:00)
	TripDeletedAt     *time.Time      // deleted at (ex: 2021-01-01 00:00:00)

	// Navigation property - lazy loading
	Route *Route // one to one
	Bus   *Bus   // one to one
	Seats []Seat // one to many
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
