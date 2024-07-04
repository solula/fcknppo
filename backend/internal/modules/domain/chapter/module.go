package chapter

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/domain/chapter/repo"
	"waterfall-backend/internal/modules/domain/chapter/service"
)

var (
	ServiceModule = fx.Module("chapter",
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.ChapterRepo) *repo.ChapterRepo { return r },
				fx.As(new(service.IChapterRepo)),
			),
			fx.Annotate(
				func(r *repo.ChapterTextRepo) *repo.ChapterTextRepo { return r },
				fx.As(new(service.IChapterTextRepo)),
			),
		),
	)
)
