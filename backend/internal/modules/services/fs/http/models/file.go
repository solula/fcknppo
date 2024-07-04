package models

import (
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/models/files"
)

type FileMetadata struct {
	Description string            // Описание
	ObjectRef   *models.ObjectRef // Ссылка на сущность, которой файл принадлежит
	Type        files.Type        // Тип файла (картинка, аватар, т.д.)
}
