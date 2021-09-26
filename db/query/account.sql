-- name: CreateAccount :one
INSERT INTO cuenta (
  propietario,
  tope,
  divisa
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM cuenta
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM cuenta
ORDER BY id
LIMIT $1
OFFSET $2; 

-- name: UpdateAccount :one
UPDATE cuenta SET tope = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM cuenta WHERE id = $1;