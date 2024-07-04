package converters

import (
	"waterfall-backend/internal/modules/domain/user/dto"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/pkg/unique"
)

func ToPermissionDTO(model *ent.Permission) *dto.Permission {
	if model == nil {
		return nil
	}
	return &dto.Permission{
		Id:          model.ID,
		Description: model.Description,
	}
}

func ToPermissionDTOs(models ent.Permissions) dto.Permissions {
	if models == nil {
		return nil
	}
	dtms := make(dto.Permissions, len(models))
	for i := range models {
		dtms[i] = ToPermissionDTO(models[i])
	}
	return dtms
}

func ToPermissionDTOsByRoles(roles ent.Roles) dto.Permissions {
	var perms dto.Permissions
	for _, role := range roles {
		perms = append(perms, ToPermissionDTOs(role.Edges.Permissions)...)
	}

	perms = unique.Unique(perms)
	return perms
}
