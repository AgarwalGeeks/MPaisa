package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCreditCard(t *testing.T) {
	creditCardDetail := AddCreditCardParams{
		BankName:   "Test Bank",
		CardName:   "Test Card",
		CardNumber: "1234567890123456",
		Cvv:        "123",
		Pin:        "123",
		Usage:      sql.NullString{String: "Test Usage", Valid: true},
		UserID:     1,
	}

	FinanceCreditCard, err := testQueries.AddCreditCard(context.Background(), creditCardDetail)
	require.NoError(t, err)
	require.NotEmpty(t, FinanceCreditCard)

	require.Equal(t, creditCardDetail.BankName, FinanceCreditCard.BankName)
	require.Equal(t, creditCardDetail.CardName, FinanceCreditCard.CardName)
	require.Equal(t, creditCardDetail.CardNumber, FinanceCreditCard.CardNumber)
	require.Equal(t, creditCardDetail.Cvv, FinanceCreditCard.Cvv)
	require.Equal(t, creditCardDetail.Pin, FinanceCreditCard.Pin)
	require.Equal(t, creditCardDetail.Usage, FinanceCreditCard.Usage)
	require.Equal(t, creditCardDetail.UserID, FinanceCreditCard.UserID)
}
