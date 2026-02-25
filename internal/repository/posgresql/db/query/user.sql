-- name: CreateUser :one
INSERT INTO users (username, name, email, phone_number, password_hash, role, data)
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;

-- name: GetUserByUsername :one
SELECT *
FROM users
where username = $1;
