package task_manager

import (
	"go.uber.org/fx"
)

var Module = fx.Module("task-manager",
	fx.Provide(NewTaskManager),
	fx.Invoke(InvokeTaskManager),
)
