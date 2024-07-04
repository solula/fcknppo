package rules

import (
	"context"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/file"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
)

// FilesFilterByCreatorUuid фильтрует по Uuid создателя
func FilesFilterByCreatorUuid() privacy.FileMutationRuleFunc {
	return func(ctx context.Context, q *ent.FileMutation) error {
		ss, ok := session.GetFromCtx(ctx)
		if !ok {
			return err_const.ErrMissingSession
		}

		if ss.UserUuid == "" {
			return privacy.Deny
		}
		q.Where(file.CreatorUUID(ss.UserUuid))

		return privacy.Skip
	}
}
