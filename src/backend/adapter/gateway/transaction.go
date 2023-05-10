package gateway

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sky0621/familiagildo/app"
	"github.com/sky0621/familiagildo/domain/repository"
	"github.com/sky0621/familiagildo/driver/db"
)

func NewTransactionRepository(cli *db.Client) repository.TransactionRepository {
	return &transactionRepository{db: cli.DB}
}

type transactionRepository struct {
	db   *sql.DB
	opts *sql.TxOptions
}

func (r *transactionRepository) ExecInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := r.db.BeginTx(ctx, r.opts)
	if err != nil {
		return app.WithStackError(err)
	}

	c := context.WithValue(ctx, app.TxCtxKey, tx)

	done := false

	defer func() {
		if tx != nil && !done {
			if err := tx.Rollback(); err != nil {
				fmt.Println(err)
			}
		}
	}()

	if err := fn(c); err != nil {
		return app.WithStackError(err)
	}

	if err := tx.Commit(); err != nil {
		return app.WithStackError(err)
	}
	done = true

	return nil
}
