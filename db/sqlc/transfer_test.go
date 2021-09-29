package db

import (
	"context"
	"testing"
	"time"

	"github.com/jdiego0102/gobank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, account1, account2 Cuentum) Transferencium {
	arg := CreateTransferParams{
		FromCuentaID: account1.ID,
		ToCuentaID:   account2.ID,
		Monto:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromCuentaID, transfer.FromCuentaID)
	require.Equal(t, arg.ToCuentaID, transfer.ToCuentaID)
	require.Equal(t, arg.Monto, transfer.Monto)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromCuentaID, transfer2.FromCuentaID)
	require.Equal(t, transfer1.ToCuentaID, transfer2.ToCuentaID)
	require.Equal(t, transfer1.Monto, transfer2.Monto)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, account1, account2)
		createRandomTransfer(t, account2, account1)
	}

	arg := ListTransfersParams{
		FromCuentaID: account1.ID,
		ToCuentaID:   account1.ID,
		Limit:        5,
		Offset:       5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.FromCuentaID == account1.ID || transfer.ToCuentaID == account1.ID)
	}
}
