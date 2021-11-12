-- name: CreateUser :one
INSERT INTO users (
  username,
  password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetAUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;