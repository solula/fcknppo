package middlewares

import (
	"github.com/labstack/echo/v4"
	"waterfall-backend/internal/models/access"
)

// WithSystem добавляет в контекст сессию системного пользователя и соответствующий ключ
func WithSystem() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()
			ctx = access.SetSystem(ctx)

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
