package worker

import (
	"go.uber.org/fx"
)

var Module = fx.Module("worker",
	fx.Provide(NewFileStorageController),
	fx.Invoke(InvokeFileStorageController),
)
