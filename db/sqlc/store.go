package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// SQLStore proporciona todas las funciones para ejecutar SQL db individualmente.
type SQLStore struct {
	db       *sql.DB
	*Queries // Composición
}

// NewStore crea un nuevo Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx ejecuta un función con una transacción db
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
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
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
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

		if arg.FromCuentaID < arg.ToCuentaID {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromCuentaID, -arg.Amount, arg.ToCuentaID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToCuentaID, arg.Amount, arg.FromCuentaID, -arg.Amount)
		}

		return nil
	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Cuentum, account2 Cuentum, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:    accountID1,
		Monto: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:    accountID2,
		Monto: amount2,
	})
	return
}
