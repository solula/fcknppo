package release

import (
	"go.uber.org/fx"
	repo2 "waterfall-backend/internal/modules/domain/release/repo"
	service2 "waterfall-backend/internal/modules/domain/release/service"
	"waterfall-backend/internal/modules/domain/release/worker"
)

var (
	ServiceModule = fx.Module("release-service",
		service2.Module,
		repo2.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo2.ReleaseRepo) *repo2.ReleaseRepo { return r },
				fx.As(new(service2.IReleaseRepo)),
			),
		),
	)

	WorkerModule = fx.Module("release-worker",
		worker.Module,
	)
)
