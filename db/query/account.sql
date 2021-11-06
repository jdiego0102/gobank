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

-- name: GetAccountForUpdate :one
SELECT * FROM cuenta
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM cuenta
ORDER BY id
LIMIT $1
OFFSET $2; 

-- name: UpdateAccount :one
UPDATE cuenta SET tope = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE cuenta 
SET tope = tope + sqlc.arg(monto)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM cuenta WHERE id = $1;