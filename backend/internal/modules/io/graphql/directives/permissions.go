package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"waterfall-backend/internal/models/access"
	"waterfall-backend/internal/models/permissions"
)

func Permission(ctx context.Context, _ interface{}, next graphql.Resolver, permission *permissions.Type) (interface{}, error) {
	err := access.CheckPermissionsFromCtx(ctx, *permission)
	if err != nil {
		return nil, err
	}

	// Разрешаем пройти дальше
	return next(ctx)
}
