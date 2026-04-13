-- name: GetUserByID :one
SELECT
    id,
    email,
    role,
    status,
    password_hash,
    email_verified_at,
    last_login_at,
    created_at,
    updated_at
FROM users
WHERE id = $1
ORDER BY created_at DESC, id DESC;

-- name: CreateUser :one
INSERT INTO users (
    email,
    password_hash,
    role,
    status
)
VALUES ($1, $2, $3, $4)
    RETURNING
    id,
    email,
    role,
    status,
    password_hash,
    email_verified_at,
    last_login_at,
    created_at,
    updated_at;


-- name: GetUserByEmail :one
SELECT
    id,
    email,
    password_hash,
    role,
    status,
    email_verified_at,
    last_login_at,
    created_at,
    updated_at
FROM users
WHERE email = $1
  AND deleted_at IS NULL;


-- name: GetUserWithProfile :one
SELECT
    u.id,
    u.email,
    u.role,
    u.status,
    u.created_at,

    p.first_name,
    p.last_name,
    p.display_name,
    p.avatar_url
FROM users u
         LEFT JOIN user_profiles p
                   ON p.user_id = u.id
WHERE u.id = $1
  AND u.deleted_at IS NULL;


-- name: GetUsers :many
SELECT
    id,
    email,
    role,
    status,
    password_hash,
    email_verified_at,
    last_login_at,
    created_at,
    updated_at
FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC, id DESC
LIMIT $1 OFFSET $2;

-- name: EmailExists :one
SELECT id,email,role
FROM users
WHERE email =  $1;
