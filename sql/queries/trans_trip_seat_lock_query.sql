-- name: GetTripSeatLockById :one
SELECT *
FROM `trip_seat_locks`
WHERE trip_seat_lock_id = ?;

-- name: DeleteTripSeatLock :exec
UPDATE trip_seat_locks
SET trip_seat_lock_deleted_at = NOW()
WHERE trip_seat_lock_id = ?;

-- name: AddTripSeatLock :execresult
INSERT INTO trip_seat_locks (
    trip_id, seat_id, locked_by_booking_id, trip_seat_lock_status, trip_seat_lock_expires_at
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: UpdateTripSeatLock :execrows
UPDATE trip_seat_locks
SET booking_updated_at = NOW(), seat_id = ?, locked_by_booking_id = ?, trip_seat_lock_status = ?,
    trip_seat_lock_expires_at = ?
WHERE trip_seat_lock_id = ? AND trip_seat_lock_updated_at = ? AND trip_seat_lock_deleted_at IS NULL;

-- name: DeleteForceTripSeatLock :exec
DELETE FROM trip_seat_locks WHERE trip_seat_lock_id = ?;

-- name: GetMapSeatLockByTripId :many
SELECT 
    s.seat_id,
    s.seat_number,
    s.seat_row_no,
    s.seat_column_no,
    s.seat_floor_no,
    s.seat_type,
    tsl.locked_by_booking_id,
    COALESCE(tsl.trip_seat_lock_status, 1) AS trip_seat_lock_status,
    tsl.trip_seat_lock_expires_at
FROM trips t
JOIN seats s ON t.bus_id = s.bus_id
LEFT JOIN trip_seat_locks tsl 
    ON s.seat_id = tsl.seat_id 
    AND tsl.trip_id = t.trip_id 
    AND tsl.trip_seat_lock_deleted_at IS NULL
WHERE t.trip_id = ?
ORDER BY s.seat_row_no, s.seat_column_no;


-- name: CreateOrUpdateSeatLock :exec
INSERT INTO trip_seat_locks (
    trip_id,
    seat_id,
    locked_by_booking_id,
    trip_seat_lock_status,
    trip_seat_lock_expires_at
) VALUES (?, ?, ?, ?, ?)
ON DUPLICATE KEY UPDATE
    trip_seat_lock_status = VALUES(trip_seat_lock_status),
    locked_by_booking_id = VALUES(locked_by_booking_id),
    trip_seat_lock_expires_at = VALUES(trip_seat_lock_expires_at);