package repo

import (
	"context"
	"waterfall-backend/internal/pkg/transaction"
)

func (r *CommentRepo) Tx(ctx context.Context) (transaction.TxRepo, transaction.Tx, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, nil, wrap(err)
	}

	return &CommentRepo{
		client: tx.Client(),
	}, tx, nil
}
