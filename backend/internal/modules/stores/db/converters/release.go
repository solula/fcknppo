package converters

import (
	"waterfall-backend/internal/modules/domain/release/dto"
	"waterfall-backend/internal/modules/stores/db/ent"
)

func ToReleaseDTO(model *ent.Release) *dto.Release {
	if model == nil {
		return nil
	}

	return &dto.Release{
		Uuid:        model.ID,
		ReleaseDate: model.ReleaseDate,
		Description: model.Description,
		Releasables: ToReleasables(model),
	}
}

func ToReleaseInfoDTO(model *ent.Release) *dto.Release {
	if model == nil {
		return nil
	}

	return &dto.Release{
		Uuid:        model.ID,
		ReleaseDate: model.ReleaseDate,
		Description: model.Description,
	}
}

func ToReleasables(model *ent.Release) dto.Releasables {
	if model == nil {
		return nil
	}

	// Обходим все выпущенные сущности и формируем общий список
	var releasables dto.Releasables
	for _, part := range model.Edges.Parts {
		releasables = append(releasables, ToPartDTO(part))
	}
	for _, chapter := range model.Edges.Chapters {
		releasables = append(releasables, ToChapterDTO(chapter))
	}

	return releasables
}

func ToReleaseDTOs(models ent.Releases) dto.Releases {
	if models == nil {
		return nil
	}
	dtms := make(dto.Releases, len(models))
	for i := range models {
		dtms[i] = ToReleaseDTO(models[i])
	}
	return dtms
}
