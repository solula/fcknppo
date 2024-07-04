package seeds

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"waterfall-backend/internal/constants/users"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/utils"
	"waterfall-backend/internal/utils/password"
)

func (s *Engine) users(ctx context.Context, cfg config.Config) error {
	if cfg.AdminEmail == "" {
		return fmt.Errorf("в конфигурации не задан email администратора")
	}
	if cfg.AdminPassword == "" {
		return fmt.Errorf("в конфигурации не задан пароль администратора")
	}
	if cfg.SystemEmailAddress == "" {
		return fmt.Errorf("в конфигурации не задан системный email")
	}

	passwordHash, err := password.GenerateHash(cfg.AdminPassword)
	if err != nil {
		return err
	}

	toCreate := []*ent.UserCreate{
		s.db.User.Create().
			SetID(users.AdminUuid).
			SetEmail(cfg.AdminEmail).
			SetPasswordHash(passwordHash).
			SetFullname(users.AdminFullname).
			SetUsername(users.AdminUsername).
			//SetSerialNumber(1).
			SetEmailVerified(true).
			AddRoleIDs(roles.Free, roles.Premium, roles.Admin), // Тут перечисляем все пользовательские роли
		s.db.User.Create().
			SetID(users.SystemUuid).
			SetEmail(cfg.SystemEmailAddress).
			SetFullname(users.SystemFullname).
			SetUsername(users.SystemUsername).
			SetEmailVerified(true).
			AddRoleIDs(roles.Free, roles.Premium, roles.Admin), // Тут перечисляем все пользовательские роли
	}

	err = s.db.User.CreateBulk(
		toCreate...,
	).OnConflict(
		sql.ConflictColumns("uuid"),
		sql.ResolveWithNewValues(),
	).Exec(ctx)

	return utils.DefaultErrorWrapper(err)
}
