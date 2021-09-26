-- name: CreateTransfer :one
INSERT INTO transferencia (
  from_cuenta_id,
  to_cuenta_id,
  monto
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transferencia
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transferencia
WHERE 
    from_cuenta_id = $1 OR
    to_cuenta_id = $2
ORDER BY id
LIMIT $3
OFFSET $4;