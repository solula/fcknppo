package authorization

import (
	"github.com/labstack/echo/v4"
	"waterfall-backend/internal/pkg/http/middlewares/jwt_extractor"
)

// WithJWTExtractor достает токен из заголовка и кладет его в контекст
func WithJWTExtractor() echo.MiddlewareFunc {
	return jwt_extractor.New(jwt_extractor.Config{
		ContextKey:  AccessTokenKey,
		TokenLookup: "header:Authorization:Bearer ",
	})
}
