package db

import (
	"context"
	"database/sql"
)

// Store provides all functions to execute db queries and transactins
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transactins
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
    tx, err := store.db.BeginTx(ctx, nil)
    if err != nil {
        return eri
    }

    q := New(tx)
    err = fn(q)
    if err != nil {
        rbErr := tx.Rollback(); rbErr != nil {
            return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
        }
        return err
    }

    return tx.Commit()
}

// TransferTxParams contains the input parameters of the transfer transactins
type TransferTxParams struct {

}

// TransferTx performs a money transfer from one account to the other.
//T It creates a transfer record, and account entries, and update accounts balance within a single database transactins
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {

}
