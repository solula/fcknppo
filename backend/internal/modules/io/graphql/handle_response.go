package graphql

import (
	"context"
	"errors"
	"github.com/99designs/gqlgen/graphql"
	"go.uber.org/zap"
	"net/http"
	"waterfall-backend/internal/modules/features/logger"
	"waterfall-backend/internal/pkg/http/error_handler/http_errors"
)

func handleResponse(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	// Получаем сформированный ответ
	resp := next(ctx)
	resp.Extensions = make(map[string]interface{})

	commonStatusCode := http.StatusOK
	// Карта полученных кодов
	statusCodes := make(map[int]struct{})

	var errorMessages []string
	for _, err := range resp.Errors {
		// Фактическая ошибка
		actualError := errors.Unwrap(err)
		err.Extensions = make(map[string]interface{})

		// Если фактической ошибки нет -> ошибка в самом запросе (BadRequest)
		if actualError == nil {
			statusCodes[http.StatusBadRequest] = struct{}{}
			continue
		}

		apiErr := http_errors.NewErrorResponse(actualError)
		err.Extensions["error_response"] = apiErr
		err.Message = apiErr.Message
		statusCodes[apiErr.Code] = struct{}{}

		errorMessages = append(errorMessages, apiErr.Details)
	}

	// Логируем ошибки, если они есть
	if len(resp.Errors) != 0 {
		logGraphqlErrors(ctx, errorMessages)
	}

	// Если присутствуют ошибки -> по умолчанию ставим код InternalServerError
	if len(resp.Errors) != 0 {
		commonStatusCode = http.StatusInternalServerError
	}

	// Обрабатываем коды в порядке их приоритета
	if _, ok := statusCodes[http.StatusNotFound]; ok {
		commonStatusCode = http.StatusNotFound
	}
	if _, ok := statusCodes[http.StatusForbidden]; ok {
		commonStatusCode = http.StatusForbidden
	}
	if _, ok := statusCodes[http.StatusBadRequest]; ok {
		commonStatusCode = http.StatusBadRequest
	}
	if _, ok := statusCodes[http.StatusInternalServerError]; ok {
		commonStatusCode = http.StatusInternalServerError
	}

	// Устанавливаем код статуса в Extensions
	resp.Extensions["status_code"] = commonStatusCode

	return resp
}

func logGraphqlErrors(ctx context.Context, errs []string) {
	opCtx := graphql.GetOperationContext(ctx)
	if opCtx == nil {
		return
	}

	logger.GetFromCtx(ctx).Error("Ошибка при обработке GraphQL запроса",
		zap.Strings("error", errs),
		zap.String("operation", opCtx.OperationName),
		zap.Any("variables", opCtx.Variables),
	)
}
