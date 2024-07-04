package seeds

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/utils"
)

func (s *Engine) permissions(ctx context.Context, _ config.Config) error {
	toCreate := []*ent.PermissionCreate{
		s.db.Permission.Create().
			SetID(permissions.FilesCreate).
			SetDescription("Создание файлов"),
		s.db.Permission.Create().
			SetID(permissions.FilesRead).
			SetDescription("Получение файлов"),
		s.db.Permission.Create().
			SetID(permissions.ReleasesCreate).
			SetDescription("Создание релизов"),
		s.db.Permission.Create().
			SetID(permissions.ReleasesRead).
			SetDescription("Получение релизов"),
		s.db.Permission.Create().
			SetID(permissions.ReleasesUpdate).
			SetDescription("Редактирование релизов"),
		s.db.Permission.Create().
			SetID(permissions.ReleasesDelete).
			SetDescription("Удаление релизов"),
		s.db.Permission.Create().
			SetID(permissions.UsersCreate).
			SetDescription("Создание пользователей"),
		s.db.Permission.Create().
			SetID(permissions.UsersRead).
			SetDescription("Получение пользователей"),
		s.db.Permission.Create().
			SetID(permissions.UsersReadSelf).
			SetDescription("Получение своего пользователя"),
		s.db.Permission.Create().
			SetID(permissions.UsersUpdate).
			SetDescription("Редактирование пользователей"),
		s.db.Permission.Create().
			SetID(permissions.UsersUpdateSelf).
			SetDescription("Редактирование своего пользователя"),
		s.db.Permission.Create().
			SetID(permissions.UsersDelete).
			SetDescription("Удаление пользователей"),
		s.db.Permission.Create().
			SetID(permissions.UsersDeleteSelf).
			SetDescription("Удаление своего пользователя"),
		s.db.Permission.Create().
			SetID(permissions.ChaptersCreate).
			SetDescription("Создание глав"),
		s.db.Permission.Create().
			SetID(permissions.ChaptersRead).
			SetDescription("Получение глав"),
		s.db.Permission.Create().
			SetID(permissions.ChaptersUpdate).
			SetDescription("Редактирование глав"),
		s.db.Permission.Create().
			SetID(permissions.ChaptersDelete).
			SetDescription("Удаление глав"),
		s.db.Permission.Create().
			SetID(permissions.PartsCreate).
			SetDescription("Создание частей"),
		s.db.Permission.Create().
			SetID(permissions.PartsRead).
			SetDescription("Получение частей"),
		s.db.Permission.Create().
			SetID(permissions.PartsUpdate).
			SetDescription("Редактирование частей"),
		s.db.Permission.Create().
			SetID(permissions.PartsDelete).
			SetDescription("Удаление частей"),
		s.db.Permission.Create().
			SetID(permissions.CommentsCreate).
			SetDescription("Создание комментариев"),
		s.db.Permission.Create().
			SetID(permissions.CommentsRead).
			SetDescription("Получение комментариев"),
		s.db.Permission.Create().
			SetID(permissions.CommentsUpdate).
			SetDescription("Редактирование комментариев"),
		s.db.Permission.Create().
			SetID(permissions.CommentsDelete).
			SetDescription("Удаление комментариев"),
	}

	err := s.db.Permission.CreateBulk(
		toCreate...,
	).OnConflict(
		sql.ConflictColumns("id"),
		sql.ResolveWithNewValues(),
	).Exec(ctx)

	return utils.DefaultErrorWrapper(err)
}
