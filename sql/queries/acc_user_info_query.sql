-- name: GetUserByUserId :one
SELECT 
    user_id,
    user_account,
    user_nickname,
    user_avatar,
    user_state,
    user_phone,
    user_gender,
    user_birthday,
    user_is_authentication,
    user_created_at,
    user_updated_at
FROM `acc_user_info`
WHERE user_id = ? LIMIT 1;

-- name: GetUsers :many
SELECT user_id,
    user_account,
    user_nickname,
    user_avatar,
    user_state,
    user_phone,
    user_gender,
    user_birthday,
    user_is_authentication,
    user_created_at,
    user_updated_at 
FROM acc_user_info 
WHERE user_id IN (?);

-- name: FindUsers :many
SELECT * FROM acc_user_info WHERE user_account LIKE ? OR user_nickname LIKE ?;

-- name: ListUsers :many
SELECT * FROM acc_user_info LIMIT ? OFFSET ?;


-- name: RemoveUser :exec
DELETE FROM acc_user_info WHERE user_id = ?;

-- name: DeleteForceUser :exec
DELETE FROM acc_user_info WHERE user_account = ?;

-- -- name: UpdatePassword :exec
-- UPDATE `acc_user_info` SET user_password = ? WHERE user_id = ?;



-- name: AddUserAutoUserId :execresult
INSERT INTO `acc_user_info` (
    user_account, user_nickname, user_avatar, user_state, user_phone, 
    user_gender, user_birthday, user_is_authentication
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: AddUserHaveUserId :execresult
INSERT INTO `acc_user_info` (
    user_id, user_account, user_nickname, user_avatar, user_state, user_phone, 
    user_gender, user_birthday,  user_is_authentication
)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: EditUserByUserId :execresult
UPDATE `acc_user_info`
SET user_nickname = ?, user_avatar = ?, user_phone = ?, 
user_gender = ?, user_birthday = ?, user_updated_at = NOW()
WHERE user_id = ? AND user_is_authentication = 1;
