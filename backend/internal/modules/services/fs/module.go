package fs

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/services/fs/http"
	"waterfall-backend/internal/modules/services/fs/i_fs"
	"waterfall-backend/internal/modules/services/fs/repo"
	"waterfall-backend/internal/modules/services/fs/s3"
	"waterfall-backend/internal/modules/services/fs/service"
	"waterfall-backend/internal/modules/services/fs/task_manager"
	"waterfall-backend/internal/modules/services/fs/worker"
)

var (
	ServiceModule = fx.Module("fs-service",
		service.Module,
		repo.Module,
		s3.Module,
		task_manager.Module,

		fx.Provide(
			fx.Annotate(
				func(r *repo.FileRepo) *repo.FileRepo { return r },
				fx.As(new(service.IFileRepo)),
			),
			fx.Annotate(
				func(s3 *s3.S3) *s3.S3 { return s3 },
				fx.As(new(service.IS3)),
			),
			fx.Annotate(
				func(r *task_manager.FileStorageTaskManager) *task_manager.FileStorageTaskManager { return r },
				fx.As(new(service.IFileStorageTaskManager)),
			),
			fx.Annotate(
				func(r *service.FileStorageService) *service.FileStorageService { return r },
				fx.As(new(i_fs.IBucketCreator)),
			),
		),
	)

	HTTPModule = fx.Module("fs-http",
		http.Module,
	)

	WorkerModule = fx.Module("fs-worker",
		worker.Module,
	)
)
