-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS trip_seat_locks (
    trip_seat_lock_id BIGINT AUTO_INCREMENT PRIMARY KEY,  -- Seat lock ID
    trip_id BIGINT NOT NULL,                             -- Trip ID
    trip_seat_lock_seat_number VARCHAR(10) NOT NULL,     -- Seat number

    locked_by_booking_id BIGINT NULL,                -- Booking ID that locked the seat
    
    trip_seat_lock_status TINYINT DEFAULT 1 NOT NULL, -- 1: AVAILABLE, 2: LOCKED, 3: BOOKED
    trip_seat_lock_expires_at TIMESTAMP NULL,         -- Seat lock expiration time
    
    trip_seat_lock_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation time
    trip_seat_lock_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    trip_seat_lock_deleted_at TIMESTAMP NULL DEFAULT NULL, -- Deletion time

    -- Ensure each seat has only one record per trip
    UNIQUE KEY unique_trip_seat (trip_id, trip_seat_lock_seat_number), -- Unique constraint for trip and seat number
    
    -- Indexes
    INDEX idx_trip_id (trip_id), -- Index for trip ID
    INDEX idx_lock_status (trip_seat_lock_status) -- Index for lock status
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to control seat status';
-- +goose StatementEnd