package main

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules"
)

func main() {
	fx.New(
		modules.WorkerModule,
	).Run()
}
