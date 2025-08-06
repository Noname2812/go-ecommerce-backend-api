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
    bus_id, seat_number, seat_row_no, seat_column_no, seat_floor_no, seat_type, seat_created_at, seat_updated_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateSeat :execrows
UPDATE seats
SET seat_updated_at = NOW(), seat_number = ?, seat_row_no = ?, seat_column_no = ?, seat_floor_no = ?, seat_type = ?
WHERE seat_id = ? AND seat_updated_at = ? AND bus_id = ? AND seat_deleted_at IS NULL;

-- name: DeleteForceSeat :exec
DELETE FROM seats WHERE seat_id = ?;

-- name: GetListSeatsByBusId :many
SELECT seat_id, bus_id, seat_number, seat_row_no, seat_column_no, seat_floor_no, seat_type
FROM seats
WHERE bus_id = ? AND seat_deleted_at IS NULL;

