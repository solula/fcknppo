package domain

import (
	chapter_serv "waterfall-backend/internal/modules/domain/chapter/service"
	comment_serv "waterfall-backend/internal/modules/domain/comment/service"
	part_serv "waterfall-backend/internal/modules/domain/part/service"
	release_serv "waterfall-backend/internal/modules/domain/release/service"
	user_serv "waterfall-backend/internal/modules/domain/user/service"
	fs_serv "waterfall-backend/internal/modules/services/fs/service"
)

type Services struct {
	Release     *release_serv.ReleaseService
	FileStorage *fs_serv.FileStorageService

	User     *user_serv.UserService
	Chapter  *chapter_serv.ChapterService
	Part     *part_serv.PartService
	Comments *comment_serv.CommentService
}

func NewServices(
	release *release_serv.ReleaseService,
	fs *fs_serv.FileStorageService,
	user *user_serv.UserService,
	chapter *chapter_serv.ChapterService,
	part *part_serv.PartService,
	comment *comment_serv.CommentService,

) *Services {
	return &Services{
		Release:     release,
		FileStorage: fs,
		Chapter:     chapter,
		Part:        part,
		User:        user,
		Comments:    comment,
	}
}
