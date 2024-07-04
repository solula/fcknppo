package task_manager

import (
	"context"
	"time"
	"waterfall-backend/internal/constants/typenames"
	"waterfall-backend/internal/modules/features/task_manager"
)

type FileStorageTaskManager struct {
	client *task_manager.TaskManager
}

func NewFileStorageTaskManager(client *task_manager.TaskManager) *FileStorageTaskManager {
	return &FileStorageTaskManager{
		client: client,
	}
}

func (t *FileStorageTaskManager) DeleteTempFileIn(ctx context.Context, fileUuid string, processIn time.Duration) (string, error) {
	return t.client.Manage(ctx, typenames.FileDeleteTemp, fileUuid, task_manager.WithProcessIn(processIn))
}
