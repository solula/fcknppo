package converters

import (
	"waterfall-backend/internal/modules/domain/comment/dto"
	"waterfall-backend/internal/modules/stores/db/ent"
)

func ToCommentDTO(model *ent.Comment) *dto.Comment {
	if model == nil {
		return nil
	}

	return &dto.Comment{
		Uuid:        model.ID,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
		Text:        model.Text,
		AuthorUuid:  model.AuthorUUID,
		ParentUuid:  model.ParentUUID,
		ChapterUuid: model.ChapterUUID,
	}
}

func ToCommentDTOs(models ent.Comments) dto.Comments {
	if models == nil {
		return nil
	}
	dtms := make(dto.Comments, len(models))
	for i := range models {
		dtms[i] = ToCommentDTO(models[i])
	}
	return dtms
}
