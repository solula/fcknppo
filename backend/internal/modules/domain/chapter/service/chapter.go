package service

import (
	"context"
	"fmt"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/modules/domain/chapter/dto"
	"waterfall-backend/internal/modules/services/fs/i_fs"
	"waterfall-backend/internal/pkg/transaction"
)

type IChapterRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.Chapter, error)
	List(ctx context.Context) (dto.Chapters, error)
	NextUuid(ctx context.Context, uuid string) (*string, error)
	PrevUuid(ctx context.Context, uuid string) (*string, error)
	FirstUuid(ctx context.Context) (string, error)
	LastAvailableUuid(ctx context.Context) (string, error)

	Create(ctx context.Context, chapter *dto.ChapterCreate) (*dto.Chapter, error)
	Update(ctx context.Context, uuid string, chapter *dto.ChapterUpdate) (*dto.Chapter, error)
	Delete(ctx context.Context, uuid string) error

	transaction.TxRepo
}

type ChapterService struct {
	repo          IChapterRepo
	bucketCreator i_fs.IBucketCreator
}

func NewChapterService(repo IChapterRepo, bucketCreator i_fs.IBucketCreator) *ChapterService {
	return &ChapterService{
		repo:          repo,
		bucketCreator: bucketCreator,
	}
}

// GetByUuid получение записи
func (r *ChapterService) GetByUuid(ctx context.Context, uuid string) (*dto.Chapter, error) {
	chapter, err := r.repo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить главу: %w", err)
	}

	return chapter, nil
}

// List получение списка записей
func (r *ChapterService) List(ctx context.Context) (dto.Chapters, error) {
	entities, err := r.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список глав: %w", err)
	}

	return entities, nil
}

// NextUuid получение uuid следующей главы по uuid предыдущей
func (r *ChapterService) NextUuid(ctx context.Context, uuid string) (*string, error) {
	nextUuid, err := r.repo.NextUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить uuid следующей главы: %w", err)
	}

	return nextUuid, nil
}

// PrevUuid получение uuid предыдущей главы по uuid следующей
func (r *ChapterService) PrevUuid(ctx context.Context, uuid string) (*string, error) {
	prevUuid, err := r.repo.PrevUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить uuid предыдущей главы: %w", err)
	}

	return prevUuid, nil
}

// FirstUuid Получение uuid первой главы
func (r *ChapterService) FirstUuid(ctx context.Context) (string, error) {
	firstUuid, err := r.repo.FirstUuid(ctx)
	if err != nil {
		return "", fmt.Errorf("не удалось получить uuid первой главы: %w", err)
	}

	return firstUuid, nil
}

// LastAvailableUuid Получение uuid последней главы
func (r *ChapterService) LastAvailableUuid(ctx context.Context) (string, error) {
	lastUuid, err := r.repo.LastAvailableUuid(ctx)
	if err != nil {
		return "", fmt.Errorf("не удалось получить uuid последней главы: %w", err)
	}

	return lastUuid, nil
}

// Create создание
func (r *ChapterService) Create(ctx context.Context, chapter *dto.ChapterCreate) (*dto.Chapter, error) {
	var newChapter *dto.Chapter
	err := transaction.WithTx(ctx, r.repo, func(txRepo IChapterRepo) error {
		var err error
		newChapter, err = txRepo.Create(ctx, chapter)
		if err != nil {
			return fmt.Errorf("не удалось создать главу: %w", err)
		}

		err = r.bucketCreator.CreateBucket(ctx, &models.ObjectRef{
			Type: models.ObjectTypeChapters,
			Ref:  newChapter.Uuid,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return newChapter, nil
}

// Update обновление
func (r *ChapterService) Update(ctx context.Context, uuid string, chapterUpdate *dto.ChapterUpdate) (*dto.Chapter, error) {
	chapter, err := r.repo.Update(ctx, uuid, chapterUpdate)
	if err != nil {
		return nil, fmt.Errorf("не удалось обновить главу: %w", err)
	}

	return chapter, nil
}

// Delete удаление записи
func (r *ChapterService) Delete(ctx context.Context, uuid string) error {
	err := r.repo.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("не удалось удалить главу: %w", err)
	}

	return nil
}
