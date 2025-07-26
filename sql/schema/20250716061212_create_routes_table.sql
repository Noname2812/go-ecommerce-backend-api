-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS routes (
    route_id INT AUTO_INCREMENT PRIMARY KEY,                           -- Route ID
    
    route_start_location VARCHAR(100) NOT NULL,                        -- Start location
    route_end_location VARCHAR(100) NOT NULL,                          -- End location
    route_estimated_duration INT NOT NULL,                             -- Estimated duration (minutes)
    
    route_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,              -- Record creation time
    route_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Record update time
    route_deleted_at TIMESTAMP NULL DEFAULT NULL,                      -- Record deletion time

    -- Indexes
    INDEX idx_route_deleted_at (route_deleted_at), -- Index for deleted time

    -- Ensure uniqueness on start+end location
    UNIQUE KEY unique_route_locations (route_start_location, route_end_location) -- Unique constraint for start and end location
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store route information';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `routes`; -- Drop routes table         
-- +goose StatementEnd
