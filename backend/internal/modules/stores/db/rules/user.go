package rules

import (
	"context"
	"go.uber.org/zap"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/access"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/features/logger"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
	"waterfall-backend/internal/modules/stores/db/ent/user"
)

// UsersAllowMutateSelf разрешает обновлять собственного пользователя
func UsersAllowMutateSelf() privacy.UserMutationRuleFunc {
	return func(ctx context.Context, m *ent.UserMutation) error {
		if m.Op().Is(ent.OpUpdate | ent.OpUpdateOne | ent.OpDelete | ent.OpDeleteOne) {
			ss, ok := session.GetFromCtx(ctx)
			if !ok {
				return privacy.Skip
			}
			currentUserUuid := ss.UserUuid

			if currentUserUuid == "" {
				return privacy.Skip
			}

			setUuid, set := m.ID()
			if !set {
				logger.GetFromCtx(ctx).Error("Значение Uuid при обновлении не установлено", zap.String(logger.Operation, "UsersAllowMutateSelf"))
				return privacy.Skip
			}

			if currentUserUuid == setUuid {
				return privacy.Allow
			}
		}

		return privacy.Skip
	}
}

// UsersAllowIfProtectionNotNeeded разрешает доступ и ограничивает выборку, если защита не требуется
func UsersAllowIfProtectionNotNeeded() privacy.UserQueryRuleFunc {
	return func(ctx context.Context, q *ent.UserQuery) error {
		protectionNeeded := access.ProtectionNeeded(ctx)
		if !protectionNeeded {
			return privacy.Allow
		}

		return privacy.Skip
	}
}

// UsersFilterBySelfUuid реализует защиту: фильтрует по Uuid текущего пользователя
func UsersFilterBySelfUuid() privacy.UserQueryRuleFunc {
	return func(ctx context.Context, q *ent.UserQuery) error {
		ss, ok := session.GetFromCtx(ctx)
		if !ok {
			return err_const.ErrMissingSession
		}

		logger.GetFromCtx(ctx).Debug("Защита пользователей применена", zap.String(logger.Operation, "UsersFilterBySelfUuid"))

		if ss.UserUuid == "" {
			return privacy.Deny
		}
		q.Where(user.ID(ss.UserUuid))

		return privacy.Skip
	}
}
