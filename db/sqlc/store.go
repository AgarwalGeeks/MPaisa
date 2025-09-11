package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, transactionError := store.db.BeginTx(ctx, nil)
	if transactionError != nil {
		return transactionError
	}

	q := New(tx)
	transactionError = fn(q)
	if transactionError != nil {
		if rollbackError := tx.Rollback(); rollbackError != nil {
			return fmt.Errorf("transaction error is: %v, rollback error is: %v", transactionError, rollbackError)
		}
		return transactionError
	}

	return tx.Commit()
}

// adds salary split and its items in a single transaction
func (store *Store) AddSalarySplitWithSplitItemsTx(ctx context.Context, salarySplit AddSalarySplitParams, splitItems []AddSalarySplitItemParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		var salarySplitResult FinanceSalarySplits

		salarySplitResult, err = q.AddSalarySplit(ctx, salarySplit)
		if err != nil {
			return err
		}

		for _, item := range splitItems {
			item.SplitID = salarySplitResult.ID
			_, err = q.AddSalarySplitItem(ctx, item)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
