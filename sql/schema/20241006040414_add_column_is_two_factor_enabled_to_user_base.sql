-- +goose Up
-- +goose StatementBegin
ALTER TABLE acc_user_base
ADD COLUMN is_two_factor_enabled INT(1) DEFAULT 0 COMMENT 'authentication is enabled for the user';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE acc_user_base
DROP COLUMN is_two_factor_enabled;
-- +goose StatementEnd
