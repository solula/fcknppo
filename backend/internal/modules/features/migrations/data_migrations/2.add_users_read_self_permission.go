package data_migrations

import (
	"context"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/utils"
)

func (s *Engine) addUsersReadSelfPermission(ctx context.Context, _ config.Config) error {
	err := utils.WithTx(ctx, s.db, func(tx *ent.Tx) error {
		err := tx.Permission.Create().
			SetID(permissions.UsersReadSelf).
			SetDescription("Получение своего пользователя").
			Exec(ctx)
		if err != nil {
			return err
		}

		err = tx.Role.UpdateOneID(roles.Free).
			AddPermissionIDs(permissions.UsersReadSelf).
			Exec(ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return utils.DefaultErrorWrapper(err)
}
