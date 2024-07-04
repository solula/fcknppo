package converters

import (
	"waterfall-backend/internal/modules/domain/user/dto"
	"waterfall-backend/internal/modules/stores/db/ent"
)

func ToUserDTO(model *ent.User) *dto.User {
	if model == nil {
		return nil
	}
	return &dto.User{
		Uuid:          model.ID,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
		DeletedAt:     model.DeletedAt,
		Email:         model.Email,
		Fullname:      model.Fullname,
		Username:      model.Username,
		Score:         model.Score,
		EmailVerified: model.EmailVerified,
	}
}

func ToExtendedUserDTO(model *ent.User) *dto.ExtendedUser {
	if model == nil {
		return nil
	}

	return &dto.ExtendedUser{
		User:         *ToUserDTO(model),
		PasswordHash: model.PasswordHash,

		Roles:       ToRoleDTOs(model.Edges.Roles),
		Permissions: ToPermissionDTOsByRoles(model.Edges.Roles),
	}
}

func ToUserDTOs(models ent.Users) dto.Users {
	if models == nil {
		return nil
	}
	dtms := make(dto.Users, len(models))
	for i := range models {
		dtms[i] = ToUserDTO(models[i])
	}
	return dtms
}
