package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jdiego0102/gobank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Cuentum {
	user := createRandomUser(t)

	arg := CreateAccountParams{
		Propietario: user.Username,
		Tope:        util.RandomMoney(),
		Divisa:      util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Propietario, account.Propietario)
	require.Equal(t, arg.Tope, account.Tope)
	require.Equal(t, arg.Divisa, account.Divisa)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Propietario, account2.Propietario)
	require.Equal(t, account1.Tope, account2.Tope)
	require.Equal(t, account1.Divisa, account2.Divisa)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:   account1.ID,
		Tope: util.RandomMoney(),
	}
	account2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Propietario, account2.Propietario)
	require.Equal(t, arg.Tope, account2.Tope)
	require.Equal(t, account1.Divisa, account2.Divisa)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Cuentum
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Propietario: lastAccount.Propietario,
		Limit:       5,
		Offset:      0,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Propietario, account.Propietario)
	}
}
