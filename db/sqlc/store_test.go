package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before: ", account1.Tope, account2.Tope)

	// Correr n concurrencias de transacción de transferencia
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromCuentaID: account1.ID,
				ToCuentaID:   account2.ID,
				Amount:       amount,
			})

			errs <- err
			results <- result
		}()
	}

	// Comprobar resultados
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// Verificar transferencia
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromCuentaID)
		require.Equal(t, account2.ID, transfer.ToCuentaID)
		require.Equal(t, amount, transfer.Monto)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// Verificar el ingreso de la trasnferencia
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.CuentaID)
		require.Equal(t, -amount, fromEntry.Monto)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.CuentaID)
		require.Equal(t, amount, toEntry.Monto)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// Verificar cuentas
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// Verificar tope de las cuentas
		fmt.Println(">> tx: ", fromAccount.Tope, toAccount.Tope)
		diff1 := account1.Tope - fromAccount.Tope
		diff2 := toAccount.Tope - account2.Tope
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0) // 1 * monto, 2 * monto, 3 * monto, ..., n * monto

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after: ", updatedAccount1.Tope, updatedAccount2.Tope)
	require.Equal(t, account1.Tope-int64(n)*amount, updatedAccount1.Tope)
	require.Equal(t, account2.Tope+int64(n)*amount, updatedAccount2.Tope)
}
