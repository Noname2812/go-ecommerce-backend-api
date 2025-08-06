package transportationconsts

import "time"

// Cache keys data
const (
	TRIPS_KEY       = "trips:%s-%s:%s:%d" // key  for trips (ex: trips:1-2:2025-01-01:1 => trips:from-to:date:page)
	TRIP_DETAIL_KEY = "trip_detail:%d"    // key for trip detail (ex: trip_detail:1 => trip_detail:id)
	SEATS_KEY       = "trip_seats:%d"     // key for seats (ex: trip_seats:1 => trip_seats:trip_id)
)

// Cache TTL (Redis)
const (
	TRIPS_KEY_REDIS_TTL         = 10 * time.Minute // Trips key redis TTL in minutes
	TRIP_DETAIL_KEY_REDIS_TTL   = 60 * time.Minute // Trip detail key redis TTL in minutes
	SEATS_KEY_REDIS_TTL_SECONDS = 300              // Seats key redis TTL in seconds
)

// Cache TTL (Local)
const (
	TRIPS_KEY_LOCAL_TTL       = 1 * time.Minute // Trips key local TTL in minutes
	TRIP_DETAIL_KEY_LOCAL_TTL = 3 * time.Second // Trip detail key local TTL
)

// Lock TTL
const (
	TRIPS_LOCK_TTL_SECONDS = 5 * time.Second // Trips lock TTL in seconds
	SEATS_LOCK_TTL_SECONDS = 3 * time.Second // Seats lock TTL in seconds
)

// Backoff delay
const (
	RETRY_GET_LIST_TRIPS_BACKOFF  = 100 * time.Millisecond // Retry get list trips backoff delay in milliseconds
	RETRY_GET_TRIP_DETAIL_BACKOFF = 100 * time.Millisecond // Retry get trip detail backoff delay in milliseconds
)
