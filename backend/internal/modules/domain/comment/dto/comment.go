package dto

import (
	"fmt"
	"time"
	"waterfall-backend/internal/constants/err_const"
)

// Comment комментарий
type Comment struct {
	Uuid        string    // Uuid комментария
	CreatedAt   time.Time // Дата создания
	UpdatedAt   time.Time // Дата обновления
	Text        string    // Текст комментария
	AuthorUuid  string    // Автор комментария
	ParentUuid  *string   // Родительский комментарий
	ChapterUuid *string   // Глава, которой этот комментарий принадлежит
}

// Comments комментарии
type Comments []*Comment

// CommentCreate модель создания комментария
type CommentCreate struct {
	Text        string  // Текст комментария
	AuthorUuid  string  // Автор комментария
	ParentUuid  *string // Родительский комментарий
	ChapterUuid *string // Глава, которой этот комментарий принадлежит
}

// CommentUpdate модель обновления комментария
type CommentUpdate struct {
	Text string // Текст комментария
}

func (c *CommentCreate) Validate() error {
	if c.ChapterUuid == nil {
		return fmt.Errorf("%w: комментарий должен принадлежать хотя бы к одной сущности", err_const.ErrValidate)
	}

	if c.Text == "" {
		return fmt.Errorf("%w: текст комментария не может быть пустым", err_const.ErrValidate)
	}

	return nil
}
