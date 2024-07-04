package http

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"waterfall-backend/internal/models/access"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/io/http/api"
	"waterfall-backend/internal/modules/io/http/auth"
	"waterfall-backend/internal/modules/io/http/public"
	"waterfall-backend/internal/pkg/http/middlewares/authorization"
	"waterfall-backend/internal/pkg/routers"
)

type Routers struct {
	fx.Out

	Public public.Router
	Api    api.Router
	Auth   auth.Router
}

// NewHTTP создает http все необходимые роутеры
func NewHTTP(server *echo.Echo, cfg config.Config, auther authorization.Auther) (Routers, error) {
	return Routers{
		Public: public.InitRouter(cfg, &routers.EchoRouter{
			Router: server,
		}),
		Api: api.InitRouter(&routers.EchoRouter{
			Router:         server.Group("/api"),
			DefaultWrapper: permissionWrapper,
		}, auther),
		Auth: auth.InitRouter(&routers.EchoRouter{
			Router: server.Group("/auth"),
		}),
	}, nil
}

func permissionWrapper(handler echo.HandlerFunc, opts ...routers.RouteOption) echo.HandlerFunc {
	var perms []permissions.Type
	for _, opt := range opts {
		switch opt.Ident() {
		case permissions.Ident{}:
			perms = append(perms, opt.Value().(permissions.Type))
		}
	}

	if len(perms) == 0 {
		return handler
	}

	if len(perms) > 1 {
		panic("указание нескольких прав доступа не поддерживаются")
	}
	permission := perms[0]

	return func(c echo.Context) error {
		if err := access.CheckPermissionsFromCtx(c.Request().Context(), permission); err != nil {
			return err
		}
		return handler(c)
	}
}
