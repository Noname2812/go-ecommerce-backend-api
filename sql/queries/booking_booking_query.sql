-- name: GetBookingById :one
SELECT *
FROM `bookings`
WHERE booking_id = ?;

-- name: DeleteBooking :exec
UPDATE bookings
SET booking_deleted_at = NOW()
WHERE booking_id = ?;

-- name: AddBooking :execresult
INSERT INTO bookings (
    trip_id, user_id, booking_total_price, booking_status, booking_contact_name,
    booking_contact_phone, booking_contact_email, booking_note, booking_created_at, booking_updated_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: UpdateBooking :execrows
UPDATE bookings
SET booking_updated_at = NOW(), booking_total_price = ?, booking_status = ?, booking_contact_name = ?,
    booking_contact_phone = ?, booking_contact_email = ?, booking_note = ?
WHERE booking_id = ? AND booking_updated_at = ? AND booking_deleted_at IS NULL;

-- name: DeleteForceBooking :exec
DELETE FROM bookings WHERE booking_id = ?;