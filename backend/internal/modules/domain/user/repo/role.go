package repo

import (
	"context"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/modules/domain/user/dto"
	"waterfall-backend/internal/modules/stores/db/converters"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/role"
)

type RoleRepo struct {
	client *ent.Client
}

func NewRoleRepo(client *ent.Client) *RoleRepo {
	return &RoleRepo{
		client: client,
	}
}

func (r *RoleRepo) GetWithPermissionsById(ctx context.Context, id roles.Type) (*dto.Role, dto.Permissions, error) {
	rl, err := r.client.Role.Query().
		Where(role.ID(id)).
		WithPermissions().
		Only(ctx)
	if err != nil {
		return nil, nil, wrap(err)
	}

	return converters.ToRoleDTO(rl), converters.ToPermissionDTOs(rl.Edges.Permissions), nil
}
