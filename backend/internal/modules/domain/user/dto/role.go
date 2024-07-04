package dto

import (
	"time"
	"waterfall-backend/internal/models/roles"
)

type Role struct {
	Id           roles.Type    // Идентификатор
	Description  string        // Описание
	ReleaseDelay time.Duration // Задержка релиза
}

type Roles []*Role
