-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS acc_user_info (
    user_id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT 'User ID', -- Primary key for user ID
    user_account VARCHAR(255) NOT NULL COMMENT 'User account', -- Account of the user
    user_nickname VARCHAR(255) COMMENT 'User nickname', -- Nickname of the user
    user_avatar VARCHAR(255) COMMENT 'User avatar', -- Avatar image URL for the user
    user_state TINYINT UNSIGNED NOT NULL COMMENT 'User state: 0-Locked, 1-Activated, 2-Not Activated', -- User state (enum)
    user_phone VARCHAR(20) COMMENT 'Mobile phone number', -- User's mobile phone number

    user_gender TINYINT UNSIGNED COMMENT 'User gender: 0-Secret, 1-Male, 2-Female', -- Gender (enum)
    user_birthday DATE COMMENT 'User birthday', -- Date of birth
    user_address VARCHAR(255) COMMENT 'User address', -- Address of the user

    user_is_authentication TINYINT UNSIGNED NOT NULL COMMENT 'Authentication status: 0-Not Authenticated, 1-Pending, 2-Authenticated, 3-Failed', -- Authentication status (enum)

    -- Add timestamps for record creation and updates
    user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation time', -- Time when the record was created
    user_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update time', -- Time when the record was last updated
    user_deleted_at TIMESTAMP NULL DEFAULT NULL COMMENT 'Record deletion time', -- Time when the record was deleted

    -- Indexes for optimized querying
    UNIQUE KEY unique_user_account (user_account), -- Ensure user_account is unique
    INDEX idx_user_phone (user_phone), -- Index for querying by user_phone
    INDEX idx_user_state (user_state), -- Index for querying by user_state
    INDEX idx_user_is_authentication (user_is_authentication) -- Index for querying by authentication status

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='acc_user_info';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `acc_user_info`;
-- +goose StatementEnd
