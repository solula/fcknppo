package dto

import "waterfall-backend/internal/models/permissions"

type Permission struct {
	Id          permissions.Type // Идентификатор
	Description string           // Описание
}

type Permissions []*Permission
