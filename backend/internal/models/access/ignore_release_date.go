package access

import "context"

type ignoreReleaseDateKey struct{}

func IgnoreReleaseDate(ctx context.Context) bool {
	_, ok := ctx.Value(ignoreReleaseDateKey{}).(struct{})
	return ok
}

func SetIgnoreReleaseDate(ctx context.Context) context.Context {
	return context.WithValue(ctx, ignoreReleaseDateKey{}, struct{}{})
}
