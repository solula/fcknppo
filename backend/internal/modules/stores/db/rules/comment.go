package rules

import (
	"context"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/comment"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
)

// CommentsFilterByAuthorUuid фильтрует комментарии по Uuid автора
func CommentsFilterByAuthorUuid() privacy.CommentMutationRuleFunc {
	return func(ctx context.Context, q *ent.CommentMutation) error {
		ss, ok := session.GetFromCtx(ctx)
		if !ok {
			return err_const.ErrMissingSession
		}

		if ss.UserUuid == "" {
			return privacy.Deny
		}
		q.Where(comment.AuthorUUID(ss.UserUuid))

		return privacy.Skip
	}
}
