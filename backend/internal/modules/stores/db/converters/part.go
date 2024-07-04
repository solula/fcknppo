package converters

import (
	"waterfall-backend/internal/modules/domain/part/dto"
	"waterfall-backend/internal/modules/stores/db/ent"
)

func ToPartDTO(model *ent.Part) *dto.Part {
	if model == nil {
		return nil
	}

	return &dto.Part{
		Uuid:        model.ID,
		Number:      model.Number,
		Title:       model.Title,
		Annotation:  model.Annotation,
		ReleaseUuid: model.ReleaseUUID,
	}
}

func ToPartDTOs(models ent.Parts) dto.Parts {
	if models == nil {
		return nil
	}
	dtms := make(dto.Parts, len(models))
	for i := range models {
		dtms[i] = ToPartDTO(models[i])
	}
	return dtms
}
