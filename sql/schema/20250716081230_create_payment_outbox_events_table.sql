-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payment_outbox_events (
    payment_outbox_event_id CHAR(36) PRIMARY KEY,               -- payment outbox event id (UUID)
    payment_id BIGINT NOT NULL,                                 -- payment related
    payment_outbox_event_type VARCHAR(100) NOT NULL,            -- PaymentCompleted, PaymentFailed
    payment_outbox_event_payload JSON NOT NULL,                 -- Data to publish (JSON)
    payment_outbox_event_status TINYINT DEFAULT 1 NOT NULL,     -- 1: PENDING, 2: PUBLISHED, 3: FAILED
    payment_outbox_event_retries INT DEFAULT 0,                 -- Number of retries
    payment_outbox_event_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation time (timestamp)
    payment_outbox_event_published_at TIMESTAMP NULL DEFAULT NULL,      -- Publication time (timestamp)
    payment_outbox_event_last_error TEXT DEFAULT NULL,                  -- Last error (text)
    -- Indexes
    INDEX idx_status (payment_outbox_event_status),                    -- Index for status
    INDEX idx_created_at (payment_outbox_event_created_at)             -- Index for creation time
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Outbox for Payment Service';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `payment outbox events`;
-- +goose StatementEnd
