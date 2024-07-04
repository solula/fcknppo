package rules

import (
	"context"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/access"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/part"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
)

// PartsHasRelease оставляет только те части, для которых выпущен релиз
func PartsHasRelease() privacy.PartQueryRuleFunc {
	return func(ctx context.Context, q *ent.PartQuery) error {
		q.Where(part.HasRelease())

		return nil
	}
}

// PartsFilterByReleaseDateOrSetSelection фильтрует по дате релиза
// или добавляет выборку разрешенных полей (если в контексте есть 'ignoreReleaseDate')
func PartsFilterByReleaseDateOrSetSelection(selection ...string) privacy.PartQueryRuleFunc {
	return func(ctx context.Context, q *ent.PartQuery) error {
		switch access.IgnoreReleaseDate(ctx) {
		case true:
			// Если требуется проигнорировать дату релиза -> ограничиваем выборку доступными полями
			q.Select(selection...)
		case false:
			// Иначе навешиваем фильтр по дате релиза с учетом задержки релиза для пользователя
			ss, ok := session.GetFromCtx(ctx)
			if !ok {
				return err_const.ErrMissingSession
			}

			q.Where(part.HasReleaseWith(filterByReleaseDate(ss.ReleaseDelay)))
		}
		return nil
	}
}
