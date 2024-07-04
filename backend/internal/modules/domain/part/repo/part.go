package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"strings"
	"time"
	"waterfall-backend/internal/constants/err_const"
	dto2 "waterfall-backend/internal/modules/domain/chapter/dto"
	"waterfall-backend/internal/modules/domain/part/dto"
	converters2 "waterfall-backend/internal/modules/stores/db/converters"
	ent2 "waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/chapter"
	"waterfall-backend/internal/modules/stores/db/ent/part"
	"waterfall-backend/internal/modules/stores/db/schema"
	"waterfall-backend/internal/modules/stores/db/utils"
)

type PartRepo struct {
	client *ent2.Client
}

func NewPartRepo(client *ent2.Client) *PartRepo {
	return &PartRepo{
		client: client,
	}
}

// GetByUuid получение записи
func (r *PartRepo) GetByUuid(ctx context.Context, uuid string) (*dto.Part, error) {
	dtm, err := r.client.Part.Get(ctx, uuid)
	if err != nil {
		return nil, wrap(err)
	}

	return converters2.ToPartDTO(dtm), nil
}

// List получение списка записей
func (r *PartRepo) List(ctx context.Context) (dto.Parts, error) {
	entities, err := r.client.Part.Query().All(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters2.ToPartDTOs(entities), nil
}

// ListChaptersByUuid получение списка глав по части
func (r *PartRepo) ListChaptersByUuid(ctx context.Context, uuid string) (dto2.Chapters, error) {
	entities, err := r.client.Chapter.Query().
		Where(chapter.PartUUID(uuid)).
		All(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters2.ToChapterDTOs(entities), nil
}

// Create создание записи
func (r *PartRepo) Create(ctx context.Context, part *dto.PartCreate) (*dto.Part, error) {
	newPart, err := r.client.Part.Create().
		SetNumber(part.Number).
		SetTitle(part.Title).
		SetNillableAnnotation(part.Annotation).
		Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters2.ToPartDTO(newPart), nil
}

// Update обновление
func (r *PartRepo) Update(ctx context.Context, uuid string, part *dto.PartUpdate) (*dto.Part, error) {
	updPart, err := r.client.Part.UpdateOneID(uuid).
		SetNumber(part.Number).
		SetTitle(part.Title).
		SetNillableAnnotation(part.Annotation).
		Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters2.ToPartDTO(updPart), nil
}

// Delete удаление записи
func (r *PartRepo) Delete(ctx context.Context, uuid string) error {
	err := r.client.Part.UpdateOneID(uuid).
		SetDeletedAt(time.Now()).
		ClearNumber().
		Exec(schema.SkipSoftDelete(ctx))
	if err != nil {
		return wrap(err)
	}

	return nil
}

func wrap(err error) error {
	if ent2.IsConstraintError(err) {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			// Код 23505 - нарушение уникальности
			if pqErr.Code == "23505" {
				if strings.Contains(pqErr.Constraint, part.FieldNumber) {
					return fmt.Errorf("%w: часть с таким номером уже существует", err_const.ErrUniqueConstraint)
				}
			}
		}
	}

	return utils.DefaultErrorWrapper(err)
}
