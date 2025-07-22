package transportationconsts

const (
	SEATS_TRIP_KEY       = "seats:trip:%s"         // seats:trip:<trip_id>
	HOLD_SEAT_KEY        = "hold:seat:%s:%s:%s:%s" // hold:seat:<trip_id>:<seat_no>:<user_id>:<timestamp>
	TRIPS_DATE_KEY       = "trips:%s"              // trips:<yyyy-mm-dd>
	TRIPS_DATE_SEATS_KEY = "trips:%s:%s:%s"        // trips:<trip_id>:<yyyy-mm-dd>:<seat_no>
)
