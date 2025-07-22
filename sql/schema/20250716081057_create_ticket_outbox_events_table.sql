-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS ticket_outbox_events (
    ticket_outbox_event_id CHAR(36) PRIMARY KEY,                        -- Ticket outbox event ID (UUID)
    ticket_id BIGINT NOT NULL,                                          -- ticket related
    ticket_outbox_event_event_type VARCHAR(100) NOT NULL,               -- TicketRefunded, TicketCancelled, TicketCreated, ....
    ticket_outbox_event_payload JSON NOT NULL,                          -- Data to publish (JSON)
    ticket_outbox_event_status TINYINT DEFAULT 1 NOT NULL,              -- 1: PENDING, 2: PUBLISHED, 3: FAILED
    ticket_outbox_event_retries INT DEFAULT 0,                          -- Number of retries
    ticket_outbox_event_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Creation time (timestamp)
    ticket_outbox_event_published_at TIMESTAMP NULL DEFAULT NULL,       -- Publication time (timestamp)
    ticket_outbox_event_last_error TEXT DEFAULT NULL,                   -- Last error (text)
    -- Indexes
    INDEX idx_status (ticket_outbox_event_status),                      -- Index for status
    INDEX idx_created_at (ticket_outbox_event_created_at)              -- Index for creation time
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Outbox for Ticket Service';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `ticket outbox events`;
-- +goose StatementEnd
