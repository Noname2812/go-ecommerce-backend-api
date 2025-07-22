-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS bookings (
    booking_id BIGINT AUTO_INCREMENT PRIMARY KEY,          -- Booking ID
    trip_id BIGINT NOT NULL,                               -- Trip ID
    user_id BIGINT NULL,                                   -- User ID

    booking_total_price DECIMAL(10, 2) NOT NULL DEFAULT 0.00,      -- Total price
    booking_status TINYINT DEFAULT 1 NOT NULL,                  -- 1: PENDING, 2: BOOKED, 3: CANCELLED, 4: COMPLETED, 5: EXPIRED, 6: REFUNDED

    booking_contact_name VARCHAR(100) NOT NULL, -- Contact name
    booking_contact_phone VARCHAR(20) NOT NULL, -- Contact phone
    booking_contact_email VARCHAR(100) NOT NULL, -- Contact email
    
    booking_note TEXT NULL, -- Note for booking

    booking_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,        -- Creation time
    booking_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    booking_deleted_at TIMESTAMP NULL DEFAULT NULL,                -- Deletion time

    -- Indexes
    INDEX idx_trip_user (trip_id, user_id), -- Index for trip and user
    INDEX idx_status (booking_status) -- Index for booking status
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store booking information';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `bookings`; -- Drop bookings table
-- +goose StatementEnd
