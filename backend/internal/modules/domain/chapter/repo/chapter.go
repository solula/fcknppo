package repo

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"slices"
	"strings"
	"time"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/access"
	"waterfall-backend/internal/modules/domain/chapter/dto"
	"waterfall-backend/internal/modules/stores/db/converters"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/chapter"
	"waterfall-backend/internal/modules/stores/db/ent/part"
	"waterfall-backend/internal/modules/stores/db/schema"
	"waterfall-backend/internal/modules/stores/db/utils"
)

type ChapterRepo struct {
	client *ent.Client
}

func NewChapterRepo(client *ent.Client) *ChapterRepo {
	return &ChapterRepo{
		client: client,
	}
}

// GetByUuid получение записи
func (r *ChapterRepo) GetByUuid(ctx context.Context, uuid string) (*dto.Chapter, error) {
	dtm, err := r.client.Chapter.Get(ctx, uuid)
	if err != nil {
		return nil, wrapChapter(err)
	}

	return converters.ToChapterDTO(dtm), nil
}

// List получение списка записей
func (r *ChapterRepo) List(ctx context.Context) (dto.Chapters, error) {
	entities, err := r.client.Chapter.Query().All(ctx)
	if err != nil {
		return nil, wrapChapter(err)
	}

	return converters.ToChapterDTOs(entities), nil
}

// NextUuid получение uuid следующей главы по uuid предыдущей
func (r *ChapterRepo) NextUuid(ctx context.Context, uuid string) (*string, error) {
	type tNextUuid struct {
		Uuid     string  `json:"uuid"`
		NextUuid *string `json:"next_uuid"`
	}

	var nextUuids []tNextUuid
	err := r.client.Chapter.Query().Modify(func(s *sql.Selector) {
		p := sql.Table(part.Table).As("p")
		s.Select(s.C(chapter.FieldID)).
			AppendSelectExprAs(sql.Window(func(b *sql.Builder) {
				b.S("lead(")
				b.Ident(s.C(chapter.FieldID))
				b.S(")")
			}).OrderBy(p.C(part.FieldNumber), s.C(chapter.FieldNumber)),
				"next_uuid",
			).
			Join(p).
			On(s.C(chapter.FieldPartUUID), p.C(part.FieldID))
	}).Scan(access.SetIgnoreReleaseDate(ctx), &nextUuids)
	if err != nil {
		return nil, wrapChapter(err)
	}

	nextUuidInd := slices.IndexFunc(nextUuids, func(nextUuid tNextUuid) bool {
		return nextUuid.Uuid == uuid
	})
	nextUuid := nextUuids[nextUuidInd].NextUuid

	return nextUuid, nil
}

// PrevUuid получение uuid предыдущей главы по uuid следующей
func (r *ChapterRepo) PrevUuid(ctx context.Context, uuid string) (*string, error) {
	type tPrevUuid struct {
		Uuid     string  `json:"uuid"`
		PrevUuid *string `json:"prev_uuid"`
	}

	var prevUuids []tPrevUuid
	err := r.client.Chapter.Query().Modify(func(s *sql.Selector) {
		p := sql.Table(part.Table).As("p")
		s.Select(s.C(chapter.FieldID)).
			AppendSelectExprAs(sql.Window(func(b *sql.Builder) {
				b.S("lag(")
				b.Ident(s.C(chapter.FieldID))
				b.S(")")
			}).OrderBy(p.C(part.FieldNumber), s.C(chapter.FieldNumber)),
				"prev_uuid",
			).
			Join(p).
			On(s.C(chapter.FieldPartUUID), p.C(part.FieldID))
	}).Scan(access.SetIgnoreReleaseDate(ctx), &prevUuids)
	if err != nil {
		return nil, wrapChapter(err)
	}

	prevUuidInd := slices.IndexFunc(prevUuids, func(prevUuid tPrevUuid) bool {
		return prevUuid.Uuid == uuid
	})
	prevUuid := prevUuids[prevUuidInd].PrevUuid

	return prevUuid, nil
}

// FirstUuid получение uuid первой главы
func (r *ChapterRepo) FirstUuid(ctx context.Context) (string, error) {
	dtm, err := r.client.Chapter.Query().
		Select(chapter.FieldID).
		WithPart().
		Order(
			chapter.ByPartField(part.FieldNumber, sql.OrderAsc()),
			chapter.ByNumber(sql.OrderAsc()),
		).
		First(ctx)
	if err != nil {
		return "", wrapChapter(err)
	}

	return dtm.ID, nil
}

// LastAvailableUuid получение uuid последней доступной главы
func (r *ChapterRepo) LastAvailableUuid(ctx context.Context) (string, error) {
	dtm, err := r.client.Chapter.Query().
		Select(chapter.FieldID).
		WithPart().
		Order(
			chapter.ByPartField(part.FieldNumber, sql.OrderDesc()),
			chapter.ByNumber(sql.OrderDesc()),
		).
		First(ctx)
	if err != nil {
		return "", wrapChapter(err)
	}

	return dtm.ID, nil
}

// Create создание записи
func (r *ChapterRepo) Create(ctx context.Context, chapter *dto.ChapterCreate) (*dto.Chapter, error) {
	newChapter, err := r.client.Chapter.Create().
		SetNumber(chapter.Number).
		SetTitle(chapter.Title).
		SetPartID(chapter.PartUuid).
		Save(ctx)
	if err != nil {
		return nil, wrapChapter(err)
	}

	return converters.ToChapterDTO(newChapter), nil
}

// Update обновление
func (r *ChapterRepo) Update(ctx context.Context, uuid string, chapter *dto.ChapterUpdate) (*dto.Chapter, error) {
	updChapter, err := r.client.Chapter.UpdateOneID(uuid).
		SetNumber(chapter.Number).
		SetTitle(chapter.Title).
		SetPartID(chapter.PartUuid).
		Save(ctx)
	if err != nil {
		return nil, wrapChapter(err)
	}

	return converters.ToChapterDTO(updChapter), nil
}

// Delete удаление записи
func (r *ChapterRepo) Delete(ctx context.Context, uuid string) error {
	err := r.client.Chapter.UpdateOneID(uuid).
		SetDeletedAt(time.Now()).
		ClearNumber().
		Exec(schema.SkipSoftDelete(ctx))
	if err != nil {
		return wrapChapter(err)
	}

	return nil
}

func wrapChapter(err error) error {
	if ent.IsConstraintError(err) {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			// Код 23505 - нарушение уникальности
			if pqErr.Code == "23505" {
				if strings.Contains(pqErr.Constraint, chapter.FieldNumber) && strings.Contains(pqErr.Constraint, chapter.FieldPartUUID) {
					return fmt.Errorf("%w: глава с таким номером в этой части уже существует", err_const.ErrUniqueConstraint)
				}
			}
		}
	}

	return utils.DefaultErrorWrapper(err)
}
