package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store proporciona todas las funciones para ejecutar consultas db individualmente.
type Store struct {
	*Queries // Composición
	db       *sql.DB
}

// NewStore crea un nuevo Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx ejecuta un función con una transacción db
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

// TransferTxParams contiene los parámetros de entrada de la transacción de la transferencia
type TransferTxParams struct {
	FromCuentaID int64 `json:"from_cuenta_id"`
	ToCuentaID   int64 `json:"to_cuenta_id"`
	Amount       int64 `json:"amount"`
}

// TransferTxResult es el resultado de la transacción de la transferencia
type TransferTxResult struct {
	Transfer    Transferencium `json:"transfer"`
	FromAccount Cuentum        `json:"from_account`
	ToAccount   Cuentum        `json:"to_account`
	FromEntry   Ingreso        `json:"from_entry`
	ToEntry     Ingreso        `json:"to_entry`
}

// TransferTx crea un registro de transferencia, agrega ingresos de cuenta
// y actualia el saldo de las cuentas dentro de una sola transacción de base de datos.
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromCuentaID: arg.FromCuentaID,
			ToCuentaID:   arg.ToCuentaID,
			Monto:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			CuentaID: arg.FromCuentaID,
			Monto:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			CuentaID: arg.ToCuentaID,
			Monto:    arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:    arg.FromCuentaID,
			Monto: -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:    arg.ToCuentaID,
			Monto: arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
