package converters

import (
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/modules/services/fs/dto"
	"waterfall-backend/internal/modules/stores/db/ent"
)

func ToFileDTO(model *ent.File) *dto.File {
	if model == nil {
		return nil
	}

	return &dto.File{
		Uuid:        model.ID,
		Filename:    model.Filename,
		MIMEType:    model.MimeType,
		Description: model.Description,
		CreatorUuid: model.CreatorUUID,
		ObjectRef: &models.ObjectRef{
			Type: model.ObjectType,
			Ref:  model.ObjectRef,
		},
		Type:           model.Type,
		Temp:           model.Temp,
		SequenceNumber: model.SequenceNumber,
	}
}

func ToFilesDTOs(models ent.Files) dto.Files {
	if models == nil {
		return nil
	}
	dtms := make(dto.Files, len(models))
	for i := range models {
		dtms[i] = ToFileDTO(models[i])
	}
	return dtms
}
