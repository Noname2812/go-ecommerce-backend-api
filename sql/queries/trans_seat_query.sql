-- name: GetSeatById :one
SELECT *
FROM `seats`
WHERE seat_id = ?;

-- name: DeleteSeat :exec
UPDATE seats
SET seat_deleted_at = NOW()
WHERE seat_id = ?;

-- name: AddSeat :execresult
INSERT INTO seats (
    bus_id, seat_number, is_available, seat_created_at, seat_updated_at
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: UpdateSeat :execrows
UPDATE seats
SET seat_updated_at = NOW(), seat_number = ?, is_available = ?
WHERE seat_id = ? AND seat_updated_at = ? AND bus_id = ? AND seat_deleted_at IS NULL;

-- name: DeleteForceSeat :exec
DELETE FROM seats WHERE seat_id = ?;