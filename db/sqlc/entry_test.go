package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/utils"
)

func CreateRandomEntry(t *testing.T) (Entry, error) {
	arg := CreateEntryParams{
		AccountID: int64(43),
		Amount:    utils.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	return entry, err
}

func TestCreateEntry(t *testing.T) {
	entry, err := CreateRandomEntry(t)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
}

func TestGetEntry(t *testing.T) {
	entry, err := CreateRandomEntry(t)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	entry2, err2 := testQueries.GetEntry(context.Background(), entry.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, entry2)
}

func TestGetEntryList(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomEntry(t)
	}

	arg := ListEntriesParams{
		Offset:    5,
		Limit:     5,
		AccountID: int64(43),
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}
