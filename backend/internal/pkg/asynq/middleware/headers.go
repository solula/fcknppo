package middleware

import (
	"context"
	"fmt"
	"github.com/hibiken/asynq"
	"waterfall-backend/internal/pkg/asynq/enqueuer"
	"waterfall-backend/internal/pkg/asynq/message"
)

// InitHeaders мидлвари при отправке для инициализации хэдеров
func InitHeaders(next enqueuer.HandlerFunc) enqueuer.HandlerFunc {
	return func(ctx context.Context, typename string, req interface{}, opts ...asynq.Option) (string, error) {
		ctx = message.OutgoingHeadersToContext(ctx, make(map[string]interface{}))

		return next(ctx, typename, req, opts...)
	}
}

// ExtractHeaders мидлвари при получении для извлечения хэдеров из Payload
func ExtractHeaders(next asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
		headers, err := message.ExtractHeaders(task.Payload())
		if err != nil {
			return fmt.Errorf("не удалось распаковать хэдеры: %w", err)
		}

		ctx = message.IncomingHeadersToContext(ctx, headers)

		return next.ProcessTask(ctx, task)
	})
}
