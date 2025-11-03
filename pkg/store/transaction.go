package sql_store

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Store struct {
	db *sqlx.DB
}

// NewStore creates a new Store instance
func NewStore(db *sqlx.DB) *Store {
	return &Store{db: db}
}

// runInTransaction is a helper function that begins a transaction,
// executes the provided function `fn`, and handles commit/rollback.
func (s *Store) RunInTransaction(ctx context.Context, fn func(*sqlx.Tx) error) (err error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "tx begin failed")
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
			if err != nil {
				err = errors.Wrap(err, "tx commit failed")
			}
		}
	}()

	err = fn(tx)
	return err
}
