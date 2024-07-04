package rules

import (
	"context"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/chapter"
	"waterfall-backend/internal/modules/stores/db/ent/chaptertext"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
)

// ChapterTextsHasRelease оставляет только те главы, для которых выпущен релиз
func ChapterTextsHasRelease() privacy.ChapterTextQueryRuleFunc {
	return func(ctx context.Context, q *ent.ChapterTextQuery) error {
		q.Where(chaptertext.HasChapterWith(chapter.HasRelease()))

		return nil
	}
}

// ChapterTextsFilterByReleaseDate фильтрует по дате релиза
func ChapterTextsFilterByReleaseDate() privacy.ChapterTextQueryRuleFunc {
	return func(ctx context.Context, q *ent.ChapterTextQuery) error {
		// Навешиваем фильтр по дате релиза с учетом задержки релиза для пользователя
		ss, ok := session.GetFromCtx(ctx)
		if !ok {
			return err_const.ErrMissingSession
		}

		q.Where(chaptertext.HasChapterWith(chapter.HasReleaseWith(filterByReleaseDate(ss.ReleaseDelay))))

		return nil
	}
}
