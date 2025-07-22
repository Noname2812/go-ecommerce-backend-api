-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employees (
    employee_id BIGINT AUTO_INCREMENT PRIMARY KEY,           -- Employee ID

    employee_name VARCHAR(100) NOT NULL,                     -- Employee name
    employee_email VARCHAR(100) UNIQUE DEFAULT NULL,         -- Email (can be NULL if not used)
    employee_phone VARCHAR(20) UNIQUE NOT NULL,              -- Phone number
    employee_gender TINYINT DEFAULT 0 NOT NULL,              -- 0: MALE, 1: FEMALE, 2: OTHER
    employee_birth_date DATE DEFAULT NULL,                   -- Birth date
    employee_base_salary DECIMAL(10,2) NOT NULL DEFAULT 0.0, -- Base salary

    employee_type TINYINT DEFAULT 0 NOT NULL,               -- 0: DRIVER, 1: ASSISTANT, 2: STAFF
    employee_license_number VARCHAR(50) DEFAULT NULL,       -- License number (only applies to DRIVER)
    employee_license_expiry_date DATE DEFAULT NULL,         -- License expiry date
    employee_status TINYINT DEFAULT 1 NOT NULL,             -- 1: ACTIVE, 2: INACTIVE

    employee_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,          -- Creation time
    employee_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Update time
    employee_deleted_at TIMESTAMP NULL DEFAULT NULL,                  -- Deletion time

    -- Indexes
    INDEX idx_employee_type (employee_type), -- Index for employee type
    INDEX idx_employee_name (employee_name), -- Index for employee name
    INDEX idx_deleted_at (employee_deleted_at) -- Index for deleted time
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Table to store employee information';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `employees`; -- Drop employees table
-- +goose StatementEnd
