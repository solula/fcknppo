package graphql

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/modules/features/logger"
	directives2 "waterfall-backend/internal/modules/io/graphql/directives"
	"waterfall-backend/internal/modules/io/graphql/generated"
	"waterfall-backend/internal/modules/io/graphql/resolvers"
	"waterfall-backend/internal/pkg/http/middlewares/authorization"

	_ "waterfall-backend/internal/pkg/http/error_handler/http_errors"
)

//go:generate go run -mod=mod github.com/99designs/gqlgen@v0.17.36

// RegisterGraphQL регистрация роута /graphql для обработки GraphQL запросов
// @tags GraphQL
// @security ApiKeyAuth
// @in header
// @summary Запросы по спецификации GraphQL
// @param payload body any true "тело запроса"
// @success 200 {object} any
// @failure 400 {object} http_errors.ErrorResponse
// @failure 401 {object} http_errors.ErrorResponse
// @failure 500 {object} http_errors.ErrorResponse
// @router /graphql [POST]
func RegisterGraphQL(router *echo.Echo, resolver *resolvers.Resolver, auther authorization.Auther) {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: resolver,
		Directives: generated.DirectiveRoot{
			Permission:        directives2.Permission,
			IgnoreReleaseDate: directives2.IgnoreReleaseDate,
			Protect:           directives2.Protect,
		},
	}))

	// Обработка паник
	srv.SetRecoverFunc(func(ctx context.Context, panicInstance interface{}) error {
		panicErr := err_const.FromPanic(panicInstance)

		logger.GetFromCtx(ctx).Error("Паника graphql", zap.Any("panic", panicInstance))

		return graphql.DefaultErrorPresenter(ctx, panicErr)
	})

	// Преобразования ошибок к необходимому виду
	srv.AroundResponses(handleResponse)

	/* Добавляем проверку аутентификации */
	router.Any("/graphql",
		echo.WrapHandler(srv),
		// Достаем токен из заголовка запроса
		authorization.WithJWTExtractor(),
		// На основе токена из контекста генерируем сессию
		authorization.WithAuth(auther),
	)
}
