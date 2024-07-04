package service

import (
	"context"
	"fmt"
	"slices"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/domain/comment/dto"
	"waterfall-backend/internal/pkg/transaction"
)

type ICommentRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.Comment, error)
	ListByUuids(ctx context.Context, uuids []string) (dto.Comments, error)
	ListByChapter(ctx context.Context, commentUuid string) (dto.Comments, error)
	Create(ctx context.Context, dtm *dto.CommentCreate) (*dto.Comment, error)
	Update(ctx context.Context, uuid string, dtm *dto.CommentUpdate) (*dto.Comment, error)
	Delete(ctx context.Context, uuid string) error
	Restore(ctx context.Context, uuid string) (*dto.Comment, error)

	transaction.TxRepo
}

type CommentService struct {
	repo ICommentRepo
}

func NewCommentService(repo ICommentRepo) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (r *CommentService) GetByUuid(ctx context.Context, uuid string) (*dto.Comment, error) {
	comment, err := r.repo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить комментарий: %w", err)
	}

	return comment, nil
}

func (r *CommentService) ListByChapter(ctx context.Context, commentUuid string) (dto.Comments, error) {
	comments, err := r.repo.ListByChapter(ctx, commentUuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить комментарии: %w", err)
	}

	return comments, nil
}

func (r *CommentService) ListByUuids(ctx context.Context, uuids []string) (dto.Comments, error) {
	comments, err := r.repo.ListByUuids(ctx, uuids)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить комментарии: %w", err)
	}

	return comments, nil
}

func (r *CommentService) Create(ctx context.Context, comment *dto.CommentCreate) (*dto.Comment, error) {
	ss, ok := session.GetFromCtx(ctx)
	if !ok {
		return nil, err_const.ErrMissingSession
	}
	comment.AuthorUuid = ss.UserUuid

	err := comment.Validate()
	if err != nil {
		return nil, err
	}

	newComment, err := r.repo.Create(ctx, comment)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать комментарий: %w", err)
	}

	return newComment, nil
}

func (r *CommentService) Update(ctx context.Context, uuid string, dtm *dto.CommentUpdate) (*dto.Comment, error) {
	ss, ok := session.GetFromCtx(ctx)
	if !ok {
		return nil, err_const.ErrMissingSession
	}
	comment, err := r.repo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить комментарий: %w", err)
	}

	// Если пользователь (на админ) не является создателем комментария -> запрещаем редактирование
	if !slices.Contains(ss.Roles, roles.Admin) && comment.AuthorUuid != ss.UserUuid {
		return nil, err_const.ErrAccessDenied
	}

	updComment, err := r.repo.Update(ctx, uuid, dtm)
	if err != nil {
		return nil, fmt.Errorf("не удалось обновить комментарий: %w", err)
	}

	return updComment, nil
}

func (r *CommentService) Delete(ctx context.Context, uuid string) error {
	ss, ok := session.GetFromCtx(ctx)
	if !ok {
		return err_const.ErrMissingSession
	}
	comment, err := r.repo.GetByUuid(ctx, uuid)
	if err != nil {
		return fmt.Errorf("не удалось получить комментарий: %w", err)
	}

	// Если пользователь (на админ) не является создателем комментария -> запрещаем удаление
	if !slices.Contains(ss.Roles, roles.Admin) && comment.AuthorUuid != ss.UserUuid {
		return err_const.ErrAccessDenied
	}

	err = r.repo.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("не удалось удалить комментарий: %w", err)
	}

	return nil
}

func (r *CommentService) Restore(ctx context.Context, uuid string) (*dto.Comment, error) {
	comment, err := r.repo.Restore(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось восстановить комментарий: %w", err)
	}

	return comment, nil
}
