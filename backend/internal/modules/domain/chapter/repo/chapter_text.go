package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"strings"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/modules/domain/chapter/dto"
	"waterfall-backend/internal/modules/stores/db/converters"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/chaptertext"
	"waterfall-backend/internal/modules/stores/db/schema"
	"waterfall-backend/internal/modules/stores/db/utils"
)

type ChapterTextRepo struct {
	client *ent.Client
}

func NewChapterTextRepo(client *ent.Client) *ChapterTextRepo {
	return &ChapterTextRepo{
		client: client,
	}
}

// GetByUuid получение записи
func (r *ChapterTextRepo) GetByUuid(ctx context.Context, uuid string) (*dto.ChapterText, error) {
	dtm, err := r.client.ChapterText.Get(ctx, uuid)
	if err != nil {
		return nil, wrapChapterText(err)
	}

	return converters.ToChapterTextDTO(dtm), nil
}

// GetByChapterUuid получение текста главы по uuid главы
func (r *ChapterTextRepo) GetByChapterUuid(ctx context.Context, chapterUuid string) (*dto.ChapterText, error) {
	dtm, err := r.client.ChapterText.Query().
		Where(chaptertext.ChapterUUID(chapterUuid)).
		Only(ctx)
	if err != nil {
		return nil, wrapChapterText(err)
	}

	return converters.ToChapterTextDTO(dtm), nil
}

// Create создание
func (r *ChapterTextRepo) Create(ctx context.Context, chapterText *dto.ChapterTextCreate) (*dto.ChapterText, error) {
	newChapterText, err := r.client.ChapterText.Create().
		SetChapterID(chapterText.ChapterUuid).
		SetText(chapterText.Text).
		Save(ctx)
	if err != nil {
		return nil, wrapChapterText(err)
	}

	return converters.ToChapterTextDTO(newChapterText), nil
}

// Update обновление
func (r *ChapterTextRepo) Update(ctx context.Context, uuid string, chapterTextUpdate *dto.ChapterTextUpdate) (*dto.ChapterText, error) {
	updChapterText, err := r.client.ChapterText.UpdateOneID(uuid).
		SetText(chapterTextUpdate.Text).
		Save(ctx)
	if err != nil {
		return nil, wrapChapterText(err)
	}

	return converters.ToChapterTextDTO(updChapterText), nil
}

// Delete удаление записи
func (r *ChapterTextRepo) Delete(ctx context.Context, uuid string) error {
	err := r.client.ChapterText.DeleteOneID(uuid).Exec(schema.SkipSoftDelete(ctx))
	if err != nil {
		return wrapChapterText(err)
	}

	return nil
}

func wrapChapterText(err error) error {
	if ent.IsConstraintError(err) {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			// Код 23505 - нарушение уникальности
			if pqErr.Code == "23505" {
				if strings.Contains(pqErr.Constraint, chaptertext.FieldChapterUUID) {
					return fmt.Errorf("%w: текст этой главы уже существует", err_const.ErrUniqueConstraint)
				}
			}
		}
	}

	return utils.DefaultErrorWrapper(err)
}
