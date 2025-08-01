-- name: GetOneUserInfo :one
SELECT user_id, user_account, user_password, user_salt
FROM `acc_user_base`
WHERE user_account = ?;

-- name: GetOneUserInfoAdmin :one
SELECT user_id, user_account, user_password, user_salt, user_login_time, user_logout_time, user_login_ip
        ,user_created_at, user_updated_at
FROM `acc_user_base`
WHERE user_account = ?;

-- name: CheckUserBaseExists :one
SELECT COUNT(*)
FROM acc_user_base
WHERE user_account = ?;

-- name: AddUserBase :execresult
INSERT INTO acc_user_base (
    user_account, user_password, user_salt, user_created_at, user_updated_at
) VALUES (
    ?, ?, ?, NOW(), NOW()
);

-- name: LoginUserBase :exec
UPDATE acc_user_base
SET user_login_time = NOW(), user_login_ip = ?
WHERE user_account = ? AND user_password = ?;

-- name: LogoutUserBase :exec
UPDATE acc_user_base
SET user_logout_time = NOW()
WHERE user_account = ?;

