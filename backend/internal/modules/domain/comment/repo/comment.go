package repo

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"waterfall-backend/internal/modules/domain/comment/dto"
	"waterfall-backend/internal/modules/stores/db/converters"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/comment"
	"waterfall-backend/internal/modules/stores/db/schema"
	"waterfall-backend/internal/modules/stores/db/utils"
)

type CommentRepo struct {
	client *ent.Client
}

func NewCommentRepo(client *ent.Client) *CommentRepo {
	return &CommentRepo{
		client: client,
	}
}

func (r *CommentRepo) GetByUuid(ctx context.Context, uuid string) (*dto.Comment, error) {
	usr, err := r.client.Comment.Get(schema.SkipSoftDelete(ctx), uuid)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToCommentDTO(usr), nil
}

func (r *CommentRepo) ListByUuids(ctx context.Context, uuids []string) (dto.Comments, error) {
	comments, err := r.client.Comment.Query().
		Where(comment.IDIn(uuids...)).
		All(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToCommentDTOs(comments), nil
}

func (r *CommentRepo) ListByChapter(ctx context.Context, chapterUuid string) (dto.Comments, error) {
	comments, err := r.client.Comment.Query().
		Where(comment.ChapterUUID(chapterUuid)).
		Order(comment.ByCreatedAt(sql.OrderAsc())).
		All(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToCommentDTOs(comments), nil
}

func (r *CommentRepo) Create(ctx context.Context, dtm *dto.CommentCreate) (*dto.Comment, error) {
	usr, err := r.client.Comment.Create().
		SetText(dtm.Text).
		SetAuthorID(dtm.AuthorUuid).
		SetNillableParentID(dtm.ParentUuid).
		SetNillableChapterID(dtm.ChapterUuid).
		Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToCommentDTO(usr), nil
}

func (r *CommentRepo) Update(ctx context.Context, uuid string, dtm *dto.CommentUpdate) (*dto.Comment, error) {
	usr, err := r.client.Comment.UpdateOneID(uuid).
		SetText(dtm.Text).
		Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToCommentDTO(usr), nil
}

func (r *CommentRepo) Delete(ctx context.Context, uuid string) error {
	err := r.client.Comment.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return wrap(err)
	}

	return nil
}

func (r *CommentRepo) Restore(ctx context.Context, uuid string) (*dto.Comment, error) {
	usr, err := r.client.Comment.UpdateOneID(uuid).ClearDeletedAt().Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToCommentDTO(usr), nil
}

func wrap(err error) error {
	return utils.DefaultErrorWrapper(err)
}
