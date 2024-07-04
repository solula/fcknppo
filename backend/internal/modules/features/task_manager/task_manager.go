package task_manager

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"io"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/pkg/asynq/enqueuer"
	"waterfall-backend/internal/pkg/asynq/middleware"
)

type TaskManager struct {
	enqueuer     enqueuer.Enqueuer
	inspector    *asynq.Inspector
	clientCloser io.Closer
}

func (r *TaskManager) Close() error {
	return r.clientCloser.Close()
}

func (r *TaskManager) Manage(ctx context.Context, typename string, req interface{}, opts ...ManageOption) (string, error) {
	return r.enqueuer.Enqueue(ctx, typename, req, opts...)
}

func (r *TaskManager) DeleteTask(priority enqueuer.Priority, taskID string) error {
	return r.inspector.DeleteTask(string(priority), taskID)
}

func NewTaskManager(redisClient *redis.Client) *TaskManager {
	client := asynq.NewClient(existingRedisClientOpt{client: redisClient})
	inspector := asynq.NewInspector(existingRedisClientOpt{client: redisClient})

	enq := enqueuer.NewEnqueuer(client)

	enq.Use(middleware.InitHeaders)
	enq.Use(middleware.StoreSession(session.GetRawFromCtx))

	return &TaskManager{
		enqueuer:     enq,
		inspector:    inspector,
		clientCloser: client,
	}
}

func InvokeTaskManager(
	taskManager *TaskManager,
	lifecycle fx.Lifecycle,
) error {
	lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			err1 := taskManager.Close()
			err2 := taskManager.inspector.Close()

			if err1 != nil {
				return err1
			}
			if err2 != nil {
				return err2
			}

			return nil
		},
	})

	return nil
}
