package api

import (
	"waterfall-backend/internal/pkg/http/middlewares/authorization"
	"waterfall-backend/internal/pkg/routers"
)

type Router routers.Router

func InitRouter(router routers.Router, auther authorization.Auther) Router {
	/* Добавляем проверку аутентификации для группы /api */

	router.Use(
		// Достаем токен из заголовка запроса
		authorization.WithJWTExtractor(),
		// На основе токена из контекста генерируем сессию
		authorization.WithAuth(auther),
	)

	return router
}
