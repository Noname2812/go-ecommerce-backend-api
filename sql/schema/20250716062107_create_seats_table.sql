-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS seats (
    seat_id INT AUTO_INCREMENT PRIMARY KEY,               -- Seat ID
    bus_id INT NOT NULL,                                  -- Bus ID
    seat_number VARCHAR(10) NOT NULL,                     -- Seat number (A1, B2, ...)
    seat_row_no TINYINT UNSIGNED NOT NULL,                         -- Seat row number (1, 2, 3, ...)
    seat_column_no TINYINT UNSIGNED NOT NULL,                      -- Seat column number (1, 2, 3, ...)
    seat_floor_no TINYINT UNSIGNED NOT NULL,                       -- Seat floor number (1, 2, 3, ...)
    seat_type TINYINT UNSIGNED NOT NULL DEFAULT 1,                 -- Seat type (1: normal, 2: VIP, 3: wheelchair)

    seat_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Record creation time
    seat_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Record update time
    seat_deleted_at TIMESTAMP NULL DEFAULT NULL,          -- Record deletion time

    UNIQUE KEY unique_bus_seat (bus_id, seat_number) -- Unique constraint for bus and seat number
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store seat information';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `seats`; -- Drop seats table   
-- +goose StatementEnd
