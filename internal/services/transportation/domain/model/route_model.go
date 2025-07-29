package transportationmodel

import "time"

// Route model
type Route struct {
	RouteId                uint64     // primary key
	RouteStartLocation     string     // start location (ex: "Hanoi", "Ho Chi Minh City")
	RouteEndLocation       string     // end location (ex: "Hanoi", "Ho Chi Minh City")
	RouteEstimatedDuration uint64     // estimated duration in minutes (ex: 600 minutes = 10 hours, 720 minutes = 12 hours)
	RouteCreatedAt         time.Time  // created at (ex: 2021-01-01 00:00:00)
	RouteUpdatedAt         time.Time  // updated at (ex: 2021-01-01 00:00:00)
	RouteDeletedAt         *time.Time // deleted at (ex: 2021-01-01 00:00:00)

	// Navigation property - lazy loading
	Trips []Trip // one to many
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
