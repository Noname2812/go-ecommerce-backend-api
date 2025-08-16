-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS seat_bookings (
    seat_booking_id BIGINT AUTO_INCREMENT PRIMARY KEY,     -- Seat booking ID
    booking_id CHAR(36) NOT NULL,                            -- Booking ID
    
    seat_booking_seat_number VARCHAR(10) NOT NULL,                      -- Seat number (A1, A2,...)
    seat_booking_price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,            -- Seat price

    passenger_name VARCHAR(100) NOT NULL, -- Passenger name
    passenger_phone VARCHAR(20) NOT NULL, -- Passenger phone

    seat_booking_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,        -- Creation time
    seat_booking_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    seat_booking_deleted_at TIMESTAMP NULL DEFAULT NULL,                -- Deletion time

    -- Ensure no duplicate seats in the same booking
    UNIQUE KEY unique_booking_seat (booking_id, seat_booking_seat_number), -- Unique constraint for booking and seat number

    -- Index for filtering quickly
    INDEX idx_booking_id (booking_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store seat booking information';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `seat_bookings`; -- Drop seat bookings table
-- +goose StatementEnd
