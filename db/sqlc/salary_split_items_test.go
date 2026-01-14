package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomSalarySplitItem(t *testing.T, SplitID int32) FinanceSalarySplitItems {
	arg := AddSalarySplitItemParams{
		SplitID:       SplitID,
		CategoryName:  "Test CategoryName",
		Amount:        "1000",
		MoveTo:        sql.NullString{String: "Test MoveTo", Valid: true},
		IsTransferred: sql.NullBool{Bool: false, Valid: true},
	}

	salarySplitItem, err := testQueries.AddSalarySplitItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, salarySplitItem)

	require.Equal(t, arg.SplitID, salarySplitItem.SplitID)
	require.Equal(t, arg.CategoryName, salarySplitItem.CategoryName)
	require.Equal(t, arg.Amount, salarySplitItem.Amount)
	require.Equal(t, arg.MoveTo, salarySplitItem.MoveTo)
	require.Equal(t, arg.IsTransferred, salarySplitItem.IsTransferred)

	return salarySplitItem
}

func TestAddSalarySplitItem(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)
	createRandomSalarySplitItem(t, salarySplit.ID)
}

func TestGetSalarySplitItemsBySplitId(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)
	salarySplitItem := createRandomSalarySplitItem(t, salarySplit.ID)

	results, err := testQueries.GetSalarySplitItemsBySplitId(context.Background(), salarySplit.ID)
	require.NoError(t, err)
	require.NotEmpty(t, results)

	found := false
	for _, result := range results {
		if result.ID == salarySplitItem.ID {
			found = true
			break
		}
	}
	require.True(t, found)
}

func TestGetAllSalarySplitItemsByUserId(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)
	salarySplitItem := createRandomSalarySplitItem(t, salarySplit.ID)

	// results, err := testQueries.GetAllSalarySplitItemsByUserId(context.Background(), 1)
	// require.NoError(t, err)
	// require.NotEmpty(t, results)

	// found := false
	// for _, result := range results {
	// 	if result.ID == salarySplitItem.ID {
	// 		found = true
	// 		break
	// 	}
	// }
	// require.True(t, found)
}

func TestUpdateSalarySplitItemAmountById(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)
	salarySplitItem := createRandomSalarySplitItem(t, salarySplit.ID)

	arg := UpdateSalarySplitItemAmountByIdParams{
		Amount: "2000",
		ID:     salarySplitItem.ID,
	}

	err := testQueries.UpdateSalarySplitItemAmountById(context.Background(), arg)
	require.NoError(t, err)
}

func TestMarkSalarySplitItemAsTransferredById(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)
	salarySplitItem := createRandomSalarySplitItem(t, salarySplit.ID)

	err := testQueries.MarkSalarySplitItemAsTransferredById(context.Background(), salarySplitItem.ID)
	require.NoError(t, err)
}

func TestDeleteSalarySplitItemsBySplitId(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)
	createRandomSalarySplitItem(t, salarySplit.ID)

	err := testQueries.DeleteSalarySplitItemsBySplitId(context.Background(), salarySplit.ID)
	require.NoError(t, err)
}
