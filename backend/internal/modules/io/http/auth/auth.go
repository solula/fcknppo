package auth

import (
	"waterfall-backend/internal/pkg/http/middlewares"
	"waterfall-backend/internal/pkg/routers"
)

type Router routers.Router

func InitRouter(router routers.Router) Router {
	router.Use(
		// Аутентификация и регистрация происходят под системным пользователем
		middlewares.WithSystem(),
	)

	return router
}
