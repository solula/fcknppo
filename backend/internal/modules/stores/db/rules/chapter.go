package rules

import (
	"context"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/access"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/chapter"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
)

// ChaptersHasRelease оставляет только те главы, для которых выпущен релиз
func ChaptersHasRelease() privacy.ChapterQueryRuleFunc {
	return func(ctx context.Context, q *ent.ChapterQuery) error {
		q.Where(chapter.HasRelease())

		return nil
	}
}

// ChaptersFilterByReleaseDateOrSetSelection фильтрует по дате релиза
// или добавляет выборку разрешенных полей (если в контексте есть 'ignoreReleaseDate')
func ChaptersFilterByReleaseDateOrSetSelection(selection ...string) privacy.ChapterQueryRuleFunc {
	return func(ctx context.Context, q *ent.ChapterQuery) error {
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

			q.Where(chapter.HasReleaseWith(filterByReleaseDate(ss.ReleaseDelay)))
		}
		return nil
	}
}
