package converters

import (
	"time"
	"waterfall-backend/internal/modules/domain/user/dto"
	"waterfall-backend/internal/modules/stores/db/ent"
)

func ToRoleDTO(model *ent.Role) *dto.Role {
	if model == nil {
		return nil
	}
	return &dto.Role{
		Id:           model.ID,
		Description:  model.Description,
		ReleaseDelay: time.Duration(model.ReleaseDelay * float64(time.Second)),
	}
}

func ToRoleDTOs(models ent.Roles) dto.Roles {
	if models == nil {
		return nil
	}
	dtms := make(dto.Roles, len(models))
	for i := range models {
		dtms[i] = ToRoleDTO(models[i])
	}
	return dtms
}
