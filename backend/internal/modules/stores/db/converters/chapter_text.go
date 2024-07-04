package converters

import (
	"waterfall-backend/internal/modules/domain/chapter/dto"
	"waterfall-backend/internal/modules/stores/db/ent"
)

func ToChapterTextDTO(model *ent.ChapterText) *dto.ChapterText {
	if model == nil {
		return nil
	}

	return &dto.ChapterText{
		Uuid:        model.ID,
		ChapterUuid: model.ChapterUUID,
		Text:        model.Text,
	}
}

func ToChapterTextDTOs(models ent.ChapterTexts) dto.ChapterTexts {
	if models == nil {
		return nil
	}
	dtms := make(dto.ChapterTexts, len(models))
	for i := range models {
		dtms[i] = ToChapterTextDTO(models[i])
	}
	return dtms
}
