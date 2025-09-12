package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomSalarySplit(t *testing.T) FinanceSalarySplits {
	arg := AddSalarySplitParams{
		UserID:             1,
		TotalSalary:        "5000",
		Month:              time.Now(),
		Notes:              sql.NullString{String: "Test Notes", Valid: true},
		IsFullyTransferred: sql.NullBool{Bool: false, Valid: true},
	}

	salarySplit, err := testQueries.AddSalarySplit(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, salarySplit)

	require.Equal(t, arg.UserID, salarySplit.UserID)
	require.Equal(t, arg.TotalSalary, salarySplit.TotalSalary)
	require.Equal(t, arg.Notes, salarySplit.Notes)
	require.Equal(t, arg.IsFullyTransferred, salarySplit.IsFullyTransferred)

	return salarySplit
}

func TestAddSalarySplit(t *testing.T) {
	createRandomSalarySplit(t)
}

func TestGetSalarySplitById(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)

	result, err := testQueries.GetSalarySplitById(context.Background(), salarySplit.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, salarySplit.ID, result.ID)
	require.Equal(t, salarySplit.UserID, result.UserID)
	require.Equal(t, salarySplit.TotalSalary, result.TotalSalary)
	require.Equal(t, salarySplit.Notes, result.Notes)
	require.Equal(t, salarySplit.IsFullyTransferred, result.IsFullyTransferred)
}

func TestGetSalarySplitsByUserId(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)

	results, err := testQueries.GetSalarySplitsByUserId(context.Background(), salarySplit.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, results)

	found := false
	for _, result := range results {
		if result.ID == salarySplit.ID {
			found = true
			break
		}
	}
	require.True(t, found)
}

func TestUpdateSalarySplitTotalById(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)

	arg := UpDateSalarySplitTotalByIdParams{
		TotalSalary: "6000",
		ID:          salarySplit.ID,
	}

	err := testQueries.UpDateSalarySplitTotalById(context.Background(), arg)
	require.NoError(t, err)

	updatedSalarySplit, err := testQueries.GetSalarySplitById(context.Background(), salarySplit.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedSalarySplit)

	require.Equal(t, arg.TotalSalary, updatedSalarySplit.TotalSalary)
}

func TestMarkSalarySplitAsFullyTransferredById(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)

	err := testQueries.MarkSalarySplitAsFullyTransferredById(context.Background(), salarySplit.ID)
	require.NoError(t, err)

	updatedSalarySplit, err := testQueries.GetSalarySplitById(context.Background(), salarySplit.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedSalarySplit)

	require.True(t, updatedSalarySplit.IsFullyTransferred.Bool)
}

func TestDeleteSalarySplitById(t *testing.T) {
	salarySplit := createRandomSalarySplit(t)

	err := testQueries.DeleteSalarySplitById(context.Background(), salarySplit.ID)
	require.NoError(t, err)

	deletedSalarySplit, err := testQueries.GetSalarySplitById(context.Background(), salarySplit.ID)
	require.Error(t, err)
	require.Empty(t, deletedSalarySplit)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
