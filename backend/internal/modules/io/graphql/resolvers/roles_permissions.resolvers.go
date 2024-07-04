package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"waterfall-backend/internal/modules/domain/user/dto"
	"waterfall-backend/internal/modules/io/graphql/generated"
)

// ReleaseDelay is the resolver for the ReleaseDelay field.
func (r *roleResolver) ReleaseDelay(ctx context.Context, obj *dto.Role) (string, error) {
	return obj.ReleaseDelay.String(), nil
}

// Role returns generated.RoleResolver implementation.
func (r *Resolver) Role() generated.RoleResolver { return &roleResolver{r} }

type roleResolver struct{ *Resolver }