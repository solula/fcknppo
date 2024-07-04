package utils

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/privacy"
)

// DefaultErrorWrapper базовый (универсальный) обработчик ошибок
func DefaultErrorWrapper(err error) error {
	if err == nil {
		return nil
	}

	// Если запись не найдена
	if ent.IsNotFound(err) {
		return err_const.ErrNotFound
	}

	// Если доступ запрещен
	if errors.Is(err, privacy.Deny) {
		return err_const.ErrAccessDenied
	}

	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		switch pqErr.Code {
		// Код 23505 - нарушение уникальности
		case "23505":
			return err_const.ErrUniqueConstraint

		// Код 22P02 - некорректное текстовое представление
		case "22P02":
			return err_const.ErrUuidValidate
		}
	}

	var entValidationErr *ent.ValidationError
	if errors.As(err, &entValidationErr) {
		fmtError := entValidationErr.Unwrap()
		if fmtError != nil {
			innerErr := errors.Unwrap(fmtError)
			if innerErr != nil {
				return fmt.Errorf("ошибка валидации %s: %w", entValidationErr.Name, innerErr)
			}
		}
	}

	return err
}
