package rules

import (
	"context"
	"slices"
	"waterfall-backend/internal/models/access"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
)

// AllowIfSystem принудительно разрешить доступ, если действие от имени системы
func AllowIfSystem() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		if access.IsSystem(ctx) {
			return privacy.Allow
		}

		return privacy.Skip
	})
}

// AllowIfAdmin принудительно разрешить доступ, если пользователь админ
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		ss, ok := session.GetFromCtx(ctx)
		if !ok {
			return privacy.Skip
		}

		if slices.Contains(ss.Roles, roles.Admin) {
			return privacy.Allow
		}

		return privacy.Skip
	})
}
