-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS transportation_outbox_events (
    transportation_outbox_event_id CHAR(36) PRIMARY KEY,                       -- Transportation outbox event ID (UUID)
    trip_id BIGINT NOT NULL,                                                   -- Trip related
    transportation_outbox_event_event_type VARCHAR(100) NOT NULL,              -- Event type
    transportation_outbox_event_payload JSON NOT NULL,                         -- Data to publish (JSON)
    transportation_outbox_event_status TINYINT DEFAULT 1 NOT NULL,             -- 1: PENDING, 2: PUBLISHED, 3: FAILED
    transportation_outbox_event_retries INT DEFAULT 0,                         -- Number of retries
    transportation_outbox_event_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation time (timestamp)
    transportation_outbox_event_published_at TIMESTAMP NULL DEFAULT NULL,       -- Publication time (timestamp)
    transportation_outbox_event_last_error TEXT DEFAULT NULL,                   -- Last error (text)
    -- Indexes
    INDEX idx_status (transportation_outbox_event_status), -- Index for status
    INDEX idx_created_at (transportation_outbox_event_created_at) -- Index for creation time       
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Outbox for Transportation Service';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `transportation_outbox_events`;
-- +goose StatementEnd
