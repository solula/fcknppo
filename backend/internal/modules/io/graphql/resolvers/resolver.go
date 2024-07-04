package resolvers

import (
	"go.uber.org/fx"
	chapter_serv "waterfall-backend/internal/modules/domain/chapter/service"
	comment_serv "waterfall-backend/internal/modules/domain/comment/service"
	part_serv "waterfall-backend/internal/modules/domain/part/service"
	release_serv "waterfall-backend/internal/modules/domain/release/service"
	user_serv "waterfall-backend/internal/modules/domain/user/service"
	fs_serv "waterfall-backend/internal/modules/services/fs/service"
	sched_serv "waterfall-backend/internal/modules/services/scheduler/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Services struct {
	fx.In

	Release     *release_serv.ReleaseService
	FileStorage *fs_serv.FileStorageService
	Scheduler   *sched_serv.SchedulerService

	User        *user_serv.UserService
	Chapter     *chapter_serv.ChapterService
	ChapterText *chapter_serv.ChapterTextService
	Part        *part_serv.PartService
	Comment     *comment_serv.CommentService
}

type Resolver struct {
	services *Services
}

func NewResolver(services Services) *Resolver {
	return &Resolver{services: &services}
}
