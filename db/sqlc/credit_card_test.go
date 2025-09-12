package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreditCardMethods(t *testing.T) {
	creditCardDetail := AddCreditCardParams{
		BankName:   "Test Bank",
		CardName:   "Test Card",
		CardNumber: "1234567890123456",
		Cvv:        "123",
		Pin:        "123",
		Usage:      sql.NullString{String: "Test Usage", Valid: true},
		UserID:     1,
	}

	newCreditCard, err := testQueries.AddCreditCard(context.Background(), creditCardDetail)
	require.NoError(t, err)
	require.NotEmpty(t, newCreditCard)

	require.Equal(t, creditCardDetail.BankName, newCreditCard.BankName)
	require.Equal(t, creditCardDetail.CardName, newCreditCard.CardName)
	require.Equal(t, creditCardDetail.CardNumber, newCreditCard.CardNumber)
	require.Equal(t, creditCardDetail.Cvv, newCreditCard.Cvv)
	require.Equal(t, creditCardDetail.Pin, newCreditCard.Pin)
	require.Equal(t, creditCardDetail.Usage, newCreditCard.Usage)
	require.Equal(t, creditCardDetail.UserID, newCreditCard.UserID)

	creditCards, err := testQueries.GetAllCreditCards(context.Background(), creditCardDetail.UserID)
	creditCardsByUserId, err := testQueries.GetCreditCardsByUserId(context.Background(), creditCardDetail.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, creditCards)
	require.GreaterOrEqual(t, len(creditCards), 1)

	require.NoError(t, err)
	require.NotEmpty(t, creditCardsByUserId)
	require.GreaterOrEqual(t, len(creditCardsByUserId), 1)

	creditCardNumber := GetCreditCardByCardNumberParams{
		CardNumber: creditCardDetail.CardNumber,
		UserID:     1,
	}

	creditCardByNumber, err := testQueries.GetCreditCardByCardNumber(context.Background(), creditCardNumber)

	require.NoError(t, err)
	require.NotEmpty(t, creditCardByNumber)
	require.Equal(t, creditCardDetail.CardNumber, creditCardByNumber.CardNumber)

	getCreditCardByUsageParams := GetCreditCardByUsageParams{
		Usage:  sql.NullString{String: "Test Usage", Valid: true},
		UserID: 1,
	}

	creditCardByUsage, err := testQueries.GetCreditCardByUsage(context.Background(), getCreditCardByUsageParams)

	require.NoError(t, err)
	require.NotEmpty(t, creditCardByUsage)
	require.Equal(t, creditCardDetail.Usage, creditCardByUsage.Usage)

	updateCreditCardUsageByCardNumberParams := UpdateCreditCardUsageByCardNumberParams{
		Usage:      sql.NullString{String: "Updated Usage", Valid: true},
		CardNumber: creditCardDetail.CardNumber,
		UserID:     1,
	}

	updateErr := testQueries.UpdateCreditCardUsageByCardNumber(context.Background(), updateCreditCardUsageByCardNumberParams)

	require.NoError(t, updateErr)

	creditCardAfterUsageUpdate, err := testQueries.GetCreditCardByCardNumber(context.Background(), creditCardNumber)

	require.NoError(t, err)
	require.NotEmpty(t, creditCardAfterUsageUpdate)
	require.Equal(t, updateCreditCardUsageByCardNumberParams.Usage, creditCardAfterUsageUpdate.Usage)

	deleteCreditCardByCardNumberParams := DeleteCreditCardByCardNumberParams{
		CardNumber: creditCardDetail.CardNumber,
		UserID:     1,
	}

	err = testQueries.DeleteCreditCardByCardNumber(context.Background(), deleteCreditCardByCardNumberParams)
	require.NoError(t, err)

	deletedCreditCard, err := testQueries.GetCreditCardByCardNumber(context.Background(), creditCardNumber)
	require.Error(t, err)
	require.Empty(t, deletedCreditCard)
	require.EqualError(t, err, sql.ErrNoRows.Error())

}
