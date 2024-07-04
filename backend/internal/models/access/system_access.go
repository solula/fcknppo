package access

import (
	"context"
	"waterfall-backend/internal/constants/users"
	"waterfall-backend/internal/models/session"
)

type systemKey struct{}

func IsSystem(ctx context.Context) bool {
	_, ok := ctx.Value(systemKey{}).(struct{})
	return ok
}

func SetSystem(ctx context.Context) context.Context {
	ss := session.Session{
		UserUuid: users.SystemUuid,
		Username: users.SystemUsername,
	}
	ctx = session.SetToCtx(ctx, ss)
	return context.WithValue(ctx, systemKey{}, struct{}{})
}
