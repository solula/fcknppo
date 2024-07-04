package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"waterfall-backend/internal/pkg/asynq/enqueuer"
	"waterfall-backend/internal/pkg/asynq/message"
)

type (
	// SessionFromContext тип функции для получения сессии из контекста
	SessionFromContext func(context.Context) (json.RawMessage, error)
	// SessionToContext тип функции для загрузки сессии в контекст
	SessionToContext func(context.Context, json.RawMessage) (context.Context, error)
)

const sessionHeaderKey = "session"

// StoreSession возвращает мидлвари при отправке для сохранения сессии в хэдерах
func StoreSession(sessionFromContext SessionFromContext) enqueuer.MiddlewareFunc {
	return func(next enqueuer.HandlerFunc) enqueuer.HandlerFunc {
		return func(ctx context.Context, typename string, req interface{}, opts ...asynq.Option) (string, error) {
			headers, ok := message.OutgoingHeadersFromContext(ctx)
			if !ok {
				return "", fmt.Errorf("хэдеры не найдены")
			}

			sessionRawMessage, err := sessionFromContext(ctx)
			if err != nil {
				return "", err
			}

			headers[sessionHeaderKey] = sessionRawMessage

			return next(ctx, typename, req, opts...)
		}
	}
}

// LoadSession возвращает мидлвари при получении для загрузки сессии из хэдеров
func LoadSession(sessionToContext SessionToContext) asynq.MiddlewareFunc {
	return func(next asynq.Handler) asynq.Handler {
		return asynq.HandlerFunc(func(ctx context.Context, task *asynq.Task) error {
			headers, ok := message.IncomingHeadersFromContext(ctx)
			if !ok {
				return fmt.Errorf("хэдеры не найдены")
			}

			// Достаем сессию
			sessionInterface, ok := headers[sessionHeaderKey]
			if !ok {
				return fmt.Errorf("сессия не найдена в хэдере")
			}

			// До этого ее уже размаршалили в мапу
			sessionMap, ok := sessionInterface.(map[string]interface{})
			if !ok {
				return fmt.Errorf("сессия не является мапой")
			}

			// Обратно маршалим сессию в RawMessage
			sessionBytes, err := json.Marshal(sessionMap)
			if err != nil {
				return fmt.Errorf("ошибка маршалинга сессии: %w", err)
			}

			ctx, err = sessionToContext(ctx, sessionBytes)
			if err != nil {
				return err
			}

			return next.ProcessTask(ctx, task)
		})
	}
}
