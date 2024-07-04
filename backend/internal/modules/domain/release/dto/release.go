package dto

import (
	"time"
	"waterfall-backend/internal/models"
)

type Releasable interface {
	Ref() models.ObjectRef
	String() string
}

type Releasables []Releasable

// Release релиз объекта (главы, части и т.д.)
type Release struct {
	Uuid        string       // Uuid
	ReleaseDate time.Time    // Дата релиза
	Description string       // Описание
	Releasables []Releasable // Список выпускаемых сущностей
}

type Releases []*Release

// ReleaseCreate модель создания релиза
type ReleaseCreate struct {
	ReleaseDate time.Time           // Дата релиза
	Description string              // Описание
	ObjectRefs  []*models.ObjectRef // Ссылки на выпускаемые объекты

	ChapterUuids []string // Uuid-ы выпускаемых глав
	PartUuids    []string // Uuid-ы выпускаемых частей
}

// ReleaseUpdate модель обновления релиза
type ReleaseUpdate struct {
	ReleaseDate time.Time // Дата релиза
	Description string    // Описание
}
