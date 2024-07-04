package files

import (
	"slices"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/pkg/file_type"
)

type Type string // Тип файла

const (
	Avatar Type = "avatar" // Аватарка пользователя
	Image  Type = "image"  // Изображение (например, к главе)
)

func (t Type) Values() []string {
	return []string{
		string(Avatar),
		string(Image),
	}
}

func (t Type) Validate() error {
	if !slices.Contains(t.Values(), string(t)) {
		return err_const.ErrInvalidRole
	}
	return nil
}

func (t Type) AvailableFileTypes() []string {
	var availableTypes []string

	switch t {
	case Image:
		availableTypes = []string{file_type.TypePNG, file_type.TypeJPEG}
	case Avatar:
		availableTypes = []string{file_type.TypePNG, file_type.TypeJPEG}

	default:
		availableTypes = []string{}
	}

	return availableTypes
}
