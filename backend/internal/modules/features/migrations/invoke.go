package migrations

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"waterfall-backend/internal/models/access"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/features/logger"
	"waterfall-backend/internal/modules/features/migrations/data_migrations"
	"waterfall-backend/internal/modules/features/migrations/seeds"
	"waterfall-backend/internal/modules/stores/db/ent"
)

func InvokeMigrations(
	client *ent.Client,
	minio *minio.Client,
	cfg config.Config,
	lg *zap.Logger,
	lifecycle fx.Lifecycle,
) error {
	ctx := context.Background()
	ctx = access.SetSystem(ctx) // Все будем делать от имени системы
	ctx = logger.SetToCtx(ctx, lg)

	// Делаем миграции только при необходимости
	if cfg.AutoMigrate {
		if err := client.Schema.Create(ctx); err != nil {
			return fmt.Errorf("ошибка при миграции схемы: %w", err)
		}

		if err := seeds.NewEngine(client, minio).Migrate(ctx, cfg); err != nil {
			return fmt.Errorf("ошибка при посевной миграции: %w", err)
		}

		if err := data_migrations.NewEngine(client, minio).Migrate(ctx, cfg); err != nil {
			return fmt.Errorf("ошибка при миграции данных: %w", err)
		}
	}

	lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return client.Close()
		},
	})

	return nil
}
