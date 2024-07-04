package middleware

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
	"time"
	"waterfall-backend/internal/pkg/asynq/message"
)

type LoggerToContext func(context.Context, *zap.Logger) (context.Context, error)

// ContextLogger возвращает мидлвари при получении для добавления логгера в контекст
func ContextLogger(lg *zap.Logger, loggerToContext LoggerToContext) asynq.MiddlewareFunc {
	return func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
			ctx, err := loggerToContext(ctx, lg)
			if err != nil {
				return err
			}

			return next.ProcessTask(ctx, task)
		})
	}
}

// TaskLogger возвращает мидлвари при получении для логирования задач
func TaskLogger(lg *zap.Logger) asynq.MiddlewareFunc {
	return func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
			var start, stop time.Time

			start = time.Now()
			err := next.ProcessTask(ctx, task)
			stop = time.Now()

			fields := []zap.Field{
				zap.String("latency", stop.Sub(start).String()),
				zap.String("typename", task.Type()),
				zap.ByteString("payload", task.Payload()),
			}

			// Опциональные поля
			headers, headersFound := message.IncomingHeadersFromContext(ctx)
			taskID, taskIDFound := asynq.GetTaskID(ctx)
			priority, priorityFound := asynq.GetQueueName(ctx)
			if headersFound {
				fields = append(fields, zap.Any("headers", headers))
			}
			if taskIDFound {
				fields = append(fields, zap.String("taskId", taskID))
			}
			if priorityFound {
				fields = append(fields, zap.String("priority", priority))
			}
			if err != nil {
				fields = append(fields, zap.Error(err))
			}

			if err != nil {
				msg := fmt.Sprintf("Внутренняя ошибка сервера: %s", err.Error())
				lg.Error(msg, fields...)
			} else {
				lg.Info("Запрос выполнен успешно", fields...)
			}

			return err
		})
	}
}
