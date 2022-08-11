package sqlc

import (
	"context"
	"database/sql"
	"testing"

	"github.com/abigiya/internal/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Firstname:    util.RandomName(),
		Lastname:     util.RandomName(),
		Emailaddress: util.RandomEmail(),
		Username:     util.RandomUserName(),
		Password:     int64(util.RandomPassword()),
	}

	acc, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.Firstname, arg.Firstname)
	require.Equal(t, acc.Lastname, arg.Lastname)
	require.Equal(t, acc.Emailaddress, arg.Emailaddress)
	require.Equal(t, acc.Username, arg.Username)
	require.Equal(t, acc.Password, arg.Password)

	require.NotZero(t, acc.ID)
	require.NotZero(t, acc.CreatedAt)

	return acc
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	t.Run("Fill username and password", func(t *testing.T) {
		followAccountOnly, err := testQueries.GetAccount(context.Background(), GetAccountParams{
			Username: account1.Username,
			Password: int64(account1.Password),
		})
		require.NoError(t, err)
		require.NotEmpty(t, followAccountOnly)
		require.Equal(t, account1.ID, followAccountOnly.ID)
		require.Equal(t, account1.Username, followAccountOnly.Username)
		require.Equal(t, account1.Password, followAccountOnly.Password)
		require.Equal(t, account1.CreatedAt, followAccountOnly.CreatedAt)
	})

	t.Run("user name not exit", func(t *testing.T) {
		followUsernameFormat, err := testQueries.GetAccount(context.Background(), GetAccountParams{
			Username: util.RandomUserName(),
			Password: int64(util.RandomPassword()),
		})
		require.EqualError(t, err, sql.ErrNoRows.Error())
		require.Empty(t, followUsernameFormat)
	})
}
