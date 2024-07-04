package repo

import (
	"context"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/modules/domain/release/dto"
	"waterfall-backend/internal/modules/stores/db/converters"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/release"
	"waterfall-backend/internal/modules/stores/db/utils"
)

type ReleaseRepo struct {
	client *ent.Client
}

func NewReleaseRepo(client *ent.Client) *ReleaseRepo {
	return &ReleaseRepo{
		client: client,
	}
}

// GetByUuid получение релиза объекта
func (r *ReleaseRepo) GetByUuid(ctx context.Context, uuid string) (*dto.Release, error) {
	rls, err := r.client.Release.Query().
		Where(release.ID(uuid)).
		WithParts().
		WithChapters().
		Only(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToReleaseDTO(rls), nil
}

// GetInfoByUuid получение информации релиза (без указания выпускаемых объектов)
func (r *ReleaseRepo) GetInfoByUuid(ctx context.Context, uuid string) (*dto.Release, error) {
	rls, err := r.client.Release.Query().
		Where(release.ID(uuid)).
		Only(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToReleaseInfoDTO(rls), nil
}

// List получение списка релизов
func (r *ReleaseRepo) List(ctx context.Context) (dto.Releases, error) {
	rls, err := r.client.Release.Query().
		WithParts().
		WithChapters().
		All(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToReleaseDTOs(rls), nil
}

// Create создание релиза
func (r *ReleaseRepo) Create(ctx context.Context, release *dto.ReleaseCreate) (*dto.Release, error) {
	newRelease, err := r.client.Release.Create().
		SetReleaseDate(release.ReleaseDate).
		SetDescription(release.Description).
		AddChapterIDs(release.ChapterUuids...).
		AddPartIDs(release.PartUuids...).
		Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToReleaseInfoDTO(newRelease), nil
}

// Update обновление релиза
func (r *ReleaseRepo) Update(ctx context.Context, uuid string, dtm *dto.ReleaseUpdate) (*dto.Release, error) {
	updRelease, err := r.client.Release.UpdateOneID(uuid).
		SetReleaseDate(dtm.ReleaseDate).
		SetDescription(dtm.Description).
		Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToReleaseInfoDTO(updRelease), nil
}

// Delete удаление релиза
func (r *ReleaseRepo) Delete(ctx context.Context, uuid string) error {
	n, err := r.client.Release.Delete().
		Where(release.ID(uuid)).
		Exec(ctx)
	if err != nil {
		return wrap(err)
	}
	if n == 0 {
		return err_const.ErrNotFound
	}

	return nil
}

func wrap(err error) error {
	return utils.DefaultErrorWrapper(err)
}
