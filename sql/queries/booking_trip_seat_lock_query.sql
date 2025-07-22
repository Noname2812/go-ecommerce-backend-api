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
    trip_id, trip_seat_lock_seat_number, locked_by_booking_id, trip_seat_lock_status, trip_seat_lock_expires_at,
    trip_seat_lock_created_at, trip_seat_lock_updated_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateTripSeatLock :execrows
UPDATE trip_seat_locks
SET booking_updated_at = NOW(), trip_seat_lock_seat_number = ?, locked_by_booking_id = ?, trip_seat_lock_status = ?,
    trip_seat_lock_expires_at = ?
WHERE trip_seat_lock_id = ? AND trip_seat_lock_updated_at = ? AND trip_seat_lock_deleted_at IS NULL;

-- name: DeleteForceTripSeatLock :exec
DELETE FROM trip_seat_locks WHERE trip_seat_lock_id = ?;