package db

import (
	"context"
	"fmt"

	"simple_go/database/model"

	_ "github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	model.Querier
}

type SQLStore struct {
	connPool *pgxpool.Pool
	model.Querier
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Querier:  model.New(connPool),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*model.Queries) error) error {
	tx, err := store.connPool.Begin(ctx)
	if err != nil {
		return err
	}
	q := model.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v rbErr: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}
