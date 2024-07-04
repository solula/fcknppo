package session

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/models/roles"
)

// Session сессия пользователя
type Session struct {
	SID          string             // Уникальный идентификатор сессии
	UserUuid     string             // Uuid пользователя
	Email        *string            // Почта пользователя
	Username     string             // Имя пользователя
	Roles        []roles.Type       // Список ролей пользователя
	Permissions  []permissions.Type // Список разрешений пользователя
	ReleaseDelay time.Duration      // Задержка доступности релизов
}

type sessionCtx struct{}

func GetFromCtx(ctx context.Context) (session Session, ok bool) {
	session, ok = ctx.Value(sessionCtx{}).(Session)
	return session, ok
}

func SetToCtx(ctx context.Context, session Session) context.Context {
	return context.WithValue(ctx, sessionCtx{}, session)
}

// GetRawFromCtx Функция для получения сессии из контекста (передается в мидлваре отправки)
func GetRawFromCtx(ctx context.Context) (json.RawMessage, error) {
	ss, ok := GetFromCtx(ctx)
	if !ok {
		return nil, fmt.Errorf("сессия не найдена")
	}

	return json.Marshal(ss)
}

// SetRawToCtx Функция для загрузки сессии в контекст (передается в мидлваре получения)
func SetRawToCtx(ctx context.Context, sessionBytes json.RawMessage) (context.Context, error) {
	var ss Session
	if err := json.Unmarshal(sessionBytes, &ss); err != nil {
		return nil, err
	}

	return SetToCtx(ctx, ss), nil
}
