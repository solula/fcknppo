package data_migrations

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/utils"
)

func (s *Engine) addCommentsPermissions(ctx context.Context, _ config.Config) error {
	err := utils.WithTx(ctx, s.db, func(tx *ent.Tx) error {
		// Новые права доступа
		toCreate := []*ent.PermissionCreate{
			tx.Permission.Create().
				SetID(permissions.CommentsCreate).
				SetDescription("Создание комментариев"),
			tx.Permission.Create().
				SetID(permissions.CommentsRead).
				SetDescription("Получение комментариев"),
			tx.Permission.Create().
				SetID(permissions.CommentsUpdate).
				SetDescription("Редактирование комментариев"),
			tx.Permission.Create().
				SetID(permissions.CommentsDelete).
				SetDescription("Удаление комментариев"),
		}

		// Пакетно создаем все новые права доступа
		err := s.db.Permission.CreateBulk(
			toCreate...,
		).OnConflict(
			sql.ConflictColumns("id"),
			sql.ResolveWithNewValues(),
		).Exec(ctx)
		if err != nil {
			return err
		}

		// Обновляем роли, добавляя новые права доступа
		err = tx.Role.UpdateOneID(roles.Guest).
			AddPermissionIDs(permissions.CommentsRead).
			Exec(ctx)
		if err != nil {
			return err
		}

		err = tx.Role.UpdateOneID(roles.Free).
			AddPermissionIDs(
				permissions.CommentsRead,
				permissions.CommentsCreate,
				permissions.CommentsUpdate,
				permissions.CommentsDelete,
			).
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
