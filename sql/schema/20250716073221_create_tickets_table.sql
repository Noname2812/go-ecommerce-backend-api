-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tickets (
    ticket_id BIGINT AUTO_INCREMENT PRIMARY KEY,           -- Ticket ID

    -- Denormalized data (do Data per Service)
    user_id BIGINT NOT NULL,                               -- User ID
    trip_id BIGINT NOT NULL,                               -- Trip ID
    seat_booking_id BIGINT NOT NULL,                       -- Seat booking ID

    ticket_code VARCHAR(50) NOT NULL UNIQUE,               -- Ticket code (QR code, barcode,…)
    ticket_status TINYINT DEFAULT 1 NOT NULL,              -- 1: ACTIVE, 2: USED, 3: REFUNDED
    ticket_issued_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Ticket issued time

    -- User info (denormalized)
    user_nickname VARCHAR(100) DEFAULT NULL,               -- User nickname
    user_gender TINYINT DEFAULT 0 NOT NULL,                -- 0: MALE, 1: FEMALE, 2: OTHER

    -- Seat_booking info (denormalized)
    seat_booking_price DECIMAL(10, 2) DEFAULT 0.00,        -- Seat booking price
    seat_booking_seat_number VARCHAR(10) NOT NULL,         -- Seat number (A1, B2,…)

    -- Trip info (denormalized)
    bus_license_plate VARCHAR(20) DEFAULT NULL,            -- Bus license plate
    trip_departure_time DATETIME DEFAULT NULL,             -- Trip departure time
    trip_arrival_time DATETIME DEFAULT NULL,               -- Trip arrival time
    route_start_location VARCHAR(100) DEFAULT NULL,        -- Route start location
    route_end_location VARCHAR(100) DEFAULT NULL,          -- Route end location

    -- Soft delete
    ticket_booking_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,        -- Ticket booking created time
    ticket_booking_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Ticket booking updated time
    ticket_booking_deleted_at TIMESTAMP NULL DEFAULT NULL,                -- Ticket booking deleted time

    -- Index
    INDEX idx_seat_booking (seat_booking_id), -- Index for seat booking ID
    INDEX idx_ticket_status (ticket_status), -- Index for ticket status
    INDEX idx_user_trip (user_id, trip_id), -- Index for user and trip
    INDEX idx_deleted_at (ticket_booking_deleted_at) -- Index for deleted time
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store ticket information';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `tickets`;
-- +goose StatementEnd
