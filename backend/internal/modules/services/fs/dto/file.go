package dto

import (
	"fmt"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/models/files"
)

// File файл
type File struct {
	Uuid           string            // Uuid
	Filename       string            // Имя файла
	MIMEType       string            // MIME тип файла
	Description    string            // Описание
	CreatorUuid    *string           // Uuid создателя файла
	ObjectRef      *models.ObjectRef // Ссылка на сущность, которой файл принадлежит
	Type           files.Type        // Тип файла (картинка, аватар, т.д.)
	Temp           bool              // Признак временного файла
	SequenceNumber *uint             // Порядковый номер файла среди всех файлов этого ObjectRef с таким типом.
}

func (r *File) String() string {
	return fmt.Sprintf("Файл \"%s\"", r.Filename)
}

type Files []*File

// S3File представление файла для сохранения в S3
type S3File struct {
	BucketName string            // Имя бакета
	Filename   string            // Имя файла
	Metadata   map[string]string // Метаданные
}

// NewFile представление нового файла для создания
type NewFile struct {
	Filename    string            // Имя файла
	MIMEType    string            // MIME тип файла
	Description string            // Описание
	ObjectRef   *models.ObjectRef // Ссылка на сущность, которой файл принадлежит
	Type        files.Type        // Тип файла (картинка, аватар, т.д.)

	CreatorUuid *string // Uuid создателя файла (опционально)
}

// FileCreate модель создания файла в БД
type FileCreate struct {
	Filename       string            // Имя файла
	MIMEType       string            // MIME тип файла
	Description    string            // Описание
	CreatorUuid    *string           // Uuid создателя файла
	ObjectRef      *models.ObjectRef // Ссылка на сущность, которой файл принадлежит
	Type           files.Type        // Тип файла (картинка, аватар, т.д.)
	Temp           bool              // Признак временного файла
	SequenceNumber *uint             // Порядковый номер файла среди всех файлов этого ObjectRef с таким типом.
}
