package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
)

// FilesMakeNotTemp is the resolver for the FilesMakeNotTemp field.
func (r *mutationResolver) FilesMakeNotTemp(ctx context.Context, ids []string) ([]string, error) {
	err := r.services.FileStorage.BulkMakeNotTemp(ctx, ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

// FilesDelete is the resolver for the FilesDelete field.
func (r *mutationResolver) FilesDelete(ctx context.Context, ids []string) ([]string, error) {
	err := r.services.FileStorage.BulkDelete(ctx, ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

// FilesReorder is the resolver for the FilesReorder field.
func (r *mutationResolver) FilesReorder(ctx context.Context, ids []string) ([]string, error) {
	err := r.services.FileStorage.Reorder(ctx, ids)
	if err != nil {
		return nil, err
	}

	return ids, nil
}