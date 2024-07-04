package middlewares

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"strings"
	"time"
	"waterfall-backend/internal/models/session"
	logger2 "waterfall-backend/internal/modules/features/logger"
	"waterfall-backend/internal/pkg/http/constants"

	"go.uber.org/zap"
)

func ContextLogger(lg *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := logger2.SetToCtx(c.Request().Context(), lg)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

func RequestLogger(lg *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var start, stop time.Time

			reqBody, dumpErr := dumpRequestBody(c.Request())
			if dumpErr != nil {
				return fmt.Errorf("ошибка при логировании: %w", dumpErr)
			}

			start = time.Now()
			err := next(c)
			stop = time.Now()

			// Обрабатываем ошибку
			if err != nil {
				// Вынимаем из echo.HTTPError обернутую ошибку, если она там указана
				var echoError *echo.HTTPError
				if errors.As(err, &echoError) {
					wrappedErr := echoError.Unwrap()
					if wrappedErr != nil {
						err = wrappedErr
					}
				}

				// Вызываем установленный обработчик
				c.Error(err)
			}

			fields := []zap.Field{
				zap.String("ip", c.RealIP()),
				zap.String("latency", stop.Sub(start).String()),
				zap.String("path", c.Request().RequestURI),
				zap.String("method", c.Request().Method),
				// zap.Any("headers", c.Request().Header),
				zap.Int("status", c.Response().Status),
			}

			// Опциональные поля
			if needLogPayload(c.Request()) {
				fields = append(fields, zap.ByteString("body", reqBody))
			}

			ss, ok := session.GetFromCtx(c.Request().Context())
			if ok {
				fields = append(fields, logger2.UserFields(ss)...)
			}

			if err != nil {
				fields = append(fields, zap.Error(err))
			}

			s := c.Response().Status
			switch {
			case s >= 500:
				msg := fmt.Sprintf("Неизвестная внутренняя ошибка")
				if err != nil {
					msg = fmt.Sprintf("Внутренняя ошибка сервера: %s", err.Error())
				}
				lg.Error(msg, fields...)
			case s >= 400:
				msg := fmt.Sprintf("Неизвестная ошибка в запросе")
				if err != nil {
					msg = fmt.Sprintf("Ошибка в запросе: %s", err.Error())
				}
				lg.Warn(msg, fields...)
			default:
				lg.Info("Запрос выполнен успешно", fields...)
			}

			// Возвращаем nil, т.к. ошибку уже обработали
			return nil
		}
	}
}

func dumpRequestBody(req *http.Request) ([]byte, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать тело запроса: %w", err)
	}
	if len(body) == 0 {
		return nil, nil
	}

	err = req.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("не удалось закрыть тело запроса: %w", err)
	}
	req.Body = io.NopCloser(bytes.NewReader(body))

	return body, nil
}

func needLogPayload(req *http.Request) bool {
	// Пропускаем тело запроса, если тело не в формате JSON...
	skipPayload := req.Header.Get(constants.HeaderContentType) != constants.MIMEApplicationJSON
	// ...или если это запрос в группе /auth
	skipPayload = skipPayload || strings.Contains(req.RequestURI, "/auth")

	return !skipPayload
}
