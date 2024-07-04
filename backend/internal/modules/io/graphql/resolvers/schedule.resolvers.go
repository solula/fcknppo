package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"time"
	"waterfall-backend/internal/modules/domain/release/dto"
)

// ScheduleReleaseCreate is the resolver for the ScheduleReleaseCreate field.
func (r *mutationResolver) ScheduleReleaseCreate(ctx context.Context, release dto.ReleaseCreate, scheduleAt time.Time) (string, error) {
	return r.services.Scheduler.CreateReleaseAt(ctx, &release, scheduleAt)
}
