-- name: GetSeatBookingById :one
SELECT *
FROM `seat_bookings`
WHERE seat_booking_id = ?;

-- name: DeleteSeatBooking :exec
UPDATE seat_bookings
SET seat_booking_deleted_at = NOW()
WHERE seat_booking_id = ?;

-- name: AddSeatBooking :execresult
INSERT INTO seat_bookings (
    booking_id, seat_booking_seat_number, seat_booking_price, passenger_name, passenger_phone,
    seat_booking_created_at, seat_booking_updated_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateSeatBooking :execrows
UPDATE seat_bookings
SET booking_updated_at = NOW(), seat_booking_seat_number = ?, seat_booking_price = ?, passenger_name = ?,
    passenger_phone = ?
WHERE seat_booking_id = ? AND seat_booking_updated_at = ? AND seat_booking_deleted_at IS NULL;

-- name: DeleteForceSeatBooking :exec
DELETE FROM seat_bookings WHERE seat_booking_id = ?;