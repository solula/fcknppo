package service

import (
	"context"
	"fmt"
	"waterfall-backend/internal/modules/domain/chapter/dto"
)

type IChapterTextRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.ChapterText, error)
	GetByChapterUuid(ctx context.Context, chapterUuid string) (*dto.ChapterText, error)

	Create(ctx context.Context, chapterText *dto.ChapterTextCreate) (*dto.ChapterText, error)
	Update(ctx context.Context, uuid string, chapterTextUpdate *dto.ChapterTextUpdate) (*dto.ChapterText, error)
	Delete(ctx context.Context, uuid string) error
}

type ChapterTextService struct {
	repo IChapterTextRepo
}

func NewChapterTextService(repo IChapterTextRepo) *ChapterTextService {
	return &ChapterTextService{
		repo: repo,
	}
}

// GetByUuid получение записи
func (r *ChapterTextService) GetByUuid(ctx context.Context, uuid string) (*dto.ChapterText, error) {
	chapter, err := r.repo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить текст главы: %w", err)
	}

	return chapter, nil
}

// GetByChapterUuid получение текста главы по uuid главы
func (r *ChapterTextService) GetByChapterUuid(ctx context.Context, chapterUuid string) (*dto.ChapterText, error) {
	chapter, err := r.repo.GetByChapterUuid(ctx, chapterUuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить текст главы: %w", err)
	}

	return chapter, nil
}

// Create создание
func (r *ChapterTextService) Create(ctx context.Context, chapterText *dto.ChapterTextCreate) (*dto.ChapterText, error) {
	newChapter, err := r.repo.Create(ctx, chapterText)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать текст главы: %w", err)
	}

	return newChapter, nil
}

// Update обновление
func (r *ChapterTextService) Update(ctx context.Context, uuid string, chapterTextUpdate *dto.ChapterTextUpdate) (*dto.ChapterText, error) {
	chapter, err := r.repo.Update(ctx, uuid, chapterTextUpdate)
	if err != nil {
		return nil, fmt.Errorf("не удалось обновить текст главы: %w", err)
	}

	return chapter, nil
}

// Delete удаление записи
func (r *ChapterTextService) Delete(ctx context.Context, uuid string) error {
	err := r.repo.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("не удалось удалить текст главы: %w", err)
	}

	return nil
}
