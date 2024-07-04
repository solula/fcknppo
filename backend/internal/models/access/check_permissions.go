package access

import (
	"context"
	"fmt"
	"slices"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/models/session"
)

func CheckPermissionsFromCtx(ctx context.Context, permission permissions.Type) error {
	ss, ok := session.GetFromCtx(ctx)
	if !ok {
		return err_const.ErrMissingSession
	}

	return CheckPermissions(ss.Permissions, permission)
}

func CheckPermissions(availablePermissions []permissions.Type, permission permissions.Type) error {
	// Если запрашиваемое разрешение не содержится в списке доступных из сессии -> запрещаем доступ
	if !slices.Contains(availablePermissions, permission) {
		return fmt.Errorf("%w: отсутствует право доступа %q", err_const.ErrAccessDenied, permission)
	}

	return nil
}
