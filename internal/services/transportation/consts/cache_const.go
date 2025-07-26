package transportationconsts

import "time"

// Cache keys
const (
	TRIPS_KEY_PREFIX = "trips"
)

// Cache TTL
const (
	TRIPS_KEY_REDIS_TTL = 10 * time.Minute // 10 minutes
	TRIPS_KEY_LOCAL_TTL = 1 * time.Minute  // 1 minute
)

// Lock configuration
const (
	TRIPS_LOCK_TTL_SECONDS       = 5 * time.Second        // Lock TTL in seconds
	RETRY_GET_LIST_TRIPS_BACKOFF = 100 * time.Millisecond // Backoff delay in milliseconds
)
