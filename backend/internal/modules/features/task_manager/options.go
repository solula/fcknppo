package task_manager

import (
	"github.com/hibiken/asynq"
	"time"
	"waterfall-backend/internal/pkg/asynq/enqueuer"
)

type ManageOption = asynq.Option

func WithPriority(priority enqueuer.Priority) ManageOption {
	return asynq.Queue(string(priority))
}

func WithProcessAt(t time.Time) ManageOption {
	return asynq.ProcessAt(t)
}

func WithProcessIn(d time.Duration) ManageOption {
	return asynq.ProcessIn(d)
}
