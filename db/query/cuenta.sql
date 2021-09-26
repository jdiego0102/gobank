-- name: CreateAccount :one
INSERT INTO cuenta (
  propietario,
  tope,
  divisa
) VALUES (
  $1, $2, $3
)
RETURNING *;