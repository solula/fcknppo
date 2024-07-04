package utils

import (
	"context"
	"fmt"
	ent2 "waterfall-backend/internal/modules/stores/db/ent"
)

func WithTx(ctx context.Context, client *ent2.Client, fn func(tx *ent2.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
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
