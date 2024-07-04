package service

import (
	"context"
	"fmt"
	"waterfall-backend/internal/models"
	dto2 "waterfall-backend/internal/modules/domain/chapter/dto"
	"waterfall-backend/internal/modules/domain/part/dto"
	"waterfall-backend/internal/modules/services/fs/i_fs"
	"waterfall-backend/internal/pkg/transaction"
)

type IPartRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.Part, error)
	List(ctx context.Context) (dto.Parts, error)
	ListChaptersByUuid(ctx context.Context, uuid string) (dto2.Chapters, error)

	Create(ctx context.Context, part *dto.PartCreate) (*dto.Part, error)
	Update(ctx context.Context, uuid string, part *dto.PartUpdate) (*dto.Part, error)
	Delete(ctx context.Context, uuid string) error

	transaction.TxRepo
}

type PartService struct {
	repo          IPartRepo
	bucketCreator i_fs.IBucketCreator
}

func NewPartService(repo IPartRepo, bucketCreator i_fs.IBucketCreator) *PartService {
	return &PartService{
		repo:          repo,
		bucketCreator: bucketCreator,
	}
}

// GetByUuid получение записи
func (r *PartService) GetByUuid(ctx context.Context, uuid string) (*dto.Part, error) {
	part, err := r.repo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить часть: %w", err)
	}

	return part, nil
}

// List получение списка записей
func (r *PartService) List(ctx context.Context) (dto.Parts, error) {
	entities, err := r.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список частей: %w", err)
	}

	return entities, nil
}

// ListChaptersByUuid получение списка глав по части
func (r *PartService) ListChaptersByUuid(ctx context.Context, uuid string) (dto2.Chapters, error) {
	entities, err := r.repo.ListChaptersByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список глав: %w", err)
	}

	return entities, nil
}

// Create создание записи в БД
func (r *PartService) Create(ctx context.Context, part *dto.PartCreate) (*dto.Part, error) {
	var newPart *dto.Part
	err := transaction.WithTx(ctx, r.repo, func(txRepo IPartRepo) error {
		var err error
		newPart, err = txRepo.Create(ctx, part)
		if err != nil {
			return fmt.Errorf("не удалось создать часть: %w", err)
		}

		err = r.bucketCreator.CreateBucket(ctx, &models.ObjectRef{
			Type: models.ObjectTypeParts,
			Ref:  newPart.Uuid,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return newPart, nil
}

// Update обновление записи
func (r *PartService) Update(ctx context.Context, uuid string, part *dto.PartUpdate) (*dto.Part, error) {
	newPart, err := r.repo.Update(ctx, uuid, part)
	if err != nil {
		return nil, fmt.Errorf("не удалось обновить часть: %w", err)
	}

	return newPart, nil
}

// Delete удаление записи
func (r *PartService) Delete(ctx context.Context, uuid string) error {
	err := r.repo.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("не удалось удалить часть: %w", err)
	}

	return nil
}
