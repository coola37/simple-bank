package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/utils"
)

func CreateRandomTransfer(t *testing.T) Transfer {
	arg := CreateTransferParams{
		ToAccountID:   int64(44),
		FromAccountID: int64(43),
		Amount:        utils.RandomMoney(),
	}
	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	//require.NotEmpty(transfer)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)
	require.NotEmpty(t, transfer)
}

func TestGetTransfer(t *testing.T) {
	transfer := CreateRandomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)
}

func TestGetTransferList(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTransfer(t)
	}
	arg := ListTransfersParams{
		ToAccountID:   int64(44),
		FromAccountID: int64(43),
		Limit:         10,
		Offset:        10,
	}

	transferList, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	for _, transfer := range transferList {
		require.NotEmpty(t, transfer)
	}
}
