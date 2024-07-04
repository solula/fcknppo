package repo

import (
	"context"
	"fmt"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/modules/services/fs/dto"
	"waterfall-backend/internal/modules/stores/db/converters"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/file"
	"waterfall-backend/internal/modules/stores/db/utils"
)

type FileRepo struct {
	client *ent.Client
}

func NewFileRepo(client *ent.Client) *FileRepo {
	return &FileRepo{
		client: client,
	}
}

// GetByUuid получение записи
func (r *FileRepo) GetByUuid(ctx context.Context, uuid string) (*dto.File, error) {
	dtm, err := r.client.File.Get(ctx, uuid)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToFileDTO(dtm), nil
}

// ListNotTempByObjectRef получает список файлов (не временных), принадлежащих указанному объекту
func (r *FileRepo) ListNotTempByObjectRef(ctx context.Context, objectRef *models.ObjectRef) (dto.Files, error) {
	dtm, err := r.client.File.Query().Where(
		file.ObjectType(objectRef.Type),
		file.ObjectRef(objectRef.Ref),
		file.Temp(false),
	).All(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToFilesDTOs(dtm), nil
}

// Create создает файл
func (r *FileRepo) Create(ctx context.Context, file *dto.FileCreate) (*dto.File, error) {
	newFile, err := r.client.File.Create().
		SetFilename(file.Filename).
		SetMimeType(file.MIMEType).
		SetDescription(file.Description).
		SetNillableCreatorUUID(file.CreatorUuid).
		SetObjectType(file.ObjectRef.Type).
		SetObjectRef(file.ObjectRef.Ref).
		SetType(file.Type).
		SetTemp(file.Temp).
		SetNillableSequenceNumber(file.SequenceNumber).
		Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToFileDTO(newFile), nil
}

// Delete удаление записи
func (r *FileRepo) Delete(ctx context.Context, uuid string) error {
	err := r.client.File.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return wrap(err)
	}

	return nil
}

// DeleteByObjectRef удаление всех файлов, принадлежащих указанному объекту
func (r *FileRepo) DeleteByObjectRef(ctx context.Context, objectRef *models.ObjectRef) error {
	_, err := r.client.File.Delete().Where(
		file.ObjectTypeEQ(objectRef.Type),
		file.ObjectRefEQ(objectRef.Ref),
	).Exec(ctx)
	if err != nil {
		return wrap(err)
	}

	return nil
}

// BulkMakeNotTemp помечает указанные файлы как не временные;
// воздействует только на свои файлы
func (r *FileRepo) BulkMakeNotTemp(ctx context.Context, uuids []string) (uint, error) {
	cnt, err := r.client.File.Update().
		Where(file.IDIn(uuids...)).
		SetTemp(false).
		Save(ctx)
	if err != nil {
		return 0, wrap(err)
	}

	return uint(cnt), nil
}

// BulkDelete удаляет указанные файлы;
// воздействует только на свои файлы
func (r *FileRepo) BulkDelete(ctx context.Context, uuids []string) (uint, error) {
	cnt, err := r.client.File.Delete().
		Where(file.IDIn(uuids...)).
		Exec(ctx)
	if err != nil {
		return 0, wrap(err)
	}

	return uint(cnt), nil
}

// UpdateSequenceNumber назначает файлу порядковый номер;
// воздействует только на свои файлы
func (r *FileRepo) UpdateSequenceNumber(ctx context.Context, uuid string, seqNum uint) error {
	err := r.client.File.UpdateOneID(uuid).
		SetSequenceNumber(seqNum).
		Exec(ctx)
	if err != nil {
		return wrap(err)
	}

	return nil
}

// CheckFilenameExists проверка наличия файла
func (r *FileRepo) CheckFilenameExists(ctx context.Context, objectRef *models.ObjectRef, filename string) (bool, error) {
	_, err := r.client.File.Query().Where(
		file.ObjectTypeEQ(objectRef.Type),
		file.ObjectRefEQ(objectRef.Ref),
		file.Filename(filename),
	).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, wrap(err)
	}

	return true, nil
}

// CheckObjectGettable проверка возможности получения объекта
func (r *FileRepo) CheckObjectGettable(ctx context.Context, objectRef *models.ObjectRef) (bool, error) {
	var err error
	switch objectRef.Type {
	case models.ObjectTypeChapters:
		_, err = r.client.Chapter.Get(ctx, objectRef.Ref)
	case models.ObjectTypeParts:
		_, err = r.client.Part.Get(ctx, objectRef.Ref)
	case models.ObjectTypeUsers:
		_, err = r.client.User.Get(ctx, objectRef.Ref)

	default:
		return false, fmt.Errorf("неразрешенный тип объекта: %s", objectRef.Type)
	}

	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, wrap(err)
	}

	return true, nil
}

// CheckObjectUpdatable проверка возможности обновления объекта
func (r *FileRepo) CheckObjectUpdatable(ctx context.Context, objectRef *models.ObjectRef) (bool, error) {
	var err error
	switch objectRef.Type {
	case models.ObjectTypeChapters:
		err = r.client.Chapter.UpdateOneID(objectRef.Ref).Exec(ctx)
	case models.ObjectTypeParts:
		err = r.client.Part.UpdateOneID(objectRef.Ref).Exec(ctx)
	case models.ObjectTypeUsers:
		err = r.client.User.UpdateOneID(objectRef.Ref).Exec(ctx)

	default:
		return false, fmt.Errorf("неразрешенный тип объекта: %s", objectRef.Type)
	}

	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, wrap(err)
	}

	return true, nil
}

func wrap(err error) error {
	return utils.DefaultErrorWrapper(err)
}
