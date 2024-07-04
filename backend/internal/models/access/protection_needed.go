package access

import "context"

type protectionNeededKey struct{}

func ProtectionNeeded(ctx context.Context) bool {
	_, ok := ctx.Value(protectionNeededKey{}).(struct{})
	return ok
}

func SetProtectionNeeded(ctx context.Context) context.Context {
	return context.WithValue(ctx, protectionNeededKey{}, struct{}{})
}
