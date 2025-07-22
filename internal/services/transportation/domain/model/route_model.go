package transportationmodel

import "time"

type Route struct {
	RouteId                uint64
	RouteStartLocation     string
	RouteEndLocation       string
	RouteEstimatedDuration uint64
	RouteCreatedAt         time.Time
	RouteUpdatedAt         time.Time
	RouteDeletedAt         *time.Time
}

func NewRoute(start, end string, estimatedDuration uint64) *Route {
	return &Route{
		RouteStartLocation:     start,
		RouteEndLocation:       end,
		RouteEstimatedDuration: estimatedDuration,
		RouteCreatedAt:         time.Now(),
		RouteUpdatedAt:         time.Now(),
		RouteDeletedAt:         nil,
	}
}
