package middleware

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"go.uber.org/zap"
)

// PanicRecover возвращает мидлвари при получении для обработки паник
func PanicRecover(logger *zap.Logger) asynq.MiddlewareFunc {
	return func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) (err error) {
			defer func() {
				if value := recover(); value != nil {
					logger.Error("Паника при обработке задачи", zap.Any("panic", value), zap.Stack("stacktrace"))
					err = fmt.Errorf("%v", value)
				}
			}()

			return next.ProcessTask(ctx, task)
		})
	}
}
