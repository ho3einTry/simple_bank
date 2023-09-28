package simpleBankDB

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_TransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1, _ := CreateRandomAccount()
	account2, _ := CreateRandomAccount()

	// run a concurrent transfer transactions
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}
	for i := 0; i < n; i++ {

		// check results
		err := <-errs
		require.NoError(t, err)
		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		// check entries

		// fromEntry
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		// toEntry
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		//check accounts
		//fromAccount := result.FromAccount
		//require.NotEmpty(t, fromAccount)
		//require.Equal(t, account1.ID, fromAccount.ID)
		//
		//toAccount := result.ToAccount
		//require.NotEmpty(t, toAccount)
		//require.Equal(t, account2.ID, toAccount.ID)
		//
		////check account's balance
		//diff1 := account1.Balance - fromAccount.Balance
		//diff2 := account2.Balance - toAccount.Balance
		//require.Equal(t, diff2, diff1)
		//require.True(t, diff1 > 0)
		//require.True(t, diff1%amount == 0)
		//
		////check the final updated balance
		//updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
		//require.NoError(t, err)
		//updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
		//require.NoError(t, err)
		//
		//require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
		//require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)

	}
}
