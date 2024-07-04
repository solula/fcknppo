package data_migrations

import (
	"context"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/stores/db/utils"
)

func (s *Engine) initial(ctx context.Context, _ config.Config) error {
	err := s.db.Migrations.Create().SetMigrated(1).Exec(ctx)

	return utils.DefaultErrorWrapper(err)
}