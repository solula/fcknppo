package error_handler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"waterfall-backend/internal/modules/features/logger"
	"waterfall-backend/internal/pkg/http/error_handler/http_errors"
)

func ErrorHandler(err error, c echo.Context) {
	var (
		apiError  *http_errors.ErrorResponse
		echoError *echo.HTTPError // Ошибка обработки роутинга и других ошибок
	)

	switch {
	case errors.As(err, &echoError):
		switch {
		case echoError.Code == http.StatusNotFound:
			apiError = &http_errors.ErrorResponse{
				Code:    http.StatusNotFound,
				Message: "Роут не найден",
				Details: "Такого роута не существует",
			}
		default:
			apiError = &http_errors.ErrorResponse{
				Code:    echoError.Code,
				Message: "Внутренняя ошибка сервера",
				Details: fmt.Sprintf("Подробности: %v", echoError.Message),
			}
		}
	case errors.As(err, &apiError):
		if apiError.Code == 0 {
			apiError.Code = http.StatusInternalServerError
		}
	default:
		apiError = http_errors.NewErrorResponse(err)
	}

	jsonErr := c.JSON(apiError.Code, apiError)
	if jsonErr != nil {
		logger.GetFromCtx(c.Request().Context()).Error("Ошибка при обработке ошибки", zap.Error(err))
	}
}
