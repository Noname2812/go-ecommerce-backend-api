-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
    payment_id BIGINT AUTO_INCREMENT PRIMARY KEY,           -- Payment ID

    booking_id BIGINT NOT NULL,                             -- Booking ID
    user_id BIGINT NOT NULL,                                -- User ID

    payment_amount DECIMAL(10, 2) NOT NULL DEFAULT 0.00,            -- Payment amount
    payment_currency VARCHAR(10) DEFAULT 'VND',                     -- Currency (VND, USD, EUR,…)
    payment_method TINYINT DEFAULT 0 NOT NULL,                     -- 0: CASH, 1: CREDIT_CARD, 2: VNPAY, 3: MOMO, 4: PAYPAL
    payment_provider_transaction_id VARCHAR(100) DEFAULT NULL,      -- Payment provider transaction ID

    payment_status TINYINT DEFAULT 1 NOT NULL,                     -- 1: PENDING, 2: SUCCESS, 3: FAILED, 4: REFUNDED

    payment_paid_at TIMESTAMP NULL DEFAULT NULL,                    -- Payment paid time
    payment_refunded_at TIMESTAMP NULL DEFAULT NULL,                -- Payment refunded time
    payment_expired_at TIMESTAMP NULL DEFAULT NULL,                 -- Payment expired time

    payment_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,        -- Creation time
    payment_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    payment_deleted_at TIMESTAMP NULL DEFAULT NULL,                -- Deletion time

    -- Index
    INDEX idx_booking_id (booking_id), -- Index for booking ID
    INDEX idx_user_id (user_id), -- Index for user ID
    INDEX idx_status (payment_status), -- Index for payment status
    INDEX idx_deleted_at (payment_deleted_at) -- Index for deleted time
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store payment information';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `payments`; -- Drop payments table
-- +goose StatementEnd
