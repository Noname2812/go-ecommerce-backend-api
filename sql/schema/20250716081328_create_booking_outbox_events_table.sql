-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS booking_outbox_events (
    booking_outbox_event_id CHAR(36) PRIMARY KEY,               -- booking outbox event id (UUID)
    booking_id BIGINT NOT NULL,                                 -- booking related
    booking_outbox_event_type VARCHAR(100) NOT NULL,            -- BookingCreated, BookingCancelled
    booking_outbox_event_payload JSON NOT NULL,                 -- Data to publish (JSON)
    booking_outbox_event_status  TINYINT DEFAULT 1 NOT NULL,    -- 1: PENDING, 2: PUBLISHED, 3: FAILED
    booking_outbox_event_retries INT DEFAULT 0,                -- Number of retries
    booking_outbox_event_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation time (timestamp)
    booking_outbox_event_published_at TIMESTAMP NULL DEFAULT NULL,      -- Publication time (timestamp)
    booking_outbox_event_last_error TEXT DEFAULT NULL,                  -- Last error (text)
    -- Indexes
    INDEX idx_status (booking_outbox_event_status),                    -- Index for status
    INDEX idx_created_at (booking_outbox_event_created_at)             -- Index for creation time
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Outbox For Booking Service';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `booking outbox events`;
-- +goose StatementEnd
