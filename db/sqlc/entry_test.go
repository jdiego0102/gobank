package db

import (
	"context"
	"testing"
	"time"

	"github.com/jdiego0102/gobank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, account Cuentum) Ingreso {
	arg := CreateEntryParams{
		CuentaID: account.ID,
		Monto:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, arg.CuentaID, entry.CuentaID)
	require.Equal(t, arg.Monto, entry.Monto)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.CuentaID, entry2.CuentaID)
	require.Equal(t, entry1.Monto, entry2.Monto)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	arg := ListEntriesParams{
		CuentaID: account.ID,
		Limit:    5,
		Offset:   5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.CuentaID, entry.CuentaID)
	}
}
