-- name: GetTripById :one
SELECT *
FROM `trips`
WHERE trip_id = ?;

-- name: DeleteTrip :exec
UPDATE trips
SET trip_deleted_at = NOW()
WHERE trip_id = ?;

-- name: AddTrip :execresult
INSERT INTO trips (
    trip_departure_time, trip_arrival_time, trip_base_price, trip_created_at, trip_updated_at
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: UpdateTrip :execrows
UPDATE trips
SET trip_updated_at = NOW(), trip_departure_time = ?, trip_arrival_time = ?, trip_base_price = ?
WHERE trip_id = ? AND trip_updated_at = ? AND trip_deleted_at IS NULL;

-- name: DeleteForceTrip :exec
DELETE FROM trips WHERE trip_id = ?;

-- name: GetListTrips :many
SELECT trip_id, routes.route_start_location, routes.route_end_location, trips.trip_departure_time, trips.trip_arrival_time, trips.trip_base_price
FROM trips
JOIN routes ON trips.route_id = routes.route_id 
WHERE trips.trip_departure_time >= ? AND routes.route_start_location = ? AND routes.route_end_location = ? AND trips.trip_deleted_at IS NULL
ORDER BY trips.trip_departure_time ASC
LIMIT 10 OFFSET ?;

-- name: GetListTripsCount :one
SELECT COUNT(trip_id)
FROM trips
JOIN routes ON trips.route_id = routes.route_id 
WHERE trips.trip_departure_time >= ? AND routes.route_start_location = ? AND routes.route_end_location = ? AND trips.trip_deleted_at IS NULL;

-- name: GetTripDetail :one
SELECT trips.trip_id, 
    routes.route_id,
    routes.route_start_location, 
    routes.route_end_location, 
    trips.trip_departure_time, 
    trips.trip_arrival_time, 
    trips.trip_base_price,
    buses.bus_license_plate, 
    buses.bus_company, 
    buses.bus_capacity,
    buses.bus_id
FROM trips
JOIN routes ON trips.route_id = routes.route_id 
JOIN buses ON trips.bus_id = buses.bus_id
WHERE trips.trip_id = ? AND trips.trip_deleted_at IS NULL;
