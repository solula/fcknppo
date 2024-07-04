package http_errors

import (
	"errors"
	"fmt"
	"net/http"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/pkg/file_type"
)

const (
	AuthenticationErrorDescription = "неправильный логин или пароль (возможно, вы использовали другой метод аутентификации?)"
)

func statusCodeAndErrorMessage(err error) (int, string, string) {
	switch {
	case errors.Is(err, err_const.ErrInvalidToken):
		return http.StatusUnauthorized, err_const.ErrInvalidToken.Error(), err.Error()
	case errors.Is(err, err_const.ErrAuthentication):
		// Не раскрываем подробности ошибки, если ошибка при аутентификации
		return http.StatusUnauthorized, err_const.ErrAuthentication.Error(), AuthenticationErrorDescription
	case errors.Is(err, err_const.ErrAccessDenied):
		return http.StatusForbidden, err_const.ErrAccessDenied.Error(), err.Error()
	case errors.Is(err, err_const.ErrNotFound):
		return http.StatusNotFound, err_const.ErrNotFound.Error(), err.Error()
	case errors.Is(err, err_const.ErrInvalidEmail):
		return http.StatusBadRequest, err_const.ErrInvalidEmail.Error(), err.Error()
	case errors.Is(err, err_const.ErrEmailNotVerified):
		return http.StatusUnauthorized, err_const.ErrEmailNotVerified.Error(), err.Error()
	case errors.Is(err, err_const.ErrPasswordTooEasy):
		return http.StatusBadRequest, err_const.ErrPasswordTooEasy.Error(), err.Error()
	case errors.Is(err, err_const.ErrValidate):
		return http.StatusBadRequest, err_const.ErrValidate.Error(), err.Error()
	case errors.Is(err, err_const.ErrUuidValidate):
		return http.StatusBadRequest, err_const.ErrUuidValidate.Error(), err.Error()
	case errors.Is(err, err_const.ErrIdValidate):
		return http.StatusBadRequest, err_const.ErrIdValidate.Error(), err.Error()
	case errors.Is(err, err_const.ErrUuidMissing):
		return http.StatusBadRequest, err_const.ErrUuidMissing.Error(), err.Error()
	case errors.Is(err, err_const.ErrIdMissing):
		return http.StatusBadRequest, err_const.ErrIdMissing.Error(), err.Error()
	case errors.Is(err, err_const.ErrMissingUser):
		return http.StatusBadRequest, err_const.ErrMissingUser.Error(), err.Error()
	case errors.Is(err, err_const.ErrMissingToken):
		return http.StatusBadRequest, err_const.ErrMissingToken.Error(), err.Error()
	case errors.Is(err, err_const.ErrUniqueConstraint):
		return http.StatusBadRequest, err_const.ErrUniqueConstraint.Error(), err.Error()
	case errors.Is(err, file_type.ErrWrongFileType):
		return http.StatusBadRequest, file_type.ErrWrongFileType.Error(), err.Error()
	case errors.Is(err, err_const.ErrInvalidNumber):
		return http.StatusBadRequest, err_const.ErrInvalidNumber.Error(), err.Error()
	case errors.Is(err, err_const.ErrInvalidIndex):
		return http.StatusBadRequest, err_const.ErrInvalidIndex.Error(), err.Error()
	case errors.Is(err, err_const.ErrInvalidObjectType):
		return http.StatusBadRequest, err_const.ErrInvalidObjectType.Error(), err.Error()
	default:
		return http.StatusInternalServerError, fmt.Errorf("ошибка обработки запроса: %w", err).Error(), err.Error()
	}
}
