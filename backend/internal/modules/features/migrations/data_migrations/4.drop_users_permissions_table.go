package data_migrations

import (
	"context"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/stores/db/utils"
)

func (s *Engine) dropUsersPermissionsTable(ctx context.Context, _ config.Config) error {
	_, err := s.db.ExecContext(ctx, `DROP TABLE IF EXISTS auth_users_permissions;`)
	if err != nil {
		return err
	}

	return utils.DefaultErrorWrapper(err)
}
