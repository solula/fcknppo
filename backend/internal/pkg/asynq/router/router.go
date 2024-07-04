package router

import (
	"github.com/hibiken/asynq"
	"waterfall-backend/internal/pkg/option"
)

type RouteOption interface {
	option.Interface
}

type Router interface {
	Handle(pattern string, handler interface{}, opts ...RouteOption) Router
	Group(prefix string) Router
	Use(middleware ...asynq.MiddlewareFunc)
}
