package repository

import "context"

type TransactionRepository interface {
	ExecInTransaction(ctx context.Context, fn func(ctx context.Context) error) error
}
