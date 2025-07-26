-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS seats (
    seat_id INT AUTO_INCREMENT PRIMARY KEY,               -- Seat ID
    bus_id INT NOT NULL,                                  -- Bus ID
    seat_number VARCHAR(10) NOT NULL,                     -- Seat number (A1, B2, ...)
    is_available BOOLEAN DEFAULT TRUE,                    -- Seat status

    seat_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Record creation time
    seat_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Record update time
    seat_deleted_at TIMESTAMP NULL DEFAULT NULL,          -- Record deletion time

    -- Indexes
    INDEX idx_seat_deleted_at (seat_deleted_at),     -- Index for deleted time
    UNIQUE KEY unique_bus_seat (bus_id, seat_number) -- Unique constraint for bus and seat number
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store seat information';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `seats`; -- Drop seats table   
-- +goose StatementEnd
