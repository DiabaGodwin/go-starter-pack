package repository

import (
	"context"
	"real-estate-app/internal/db"
)

func (s *Store) InTx(ctx context.Context, fn func(q *db.Queries) error) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := s.queries.WithTx(tx)

	if err := fn(qtx); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
