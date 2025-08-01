-- name: GetValidOTP :one
SELECT verify_otp, verify_key_hash, verify_key, verify_id
FROM `acc_user_verify`
WHERE verify_key_hash = ? AND is_verified = 0;

-- name: UpdateUserVerificationStatus :exec
UPDATE `acc_user_verify`
SET is_verified = 1,
    verify_updated_at = now()
WHERE verify_key_hash = ?;

-- name: InsertOTPVerify :execresult
INSERT INTO `acc_user_verify` (
    verify_otp,
    verify_key,
    verify_key_hash, 
    verify_type, 
    is_verified, 
    is_deleted, 
    verify_created_at, 
    verify_updated_at
)
VALUES (?, ?, ?, ?, 0, 0, NOW(), NOW());

-- name: GetInfoOTP :one
SELECT verify_id, verify_otp, verify_key, verify_key_hash, verify_type, is_verified, is_deleted, verify_created_at, verify_updated_at
FROM `acc_user_verify`
WHERE verify_key_hash = ?;
