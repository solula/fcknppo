package worker

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/pkg/asynq/router"
)

type Router router.Router

type Routers struct {
	fx.Out

	Router Router
}

// NewWorker создает все необходимые роутеры очередей
func NewWorker(asynqRouter router.Router) Routers {
	return Routers{
		Router: asynqRouter,
	}
}
