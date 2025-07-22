-- name: EnableTwoFactorTypeEmail :exec
INSERT INTO acc_user_two_factor (user_id, two_factor_auth_type, two_factor_email, two_factor_auth_secret, two_factor_is_active, two_factor_created_at, two_factor_updated_at)
VALUES (?, ?, ?, "OTP", FALSE, NOW(), NOW());

-- name: DisableTwoFactor :exec
UPDATE acc_user_two_factor
SET two_factor_is_active = FALSE, 
    two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ?;

-- name: UpdateTwoFactorStatus :exec
UPDATE acc_user_two_factor
SET two_factor_is_active = TRUE, two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ? AND two_factor_is_active = FALSE;

-- name: VerifyTwoFactor :one
SELECT COUNT(*)
FROM acc_user_two_factor
WHERE user_id = ? AND two_factor_auth_type = ? AND two_factor_is_active = TRUE;

-- name: GetTwoFactorStatus :one
SELECT two_factor_is_active
FROM acc_user_two_factor
WHERE user_id = ? AND two_factor_auth_type = ?;

-- name: IsTwoFactorEnabled :one
SELECT COUNT(*)
FROM acc_user_two_factor
WHERE user_id = ? AND two_factor_is_active = TRUE;

-- name: AddOrUpdatePhoneNumber :exec
INSERT INTO acc_user_two_factor (user_id, two_factor_phone, two_factor_is_active)
VALUES (?, ?, TRUE)
ON DUPLICATE KEY UPDATE 
    two_factor_phone = ?, 
    two_factor_updated_at = NOW();

-- name: AddOrUpdateEmail :exec
INSERT INTO acc_user_two_factor (user_id, two_factor_email, two_factor_is_active)
VALUES (?, ?, TRUE)
ON DUPLICATE KEY UPDATE 
    two_factor_email = ?, 
    two_factor_updated_at = NOW();

-- name: GetUserTwoFactorMethods :many
SELECT two_factor_id, user_id, two_factor_auth_type, two_factor_auth_secret, 
       two_factor_phone, two_factor_email, 
       two_factor_is_active, two_factor_created_at, two_factor_updated_at
FROM acc_user_two_factor
WHERE user_id = ?;

-- name: ReactivateTwoFactor :exec
UPDATE acc_user_two_factor
SET two_factor_is_active = TRUE, 
    two_factor_updated_at = NOW()
WHERE user_id = ? AND two_factor_auth_type = ?;

-- name: RemoveTwoFactor :exec
DELETE FROM acc_user_two_factor
WHERE user_id = ? AND two_factor_auth_type = ?;

-- name: CountActiveTwoFactorMethods :one
SELECT COUNT(*)
FROM acc_user_two_factor
WHERE user_id = ? AND two_factor_is_active = TRUE;

-- name: GetTwoFactorMethodByID :one
SELECT two_factor_id, user_id, two_factor_auth_type, two_factor_auth_secret, 
       two_factor_phone, two_factor_email, 
       two_factor_is_active, two_factor_created_at, two_factor_updated_at
FROM acc_user_two_factor
WHERE two_factor_id = ?;

-- name: GetTwoFactorMethodByIDAndType :one
SELECT two_factor_id, user_id, two_factor_auth_type, two_factor_auth_secret, 
       two_factor_phone, two_factor_email, 
       two_factor_is_active, two_factor_created_at, two_factor_updated_at
FROM acc_user_two_factor
WHERE user_id = ? AND two_factor_auth_type = ?;
