-- name: GetAllUsers :many
SELECT id, name, email
FROM users;

-- name: CreateUser :one
INSERT INTO users (name, email)
VALUES ($1, $2) RETURNING id;

-- name: GetUserByID :one
SELECT id, name, email
FROM users
WHERE id = $1;
