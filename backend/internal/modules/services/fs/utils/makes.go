package utils

import (
	"fmt"
	"path/filepath"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/modules/services/fs/dto"
)

func MakeBucketName(objectRef *models.ObjectRef) string {
	return fmt.Sprintf("%s-%s", objectRef.Type, objectRef.Ref)
}

func MakeFilename(file *dto.File) string {
	// Для уникальности файлы называем как uuid + расширение
	return file.Uuid + filepath.Ext(file.Filename)
}

func MakeMetadata(file *dto.NewFile) map[string]string {
	metadata := make(map[string]string)

	metadata["filename"] = file.Filename
	metadata["mimetype"] = file.MIMEType
	metadata["description"] = file.Description

	return metadata
}
