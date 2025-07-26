-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS buses (
    bus_id INT AUTO_INCREMENT PRIMARY KEY,                           -- Bus ID
    bus_license_plate VARCHAR(20) NOT NULL UNIQUE,                   -- Bus license plate (unique)

    bus_company VARCHAR(100) NOT NULL,                               -- Bus company
    bus_price DECIMAL(10,2) NOT NULL DEFAULT 0.0,                    -- Default price for bus
    bus_capacity INT NOT NULL,                                       -- Bus capacity
    
    bus_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,              -- Record creation time
    bus_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Record update time
    bus_deleted_at TIMESTAMP NULL DEFAULT NULL,                      -- Record deletion time

    -- Indexes
    INDEX idx_bus_deleted_at (bus_deleted_at), -- Index for deleted time

    -- Ensure uniqueness on license plate
    UNIQUE KEY unique_bus_license (bus_license_plate) -- Unique constraint for bus license plate
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store bus information';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `buses`; -- Drop buses table   
-- +goose StatementEnd
