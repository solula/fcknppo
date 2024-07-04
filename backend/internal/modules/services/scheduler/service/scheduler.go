package service

import (
	"context"
	"time"
	"waterfall-backend/internal/modules/domain/release/dto"
)

type ISchedulerTaskManager interface {
	CreateReleaseAt(ctx context.Context, release *dto.ReleaseCreate, scheduleAt time.Time) (string, error)
}

type SchedulerService struct {
	taskManager ISchedulerTaskManager
}

func NewSchedulerService(taskManager ISchedulerTaskManager) *SchedulerService {
	return &SchedulerService{
		taskManager: taskManager,
	}
}

// CreateReleaseAt запланировать создание релиза на дату scheduleAt
func (r *SchedulerService) CreateReleaseAt(ctx context.Context, release *dto.ReleaseCreate, scheduleAt time.Time) (string, error) {
	return r.taskManager.CreateReleaseAt(ctx, release, scheduleAt)
}
