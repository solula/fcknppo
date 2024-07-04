package asynq_server

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"time"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/features/logger"
	"waterfall-backend/internal/pkg/asynq/enqueuer"
	"waterfall-backend/internal/pkg/asynq/middleware"
	"waterfall-backend/internal/pkg/asynq/router"
)

func NewAsynqServer(redisClient *redis.Client) *asynq.Server {
	asynqServer := newAsynqServer(redisClient)

	return asynqServer
}

func NewAsynqRouter(lg *zap.Logger) *router.AsynqRouter {
	asynqRouter := newRouter(lg)

	return asynqRouter
}

func InvokeAsynqServer(
	server *asynq.Server,
	router *router.AsynqRouter,
	lifecycle fx.Lifecycle,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Run(router)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.Shutdown()
			return nil
		},
	})
}

func newAsynqServer(redisClient *redis.Client) *asynq.Server {
	srv := asynq.NewServer(
		existingRedisClientOpt{client: redisClient},
		asynq.Config{
			// Максимальное кол-во воркеров
			Concurrency: 10,
			// Очереди с их приортетами
			Queues: map[string]int{
				string(enqueuer.PriorityCritical): 6,
				string(enqueuer.PriorityDefault):  3,
				string(enqueuer.PriorityLow):      1,
			},
			// Задержка между повторной обработкой задач
			RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
				return 1 * time.Minute
			},
		},
	)

	return srv
}

func newRouter(lg *zap.Logger) *router.AsynqRouter {
	var loggerToContext = func(ctx context.Context, lg *zap.Logger) (context.Context, error) {
		var extraFields []zap.Field

		taskID, ok := asynq.GetTaskID(ctx)
		if ok {
			extraFields = append(extraFields, zap.String("task_id", taskID))
		}

		priority, ok := asynq.GetQueueName(ctx)
		if ok {
			extraFields = append(extraFields, zap.String("priority", priority))
		}

		ctx = logger.SetToCtx(ctx, lg.With(extraFields...))
		return ctx, nil
	}

	asynqRouter := router.NewAsynqRouter()

	asynqRouter.Use(middleware.ExtractHeaders)
	asynqRouter.Use(middleware.ContextLogger(lg, loggerToContext))
	asynqRouter.Use(middleware.TaskLogger(lg))
	asynqRouter.Use(middleware.LoadSession(session.SetRawToCtx))
	asynqRouter.Use(middleware.PanicRecover(lg))

	return asynqRouter
}
