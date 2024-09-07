-- name: GetUserByName :one
SELECT * FROM users
WHERE name = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByApiKey :one
SELECT * FROM users
WHERE api_key = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
    id,
    created_at,
    updated_at,
    name,
    api_key
) VALUES (
    $1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex')
) RETURNING *;