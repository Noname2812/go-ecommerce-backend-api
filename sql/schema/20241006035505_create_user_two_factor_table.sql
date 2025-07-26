-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `acc_user_two_factor` (
    `two_factor_id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,        -- Primary key
    `user_id` INT UNSIGNED NOT NULL,                                -- Foreign key linked to user table
    `two_factor_auth_type` ENUM('SMS', 'EMAIL', 'APP') NOT NULL,    -- 2FA method (SMS, Email, App like Google Authenticator)
    `two_factor_auth_secret` VARCHAR(255) NOT NULL,                 -- Secret information for 2FA (e.g. TOTP secret for 2FA app) 
    `two_factor_phone` VARCHAR(20) NULL,                            -- Phone number for 2FA via SMS (if applicable)
    `two_factor_email` VARCHAR(255) NULL,                           -- Email address for 2FA via Email (if applicable)
    `two_factor_is_active` BOOLEAN NOT NULL DEFAULT TRUE,           -- Status of 2FA method
    `two_factor_created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    -- Creation time
    `two_factor_updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time

    -- Indexes
    INDEX `idx_user_id` (`user_id`),               -- Index for user ID
    INDEX `idx_auth_type` (`two_factor_auth_type`) -- Index for auth type
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store user two factor information';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `acc_user_two_factor`; -- Drop user two factor table       
-- +goose StatementEnd
