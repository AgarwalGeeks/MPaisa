package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddSalarySplitWithSplitItemsTx(t *testing.T) {
	salarySplit := AddSalarySplitParams{
		UserID:             "1", // Changed to string
		TotalSalary:        "10000",
		Month:              time.Now(),
		Notes:              sql.NullString{String: "Test Salary Split", Valid: true},
		IsFullyTransferred: sql.NullBool{Bool: false, Valid: true},
	}

	splitItems := []AddSalarySplitItemParams{
		{
			CategoryName:  "Rent",
			Amount:        "2000",
			MoveTo:        sql.NullString{String: "Checking Account", Valid: true},
			IsTransferred: sql.NullBool{Bool: false, Valid: true},
		},
		{
			CategoryName:  "Food",
			Amount:        "1000",
			MoveTo:        sql.NullString{String: "Credit Card", Valid: true},
			IsTransferred: sql.NullBool{Bool: false, Valid: true},
		},
	}

	err := testStore.AddSalarySplitWithSplitItemsTx(context.Background(), salarySplit, splitItems)
	require.NoError(t, err)

	// Verify that the salary split was created
	// salarySplits, err := testStore.GetSalarySplitsByUserId(context.Background(), salarySplit.UserID)
	// require.NoError(t, err)
	// require.NotEmpty(t, salarySplits)

	// var createdSalarySplit FinanceSalarySplits
	// for _, s := range salarySplits {
	// 	if s.TotalSalary == salarySplit.TotalSalary {
	// 		createdSalarySplit = s
	// 		break
	// 	}
	// }
	// require.NotEmpty(t, createdSalarySplit)

	// // Verify that the split items were created
	// splitItemsResult, err := testStore.GetSalarySplitItemsBySplitId(context.Background(), createdSalarySplit.ID)
	// require.NoError(t, err)
	// require.Len(t, splitItemsResult, len(splitItems))

	// for _, item := range splitItemsResult {
	// 	require.Equal(t, createdSalarySplit.ID, item.SplitID)
	// }
}
