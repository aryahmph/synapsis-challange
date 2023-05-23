-- name: CreateUser :one
INSERT INTO users (email,
                   password_hash,
                   name,
                   address)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: GetUser :one
SELECT id, email, password_hash, name, address
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, email, password_hash, name, address
FROM users
WHERE email = $1
LIMIT 1;