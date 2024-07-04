package http_errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"net/http"
	"waterfall-backend/internal/constants/err_const"
)

// ErrorResponse ответ описание с ошибками
type ErrorResponse struct {
	Message string  // Текстовое сообщение для вывода клиенту
	Details string  // Сообщение техническое
	Code    int     // Код ошибки (добавлено для возможности регулирование из кода)
	Errors  []error `json:",omitempty"` // Набор ошибок если не одна
}

func (r ErrorResponse) Error() string {
	if len(r.Errors) > 1 {
		return multierr.Combine(r.Errors...).Error()
	}
	return r.Message
}

func (r ErrorResponse) MarshalJSON() ([]byte, error) {

	var errJson struct {
		Message string   `json:"message"`
		Details string   `json:"details"`
		Errors  []string `json:",omitempty"` // Не выводим поле если пусто
	}

	errJson.Message = r.Message
	errJson.Details = r.Details

	for _, err := range r.Errors {
		errJson.Errors = append(errJson.Errors, err.Error())
	}

	return json.Marshal(&errJson)
}

func NewErrorResponse(err error) *ErrorResponse {
	apiError := &ErrorResponse{}
	// Обработка ошибок, влияющих на статус
	apiError.Code, apiError.Message, apiError.Details = statusCodeAndErrorMessage(err)

	// Обработка multierr
	mErrs := multierr.Errors(err)
	if len(mErrs) > 1 {
		apiError.Errors = mErrs
		apiError.Details = "ответ содержит несколько ошибок"
	}

	// Обработка ошибок, полученных из паник
	var panicErr *err_const.PanicError
	if errors.As(err, &panicErr) {
		apiError.Message = panicErr.Message
		apiError.Details = panicErr.Detail
	}

	apiError = decorateErrorResponse(apiError)
	return apiError
}

func decorateErrorResponse(err *ErrorResponse) *ErrorResponse {
	prefix := "Ошибка при обработке запроса"
	switch err.Code {
	case http.StatusBadRequest:
		prefix = "Некорректный запрос"
	case http.StatusNotFound:
		prefix = "Ошибка 404"
	case http.StatusInternalServerError:
		prefix = "Внутренняя ошибка"

	}

	return &ErrorResponse{
		Message: fmt.Sprintf("%s: %s", prefix, err.Message),
		Details: fmt.Sprintf("Подробности: %s", err.Details),
		Code:    err.Code,
	}
}
