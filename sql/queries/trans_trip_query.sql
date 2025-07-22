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