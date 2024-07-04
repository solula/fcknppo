package comment

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/domain/comment/repo"
	"waterfall-backend/internal/modules/domain/comment/service"
)

var ServiceModule = fx.Module("comment",
	service.Module,
	repo.Module,

	fx.Provide(
		fx.Annotate(
			func(r *repo.CommentRepo) *repo.CommentRepo { return r },
			fx.As(new(service.ICommentRepo)),
		),
	),
)
