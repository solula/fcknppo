package service

import (
	"context"
	"fmt"
	"time"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/domain/release/dto"
)

type IReleaseRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.Release, error)
	GetInfoByUuid(ctx context.Context, uuid string) (*dto.Release, error)
	List(ctx context.Context) (dto.Releases, error)

	Create(ctx context.Context, release *dto.ReleaseCreate) (*dto.Release, error)
	Update(ctx context.Context, uuid string, dtm *dto.ReleaseUpdate) (*dto.Release, error)
	Delete(ctx context.Context, uuid string) error
}

type ReleaseService struct {
	repo IReleaseRepo
}

func NewReleaseService(repo IReleaseRepo) *ReleaseService {
	return &ReleaseService{
		repo: repo,
	}
}

// GetByUuid получение релиза объекта.
func (r *ReleaseService) GetByUuid(ctx context.Context, uuid string) (*dto.Release, error) {
	rls, err := r.repo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить релиз: %w", err)
	}

	return rls, nil
}

// GetInfoByUuid получение информации релиза (без указания выпускаемых объектов)
func (r *ReleaseService) GetInfoByUuid(ctx context.Context, uuid string) (*dto.Release, error) {
	rls, err := r.repo.GetInfoByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить информацию о релизе: %w", err)
	}

	return rls, nil
}

// GetAvailableAtByUuid получение даты, когда релиз будет доступен текущему пользователю
func (r *ReleaseService) GetAvailableAtByUuid(ctx context.Context, uuid string) (*time.Time, error) {
	ss, ok := session.GetFromCtx(ctx)
	if !ok {
		return nil, err_const.ErrMissingSession
	}

	rls, err := r.repo.GetInfoByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить информацию о релизе: %w", err)
	}

	availableAt := rls.ReleaseDate.Add(ss.ReleaseDelay)
	return &availableAt, nil
}

// List получение списка релизов
func (r *ReleaseService) List(ctx context.Context) (dto.Releases, error) {
	rls, err := r.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список релизов: %w", err)
	}

	return rls, nil
}

// Create создание записи в БД
func (r *ReleaseService) Create(ctx context.Context, release *dto.ReleaseCreate) (*dto.Release, error) {
	var chapterUuids, partUuids []string

	// Заполняем массивы Uuid-ов сущностей
	for _, objectRef := range release.ObjectRefs {
		switch objectRef.Type {
		case models.ObjectTypeChapters:
			chapterUuids = append(chapterUuids, objectRef.Ref)
		case models.ObjectTypeParts:
			partUuids = append(partUuids, objectRef.Ref)
		}
	}

	// Устанавливаем заполненные массивы в релиз
	release.ChapterUuids = chapterUuids
	release.PartUuids = partUuids

	newRelease, err := r.repo.Create(ctx, release)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать релиз: %w", err)
	}

	return newRelease, nil
}

// Update обновление записи
func (r *ReleaseService) Update(ctx context.Context, uuid string, release *dto.ReleaseUpdate) (*dto.Release, error) {
	newRelease, err := r.repo.Update(ctx, uuid, release)
	if err != nil {
		return nil, fmt.Errorf("не удалось обновить релиз: %w", err)
	}

	return newRelease, nil
}

// Delete удаление записи
func (r *ReleaseService) Delete(ctx context.Context, uuid string) error {
	err := r.repo.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("не удалось удалить релиз: %w", err)
	}

	return nil
}
