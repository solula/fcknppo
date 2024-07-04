package transaction

import (
	"context"
	"fmt"
)

// TxRepo репозиторий, который может порождать транзакции
type TxRepo interface {
	Tx(ctx context.Context) (TxRepo, Tx, error)
}

// Tx транзакция
type Tx interface {
	Commit() error
	Rollback() error
}

func WithTx[Repo TxRepo](ctx context.Context, repo TxRepo, fn func(tx Repo) error) error {
	txRepo, tx, err := repo.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()

	typedRepo, ok := txRepo.(Repo)
	if !ok {
		return fmt.Errorf("указанный в WithTx тип %T не является Repo", txRepo)
	}
	if err := fn(typedRepo); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%w: ошибка при откате транзакции: %v", err, rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("ошибка при завершении транзакции: %w", err)
	}
	return nil
}
