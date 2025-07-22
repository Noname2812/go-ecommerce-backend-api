-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS trips (
    trip_id BIGINT AUTO_INCREMENT PRIMARY KEY,                          -- Trip ID

    route_id BIGINT NOT NULL,                                           -- Route ID (denormalized from Route Service)
    bus_id BIGINT NOT NULL,                                             -- Bus ID (denormalized from Bus Service)
    
    trip_departure_time DATETIME NOT NULL,                              -- Departure time
    trip_arrival_time DATETIME NOT NULL,                                -- Arrival time
    trip_base_price DECIMAL(10,2) NOT NULL DEFAULT 0.0,                 -- Base price

    trip_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,                -- Creation time
    trip_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    trip_deleted_at TIMESTAMP NULL DEFAULT NULL,                        -- Deletion time

    -- Indexes
    INDEX idx_trip_departure_time (trip_departure_time), -- Index for departure time
    INDEX idx_trip_deleted_at (trip_deleted_at), -- Index for deleted time

    -- Ensure uniqueness: one bus cannot run two trips at the same time
    UNIQUE KEY unique_trip_schedule (bus_id, trip_departure_time) -- Unique constraint for bus and departure time
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store trip information';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `trips`; -- Drop trips table
-- +goose StatementEnd
