package task_manager

import (
	"context"
	"time"
	"waterfall-backend/internal/constants/typenames"
	"waterfall-backend/internal/modules/domain/release/dto"
	"waterfall-backend/internal/modules/features/task_manager"
)

type SchedulerTaskManager struct {
	client *task_manager.TaskManager
}

func NewSchedulerTaskManager(client *task_manager.TaskManager) *SchedulerTaskManager {
	return &SchedulerTaskManager{
		client: client,
	}
}

func (t *SchedulerTaskManager) CreateReleaseAt(ctx context.Context, release *dto.ReleaseCreate, scheduleAt time.Time) (string, error) {
	return t.client.Manage(ctx, typenames.ReleaseCreate, release, task_manager.WithProcessAt(scheduleAt))
}
