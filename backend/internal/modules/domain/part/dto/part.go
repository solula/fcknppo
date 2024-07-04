package dto

import (
	"fmt"
	"waterfall-backend/internal/models"
)

// Part часть книги
type Part struct {
	Uuid        string  // Uuid
	Number      int     // Номер
	Title       string  // Название
	Annotation  *string // Аннотация
	ReleaseUuid *string // Uuid релиза
}

func (r *Part) Ref() models.ObjectRef {
	return models.ObjectRef{
		Type: models.ObjectTypeParts,
		Ref:  r.Uuid,
	}
}

func (r *Part) String() string {
	return fmt.Sprintf("Часть %d: %s", r.Number, r.Title)
}

type Parts []*Part

// PartCreate модель создания части
type PartCreate struct {
	Number     int     // Номер
	Title      string  // Название
	Annotation *string // Аннотация
}

// PartUpdate модель обновления
type PartUpdate struct {
	Number     int     // Номер
	Title      string  // Название
	Annotation *string // Аннотация
}
