package part

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/domain/part/repo"
	"waterfall-backend/internal/modules/domain/part/service"
)

var (
	ServiceModule = fx.Module("part",
		service.Module,
		repo.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.PartRepo) *repo.PartRepo { return r },
				fx.As(new(service.IPartRepo)),
			),
		),
	)
)
