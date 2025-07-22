-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS notification_logs (
    notification_id BIGINT AUTO_INCREMENT PRIMARY KEY,         -- Notification ID
    user_id BIGINT NOT NULL,                                   -- User ID
    notification_type VARCHAR(100) NOT NULL,                   -- Notification type
    notification_recipient VARCHAR(255) NOT NULL,              -- Notification recipient
    notification_subject VARCHAR(255) DEFAULT NULL,            -- Subject (if email)
    notification_content TEXT NOT NULL,                        -- Notification content
    notification_status TINYINT DEFAULT 1 NOT NULL,            -- 1: PENDING, 2: SENT, 3: FAILED
    notification_error_message TEXT DEFAULT NULL,                           -- Error message if failed
    notification_sent_at TIMESTAMP NULL DEFAULT NULL,                       -- Sent time
    notification_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,            -- Creation time
    notification_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    notification_deleted_at TIMESTAMP NULL DEFAULT NULL,                    -- Deletion time

    -- Indexes
    INDEX idx_user_id (user_id), -- Index for user ID
    INDEX idx_type_status (notification_type, notification_status), -- Index for type and status
    INDEX idx_sent_at (notification_sent_at) -- Index for sent time
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='notification_logs_table';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `notification logs`;
-- +goose StatementEnd
