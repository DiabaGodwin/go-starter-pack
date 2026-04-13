-- =========================================
-- USER PROFILES
-- =========================================

-- name: GetAllUserProfiles :many
SELECT
    id,
    user_id,
    first_name,
    last_name,
    display_name,
    phone,
    avatar_url,
    bio,
    date_of_birth,
    gender,
    country,
    city,
    address,
    created_at,
    updated_at
FROM user_profiles
ORDER BY created_at DESC, id DESC;


-- name: GetUserProfileByID :one
SELECT
    id,
    user_id,
    first_name,
    last_name,
    display_name,
    phone,
    avatar_url,
    bio,
    date_of_birth,
    gender,
    country,
    city,
    address,
    created_at,
    updated_at
FROM user_profiles
WHERE id = $1;


-- name: GetUserProfileByUserID :one
SELECT
    id,
    user_id,
    first_name,
    last_name,
    display_name,
    phone,
    avatar_url,
    bio,
    date_of_birth,
    gender,
    country,
    city,
    address,
    created_at,
    updated_at
FROM user_profiles
WHERE user_id = $1;


-- name: CreateUserProfile :one
INSERT INTO user_profiles (
    user_id,
    first_name,
    last_name,
    display_name,
    phone,
    avatar_url,
    bio,
    date_of_birth,
    gender,
    country,
    city,
    address
)
VALUES (
           $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
       )
    RETURNING
    id,
    user_id,
    first_name,
    last_name,
    display_name,
    phone,
    avatar_url,
    bio,
    date_of_birth,
    gender,
    country,
    city,
    address,
    created_at,
    updated_at;


-- name: UpdateUserProfile :one
UPDATE user_profiles
SET
    first_name = $2,
    last_name = $3,
    display_name = $4,
    phone = $5,
    avatar_url = $6,
    bio = $7,
    date_of_birth = $8,
    gender = $9,
    country = $10,
    city = $11,
    address = $12,
    updated_at = NOW()
WHERE user_id = $1
    RETURNING
    id,
    user_id,
    first_name,
    last_name,
    display_name,
    phone,
    avatar_url,
    bio,
    date_of_birth,
    gender,
    country,
    city,
    address,
    created_at,
    updated_at;


-- name: DeleteUserProfile :exec
DELETE FROM user_profiles
WHERE user_id = $1;