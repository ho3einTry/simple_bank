package simpleBankDB

import (
	"context"
	"github.com/stretchr/testify/require"
	"simpleBank/util"
	"testing"
	"time"
)

func CreateRandomAccount(params ...CreateAccountParams) (Account, error) {
	var args CreateAccountParams
	if len(params) == 0 {
		args = CreateAccountParams{
			Owner:    util.RandomOwner(),
			Balance:  util.RandomMoney(),
			Currency: util.RandomCurrency(),
		}
	} else {
		args = params[0]
	}

	account, err := testQueries.CreateAccount(context.Background(), args)
	return account, err
}
func TestCreateAccount(t *testing.T) {
	param := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := CreateRandomAccount(param)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, param.Owner, account.Owner)
	require.Equal(t, param.Currency, account.Currency)
	require.Equal(t, param.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	createdAccount, creatingErr := CreateRandomAccount()
	require.NoError(t, creatingErr)

	account, err := testQueries.GetAccount(context.Background(), createdAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, createdAccount.ID, account.ID)
	require.Equal(t, createdAccount.Owner, account.Owner)
	require.Equal(t, createdAccount.Currency, account.Currency)
	require.Equal(t, createdAccount.Balance, account.Balance)
	require.WithinDuration(t, createdAccount.CreatedAt, account.CreatedAt, time.Second)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	createdAccount, creatingErr := CreateRandomAccount()
	require.NoError(t, creatingErr)

	args := UpdateAccountParams{
		ID:      createdAccount.ID,
		Balance: util.RandomMoney(),
	}
	account, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, createdAccount.ID, account.ID)
	require.Equal(t, createdAccount.Owner, account.Owner)
	require.Equal(t, args.Balance, account.Balance)
	require.Equal(t, createdAccount.Currency, account.Currency)
	require.WithinDuration(t, createdAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	createdAccount, creatingErr := CreateRandomAccount()
	require.NoError(t, creatingErr)

	err := testQueries.DeleteAccount(context.Background(), createdAccount.ID)
	require.NoError(t, err)

	_, err1 := testQueries.GetAccount(context.Background(), createdAccount.ID)
	require.EqualError(t, err1, "sql: no rows in result set")
}

func TestGetListAccount(t *testing.T) {
	var accounts []Account
	for i := 0; i < 5; i++ {
		createdAccount, _ := CreateRandomAccount()
		accounts = append(accounts, createdAccount)
	}

	listAccounts, _ := testQueries.ListAccounts(context.Background(), ListAccountsParams{
		Limit:  5,
		Offset: 5,
	})

	for _, account := range listAccounts {
		acc, err := testQueries.GetAccount(context.Background(), account.ID)
		require.NoError(t, err)
		require.NotEmpty(t, account)

		require.Equal(t, acc.ID, account.ID)
		require.Equal(t, acc.Owner, account.Owner)
		require.Equal(t, acc.Currency, account.Currency)
		require.Equal(t, acc.Balance, account.Balance)
		require.WithinDuration(t, acc.CreatedAt, account.CreatedAt, time.Second)

	}
}
