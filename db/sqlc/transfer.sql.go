// Code generated by sqlc. DO NOT EDIT.
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transferencia (
  from_cuenta_id,
  to_cuenta_id,
  monto
) VALUES (
  $1, $2, $3
) RETURNING id, from_cuenta_id, to_cuenta_id, monto, created_at
`

type CreateTransferParams struct {
	FromCuentaID int64 `json:"from_cuenta_id"`
	ToCuentaID   int64 `json:"to_cuenta_id"`
	Monto        int64 `json:"monto"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transferencium, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromCuentaID, arg.ToCuentaID, arg.Monto)
	var i Transferencium
	err := row.Scan(
		&i.ID,
		&i.FromCuentaID,
		&i.ToCuentaID,
		&i.Monto,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_cuenta_id, to_cuenta_id, monto, created_at FROM transferencia
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transferencium, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transferencium
	err := row.Scan(
		&i.ID,
		&i.FromCuentaID,
		&i.ToCuentaID,
		&i.Monto,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, from_cuenta_id, to_cuenta_id, monto, created_at FROM transferencia
WHERE 
    from_cuenta_id = $1 OR
    to_cuenta_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTransfersParams struct {
	FromCuentaID int64 `json:"from_cuenta_id"`
	ToCuentaID   int64 `json:"to_cuenta_id"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transferencium, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers,
		arg.FromCuentaID,
		arg.ToCuentaID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transferencium
	for rows.Next() {
		var i Transferencium
		if err := rows.Scan(
			&i.ID,
			&i.FromCuentaID,
			&i.ToCuentaID,
			&i.Monto,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}