package scheduler

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/services/scheduler/service"
	"waterfall-backend/internal/modules/services/scheduler/task_manager"
)

var (
	ServiceModule = fx.Module("scheduler-service",
		service.Module,
		task_manager.Module,

		fx.Provide(
			fx.Annotate(
				func(r *task_manager.SchedulerTaskManager) *task_manager.SchedulerTaskManager { return r },
				fx.As(new(service.ISchedulerTaskManager)),
			),
		),
	)
)
