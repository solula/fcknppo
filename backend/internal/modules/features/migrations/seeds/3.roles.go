package seeds

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"time"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/utils"
)

// Соотношение ролей и из прав доступа
var rolesPermissions = map[roles.Type][]permissions.Type{
	// Роль гостя сама по себе (не объединяется с другими ролями)
	roles.Guest: {
		permissions.FilesRead,
		permissions.ReleasesRead,
		permissions.UsersCreate,
		permissions.UsersRead,
		permissions.ChaptersRead,
		permissions.PartsRead,
		permissions.CommentsRead,
	},
	roles.Free: {
		permissions.FilesRead,
		permissions.FilesCreate,
		permissions.ReleasesRead,
		permissions.UsersCreate,
		permissions.UsersRead,
		permissions.UsersReadSelf,
		permissions.UsersUpdateSelf,
		permissions.UsersDeleteSelf,
		permissions.ChaptersRead,
		permissions.PartsRead,
		permissions.CommentsRead,
		permissions.CommentsCreate,
		permissions.CommentsUpdate,
		permissions.CommentsDelete,
	},
	roles.Premium: {},
	roles.Admin: {
		permissions.ReleasesCreate,
		permissions.ReleasesUpdate,
		permissions.ReleasesDelete,
		permissions.UsersUpdate,
		permissions.UsersDelete,
		permissions.ChaptersCreate,
		permissions.ChaptersUpdate,
		permissions.ChaptersDelete,
		permissions.PartsCreate,
		permissions.PartsUpdate,
		permissions.PartsDelete,
	},
}

func (s *Engine) roles(ctx context.Context, _ config.Config) error {
	zeroReleaseDelay := 0 * time.Hour.Seconds()
	freeReleaseDelay := 28 * 24 * time.Hour.Seconds()

	toCreate := []*ent.RoleCreate{
		s.db.Role.Create().
			SetID(roles.Admin).
			SetDescription("Администратор").
			SetReleaseDelay(zeroReleaseDelay).
			AddPermissionIDs(rolesPermissions[roles.Admin]...),
		s.db.Role.Create().
			SetID(roles.Premium).
			SetDescription("Платный пользователь").
			SetReleaseDelay(zeroReleaseDelay).
			AddPermissionIDs(rolesPermissions[roles.Premium]...),
		s.db.Role.Create().
			SetID(roles.Free).
			SetDescription("Бесплатный пользователь").
			SetReleaseDelay(freeReleaseDelay).
			AddPermissionIDs(rolesPermissions[roles.Free]...),
		s.db.Role.Create().
			SetID(roles.Guest).
			SetDescription("Гость (неавторизованный пользователь)").
			SetReleaseDelay(freeReleaseDelay).
			AddPermissionIDs(rolesPermissions[roles.Guest]...),
	}

	err := s.db.Role.CreateBulk(
		toCreate...,
	).OnConflict(
		sql.ConflictColumns("id"),
		sql.ResolveWithNewValues(),
	).Exec(ctx)

	return utils.DefaultErrorWrapper(err)
}
