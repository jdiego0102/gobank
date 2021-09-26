-- name: CreateEntry :one
INSERT INTO ingreso (
  cuenta_id,
  monto
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetEntry :one
SELECT * FROM ingreso
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM ingreso
WHERE cuenta_id = $1
ORDER BY id
LIMIT $2
OFFSET $3; 