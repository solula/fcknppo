package permissions

import (
	"slices"
	"waterfall-backend/internal/constants/err_const"
)

type Type string // Тип права доступа

const (
	read     = "_read"        // Право на чтение
	readSelf = "_read_self"   // Право на чтение своей записи (пользователя, комментария, т.д.)
	crt      = "_create"      // Право на создание
	upd      = "_update"      // Право на обновление
	updSelf  = "_update_self" // Право на обновление своей записи (пользователя, комментария, т.д.)
	del      = "_delete"      // Право на удаление
	delSelf  = "_delete_self" // Право на удаление своей записи (пользователя, комментария, т.д.)

	files    = "files"    // Определяет доступ к файлам
	releases = "releases" // Определяет доступ к релизам
	users    = "users"    // Определяет доступ к пользователям
	chapters = "chapters" // Определяет доступ к главам
	parts    = "parts"    // Определяет доступ к частям
	comments = "comments" // Определяет доступ к комментариям
)

const (
	FilesRead   Type = files + read
	FilesCreate Type = files + crt

	ReleasesRead   Type = releases + read
	ReleasesCreate Type = releases + crt
	ReleasesUpdate Type = releases + upd
	ReleasesDelete Type = releases + del

	UsersRead       Type = users + read
	UsersReadSelf   Type = users + readSelf
	UsersCreate     Type = users + crt
	UsersUpdate     Type = users + upd
	UsersUpdateSelf Type = users + updSelf
	UsersDelete     Type = users + del
	UsersDeleteSelf Type = users + delSelf

	ChaptersRead   Type = chapters + read
	ChaptersCreate Type = chapters + crt
	ChaptersUpdate Type = chapters + upd
	ChaptersDelete Type = chapters + del

	PartsRead   Type = parts + read
	PartsCreate Type = parts + crt
	PartsUpdate Type = parts + upd
	PartsDelete Type = parts + del

	CommentsRead   Type = comments + read
	CommentsCreate Type = comments + crt
	CommentsUpdate Type = comments + upd
	CommentsDelete Type = comments + del
)

func (t Type) Values() []string {
	return []string{
		string(FilesRead),
		string(FilesCreate),
		string(ReleasesRead),
		string(ReleasesCreate),
		string(ReleasesUpdate),
		string(ReleasesDelete),
		string(UsersRead),
		string(UsersReadSelf),
		string(UsersCreate),
		string(UsersUpdate),
		string(UsersUpdateSelf),
		string(UsersDelete),
		string(UsersDeleteSelf),
		string(ChaptersRead),
		string(ChaptersCreate),
		string(ChaptersUpdate),
		string(ChaptersDelete),
		string(PartsRead),
		string(PartsCreate),
		string(PartsUpdate),
		string(PartsDelete),
		string(CommentsRead),
		string(CommentsCreate),
		string(CommentsUpdate),
		string(CommentsDelete),
	}
}

func (t Type) Validate() error {
	if !slices.Contains(t.Values(), string(t)) {
		return err_const.ErrInvalidPermission
	}
	return nil
}
