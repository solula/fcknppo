package data_migrations

import (
	"context"
	"fmt"
	"waterfall-backend/internal/modules/features/config"
	ent2 "waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/utils"
)

func (s *Engine) Migrate(ctx context.Context, cfg config.Config) error {
	err := s.migrateWithFuncs(ctx, cfg,
		s.initial,
		s.addUsersReadSelfPermission,
		s.addCommentsPermissions,
		s.dropUsersPermissionsTable,
	)
	if err != nil {
		return fmt.Errorf("ошибка при миграции: %w", err)
	}

	return nil
}

func (s *Engine) migrateWithFuncs(ctx context.Context, cfg config.Config, funcs ...func(context.Context, config.Config) error) error {
	var migrationErr error
	err := utils.WithTx(ctx, s.db, func(tx *ent2.Tx) error {
		migration, err := tx.Migrations.Query().First(ctx)
		if err != nil && !ent2.IsNotFound(err) {
			return fmt.Errorf("ошибка при получении миграции: %w", err)
		}
		// Если последняя миграция не найдена, то prevMigrationNumber будет 0.
		// Запись в таблице миграций будет создана первой же миграцией.
		prevMigrationNumber := 0
		if migration != nil {
			prevMigrationNumber = migration.Migrated
		}

		migratedSeccessfully := prevMigrationNumber
		for i, f := range funcs {
			migrationNumber := i + 1
			if migrationNumber > prevMigrationNumber {
				if err := f(ctx, cfg); err != nil {
					migrationErr = err
					break
				}
				migratedSeccessfully++
			}
		}

		err = tx.Migrations.Update().SetMigrated(migratedSeccessfully).Exec(ctx)
		if err != nil {
			return fmt.Errorf("ошибка при обновлении миграции: %w", err)
		}

		return nil
	})
	if err != nil {
		return err
	}

	if migrationErr != nil {
		return migrationErr
	}

	return nil
}
