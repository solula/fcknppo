package repo

import (
	"context"
	"waterfall-backend/internal/pkg/transaction"
)

func (r *PartRepo) Tx(ctx context.Context) (transaction.TxRepo, transaction.Tx, error) {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return nil, nil, wrap(err)
	}

	return &PartRepo{
		client: tx.Client(),
	}, tx, nil
}
