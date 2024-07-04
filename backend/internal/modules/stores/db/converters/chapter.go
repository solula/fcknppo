package converters

import (
	"waterfall-backend/internal/modules/domain/chapter/dto"
	"waterfall-backend/internal/modules/stores/db/ent"
)

func ToChapterDTO(model *ent.Chapter) *dto.Chapter {
	if model == nil {
		return nil
	}

	return &dto.Chapter{
		Uuid:        model.ID,
		Number:      model.Number,
		Title:       model.Title,
		PartUuid:    model.PartUUID,
		ReleaseUuid: model.ReleaseUUID,
	}
}

func ToChapterDTOs(models ent.Chapters) dto.Chapters {
	if models == nil {
		return nil
	}
	dtms := make(dto.Chapters, len(models))
	for i := range models {
		dtms[i] = ToChapterDTO(models[i])
	}
	return dtms
}
