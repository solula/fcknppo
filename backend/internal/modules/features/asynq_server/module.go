package asynq_server

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/pkg/asynq/router"
)

var Module = fx.Module("asynq-server",
	fx.Provide(
		NewAsynqServer,
		NewAsynqRouter,
		// Аннотируем роутер как интерфейс Router
		fx.Annotate(
			func(r *router.AsynqRouter) *router.AsynqRouter { return r },
			fx.As(new(router.Router)),
		),
	),
	fx.Invoke(InvokeAsynqServer),
)
