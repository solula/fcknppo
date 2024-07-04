package router

import (
	"context"
	"github.com/hibiken/asynq"
)

type AsynqRouter struct {
	router        *asynq.ServeMux
	patternPrefix string
}

func NewAsynqRouter() *AsynqRouter {
	return &AsynqRouter{
		router:        asynq.NewServeMux(),
		patternPrefix: "",
	}
}

func (r *AsynqRouter) Handle(pattern string, handler interface{}, opts ...RouteOption) Router {
	r.router.Handle(r.patternPrefix+pattern, r.wrap(handler, opts...))
	return r
}

func (r *AsynqRouter) Group(prefix string) Router {
	groupPattern := r.patternPrefix + prefix

	groupMux := asynq.NewServeMux()
	r.router.Handle(groupPattern, groupMux)

	return &AsynqRouter{
		router:        groupMux,
		patternPrefix: groupPattern,
	}
}

func (r *AsynqRouter) Use(middlewares ...asynq.MiddlewareFunc) {
	r.router.Use(middlewares...)
}

func (r *AsynqRouter) wrap(handler interface{}, opts ...RouteOption) asynq.HandlerFunc {
	var internalHandler asynq.HandlerFunc

	// TODO Сделать обработку опций.
	//  Пока опций нет и обработки нет

	switch h := handler.(type) {
	case asynq.HandlerFunc:
		internalHandler = h
	default:
		internalHandler = sugarHandler(handler)
	}

	// Обработать опции
	return internalHandler
}

// ProcessTask метод, чтобы AsynqRouter был asynq.Handler
func (r *AsynqRouter) ProcessTask(ctx context.Context, task *asynq.Task) error {
	return r.router.ProcessTask(ctx, task)
}
