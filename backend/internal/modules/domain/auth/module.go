package auth

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/domain/auth/http"
	"waterfall-backend/internal/modules/domain/auth/service"
	"waterfall-backend/internal/modules/domain/user/repo"
	email_serv "waterfall-backend/internal/modules/services/email/service"
	fs_serv "waterfall-backend/internal/modules/services/fs/service"
	"waterfall-backend/internal/pkg/http/middlewares/authorization"
)

var ServiceModule = fx.Module("auth-service",
	service.Module,

	fx.Provide(
		fx.Annotate(
			func(r *repo.UserRepo) *repo.UserRepo { return r },
			fx.As(new(service.IUserRepo)),
		),
		fx.Annotate(
			func(r *repo.RoleRepo) *repo.RoleRepo { return r },
			fx.As(new(service.IRoleRepo)),
		),
		fx.Annotate(
			func(r *fs_serv.FileStorageService) *fs_serv.FileStorageService { return r },
			fx.As(new(service.IFileStorageService)),
		),
		fx.Annotate(
			func(r *email_serv.EmailService) *email_serv.EmailService { return r },
			fx.As(new(service.IEmailService)),
		),

		// Аннотируем сервис как интерфейс Auther
		fx.Annotate(
			func(r *service.AuthService) *service.AuthService { return r },
			fx.As(new(authorization.Auther)),
		),
	),
)

var HTTPModule = fx.Module("auth",
	http.Module,
)
