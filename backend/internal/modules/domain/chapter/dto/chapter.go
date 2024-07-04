package dto

import (
	"fmt"
	"waterfall-backend/internal/models"
)

// Chapter глава
type Chapter struct {
	Uuid        string  // Uuid
	Number      int     // Номер
	Title       string  // Название
	PartUuid    string  // Часть, которой глава принадлежит
	ReleaseUuid *string // Uuid релиза
}

func (r *Chapter) Ref() models.ObjectRef {
	return models.ObjectRef{
		Type: models.ObjectTypeChapters,
		Ref:  r.Uuid,
	}
}

func (r *Chapter) String() string {
	return fmt.Sprintf("Глава %d: %s", r.Number, r.Title)
}

type Chapters []*Chapter

// ChapterCreate модель создания главы
type ChapterCreate struct {
	Number   int    // Номер
	Title    string // Название
	PartUuid string // Часть, которой глава принадлежит
}

// ChapterUpdate модель обновления главы
type ChapterUpdate struct {
	Number   int    // Номер
	Title    string // Название
	PartUuid string // Часть, которой глава принадлежит
}
