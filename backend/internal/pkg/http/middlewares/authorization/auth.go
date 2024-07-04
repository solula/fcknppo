package authorization

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/domain/auth/dto"
	"waterfall-backend/internal/modules/features/logger"
)

const AccessTokenKey = "access_token"

// Auther парсит токен
type Auther interface {
	AuthAccessToken(_ context.Context, accessToken string) (*dto.AccessTokenPayload, error)
	GenerateGuestSession(ctx context.Context) (*session.Session, error)
	GenerateUserSession(ctx context.Context, userUuid string) (*session.Session, error)
}

// WithAuth парсит токен и кладет контекст сессию пользователя
func WithAuth(auther Auther) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var ss *session.Session
			ctx := c.Request().Context()

			token, ok := c.Get(AccessTokenKey).(string)
			if !ok {
				// Если токен не найден - создаем сессию гостя
				guestSession, err := auther.GenerateGuestSession(ctx)
				if err != nil {
					return fmt.Errorf("не удалось сгенерировать гостевую сессию: %w", err)
				}

				ss = guestSession
			} else {
				// Если токен найден - парсим и получаем сессию
				tokenPayload, err := auther.AuthAccessToken(ctx, token)
				if err != nil {
					// Если в токене ошибка - возвращаем ошибку
					return err
				}

				userSession, err := auther.GenerateUserSession(ctx, tokenPayload.UserUuid)
				if err != nil {
					return fmt.Errorf("не удалось сгенерировать сессию пользователя: %w", err)
				}

				ss = userSession
			}

			// Добавляем в логгер информацию о пользователе
			lg := logger.GetFromCtx(ctx)
			ctx = logger.SetToCtx(ctx, lg.With(logger.UserFields(*ss)...))

			// Далее все действия идут уже от имени пользователя
			ctx = session.SetToCtx(ctx, *ss)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
