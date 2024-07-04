package user

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/domain/user/repo"
	"waterfall-backend/internal/modules/domain/user/service"
)

var ServiceModule = fx.Module("user",
	service.Module,
	repo.Module,

	fx.Provide(
		fx.Annotate(
			func(r *repo.UserRepo) *repo.UserRepo { return r },
			fx.As(new(service.IUserRepo)),
		),
	),
)
