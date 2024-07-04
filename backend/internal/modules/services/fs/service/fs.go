package service

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"io"
	"slices"
	"time"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models"
	"waterfall-backend/internal/models/access"
	"waterfall-backend/internal/models/files"
	"waterfall-backend/internal/models/permissions"
	"waterfall-backend/internal/models/session"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/features/logger"
	"waterfall-backend/internal/modules/services/fs/dto"
	fs_utils "waterfall-backend/internal/modules/services/fs/utils"
	"waterfall-backend/internal/pkg/file_type"
	"waterfall-backend/internal/pkg/transaction"
)

type IFileRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.File, error)
	ListNotTempByObjectRef(ctx context.Context, objectRef *models.ObjectRef) (dto.Files, error)
	Create(ctx context.Context, file *dto.FileCreate) (*dto.File, error)
	Delete(ctx context.Context, uuid string) error
	BulkDelete(ctx context.Context, uuids []string) (uint, error)
	BulkMakeNotTemp(ctx context.Context, uuids []string) (uint, error)
	UpdateSequenceNumber(ctx context.Context, uuid string, seqNum uint) error

	DeleteByObjectRef(ctx context.Context, objectRef *models.ObjectRef) error
	CheckObjectGettable(ctx context.Context, objectRef *models.ObjectRef) (bool, error)
	CheckObjectUpdatable(ctx context.Context, objectRef *models.ObjectRef) (bool, error)

	transaction.TxRepo
}

type IS3 interface {
	GetFile(ctx context.Context, bucketName string, fileName string) ([]byte, error)
	CreateBucket(ctx context.Context, bucketName string) error
	SaveFile(ctx context.Context, file *dto.S3File, content io.Reader) (int, error)
	CopyFile(ctx context.Context, fromBucketName, fromFilename string, toBucketName, toFilename string) error
	DeleteFile(ctx context.Context, bucketName string, fileName string) error
	DeleteBucket(ctx context.Context, bucketName string) error
}

type IFileStorageTaskManager interface {
	DeleteTempFileIn(ctx context.Context, fileUuid string, processIn time.Duration) (string, error)
}

type FileStorageService struct {
	repo        IFileRepo
	fileStorage IS3
	taskManager IFileStorageTaskManager
	filesTTL    time.Duration
}

func NewFileStorageService(repo IFileRepo, fileStorage IS3, taskManager IFileStorageTaskManager, cfg config.Config) (*FileStorageService, error) {
	return &FileStorageService{
		repo:        repo,
		fileStorage: fileStorage,
		taskManager: taskManager,
		filesTTL:    cfg.FilesTTL,
	}, nil
}

func (s *FileStorageService) GetFileByUuid(ctx context.Context, uuid string) (*dto.File, []byte, error) {
	file, err := s.getByUuid(ctx, uuid)
	if err != nil {
		return nil, nil, fmt.Errorf("не удалось получить файл: %w", err)
	}

	fileContent, err := s.fileStorage.GetFile(ctx, fs_utils.MakeBucketName(file.ObjectRef), fs_utils.MakeFilename(file))
	if err != nil {
		return nil, nil, fmt.Errorf("не удалось получить содержимое файла: %w", err)
	}

	return file, fileContent, nil
}

func (s *FileStorageService) ListNotTempByObjectRef(ctx context.Context, objectRef *models.ObjectRef) (dto.Files, error) {
	filesList, err := s.listNotTempByObjectRef(ctx, objectRef)
	if err != nil {
		return nil, err
	}

	return filesList, nil
}

func (s *FileStorageService) CreateBucket(ctx context.Context, objectRef *models.ObjectRef) error {
	err := s.fileStorage.CreateBucket(ctx, fs_utils.MakeBucketName(objectRef))
	if err != nil {
		return fmt.Errorf("не удалось создать бакет: %w", err)
	}

	return nil
}

// CreateTempFile создает временный файл
func (s *FileStorageService) CreateTempFile(ctx context.Context, newFile *dto.NewFile, content io.Reader) (*dto.File, error) {
	createdFile, err := s.createFile(ctx, newFile, true, content)
	if err != nil {
		return nil, err
	}

	// Планируем удаление временного файла по истечению его TTL
	_, err = s.taskManager.DeleteTempFileIn(ctx, createdFile.Uuid, s.filesTTL)
	if err != nil {
		logger.GetFromCtx(ctx).Warn("не удалось запланировать удаление временного файла", zap.String("uuid", createdFile.Uuid), zap.Error(err))
	}

	return createdFile, nil
}

// CreateFile создает постоянный файл
func (s *FileStorageService) CreateFile(ctx context.Context, newFile *dto.NewFile, content io.Reader) (*dto.File, error) {
	createdFile, err := s.createFile(ctx, newFile, false, content)
	if err != nil {
		return nil, err
	}

	return createdFile, nil
}

func (s *FileStorageService) DeleteTempFile(ctx context.Context, uuid string) error {
	file, err := s.getByUuid(ctx, uuid)
	if err != nil {
		return fmt.Errorf("не удалось получить файл: %w", err)
	}

	// Если файл уже не временный, то удалять его не нужно
	if !file.Temp {
		return nil
	}

	err = s.fileStorage.DeleteFile(ctx, fs_utils.MakeBucketName(file.ObjectRef), fs_utils.MakeFilename(file))
	if err != nil {
		return fmt.Errorf("не удалось удалить содержимое файла: %w", err)
	}

	err = s.repo.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("не удалось удалить файл: %w", err)
	}

	return nil
}

func (s *FileStorageService) CopyFile(ctx context.Context, uuid string, toObjectRef *models.ObjectRef, newCreatorUuid string) error {
	if !access.IsSystem(ctx) {
		return fmt.Errorf("%w: копировать файлы может только системный пользователь", err_const.ErrAccessDenied)
	}

	file, err := s.getByUuid(ctx, uuid)
	if err != nil {
		return fmt.Errorf("не удалось получить файл: %w", err)
	}

	if file.Temp {
		return fmt.Errorf("временный файлы копировать запрещено")
	}

	// Новый файл создается в бакете объекта
	fileCreate := &dto.FileCreate{
		Filename:       file.Filename,
		MIMEType:       file.MIMEType,
		Description:    file.Description,
		CreatorUuid:    &newCreatorUuid,
		ObjectRef:      toObjectRef,
		Type:           file.Type,
		Temp:           file.Temp,
		SequenceNumber: file.SequenceNumber,
	}

	newFile, err := s.repo.Create(ctx, fileCreate)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s: %w", file.Filename, err)
	}

	err = s.fileStorage.CopyFile(ctx, fs_utils.MakeBucketName(file.ObjectRef), fs_utils.MakeFilename(file), fs_utils.MakeBucketName(toObjectRef), fs_utils.MakeFilename(newFile))
	if err != nil {
		return fmt.Errorf("не удалось переместить файл: %w", err)
	}

	return nil
}

// BulkMakeNotTemp помечает указанные файлы как не временные;
// воздействует только на свои файлы
func (s *FileStorageService) BulkMakeNotTemp(ctx context.Context, uuids []string) error {
	cnt, err := s.repo.BulkMakeNotTemp(ctx, uuids)
	if err != nil {
		return fmt.Errorf("не удалось пометить файлы как не временные: %w", err)
	}

	if len(uuids) != int(cnt) {
		logger.GetFromCtx(ctx).Warn("Не удалось пометить все файлы как не временные", zap.Uint("cnt", cnt), zap.Strings("uuids", uuids))
	}

	return nil
}

// BulkDelete удаляет указанные файлы;
// воздействует только на свои файлы
func (s *FileStorageService) BulkDelete(ctx context.Context, uuids []string) error {
	cnt, err := s.repo.BulkDelete(ctx, uuids)
	if err != nil {
		return fmt.Errorf("не удалось удалить файлы: %w", err)
	}

	if len(uuids) != int(cnt) {
		logger.GetFromCtx(ctx).Warn("Не удалось удалить все указанные файлы", zap.Uint("cnt", cnt), zap.Strings("uuids", uuids))
	}

	return nil
}

// Reorder назначает файлам порядковые номера в соответствии с последовательностью ids;
// воздействует только на свои файлы
func (s *FileStorageService) Reorder(ctx context.Context, uuids []string) error {
	err := transaction.WithTx(ctx, s.repo, func(txRepo IFileRepo) error {
		for i, uuid := range uuids {
			err := txRepo.UpdateSequenceNumber(ctx, uuid, uint(i))
			if err != nil {
				return fmt.Errorf("не удалось назначить порядковый номер %d для файла %s: %w", i, uuid, err)
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("ошибка при назначении порядковых номеров: %w", err)
	}

	return nil
}

func (s *FileStorageService) DeleteBucket(ctx context.Context, objectRef *models.ObjectRef) error {
	err := s.repo.DeleteByObjectRef(ctx, objectRef)
	if err != nil {
		return fmt.Errorf("не удалось удалить файлы: %w", err)
	}

	err = s.fileStorage.DeleteBucket(ctx, fs_utils.MakeBucketName(objectRef))
	if err != nil {
		return fmt.Errorf("не удалось удалить бакет: %w", err)
	}

	return nil
}

// getByUuid получает файл по uuid
func (s *FileStorageService) getByUuid(ctx context.Context, uuid string) (*dto.File, error) {
	file, err := s.repo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить файл: %w", err)
	}

	// Публичные файлы доступны всем пользователям
	if file.ObjectRef.Type == models.ObjectTypePublic {
		return file, nil
	}

	// Пользователю доступны только его временные файлы
	if file.Temp {
		ss, ok := session.GetFromCtx(ctx)
		if !ok || ss.UserUuid == "" {
			return nil, fmt.Errorf("временные файлы могут запрашивать только авторизованные пользователи")
		}

		if file.CreatorUuid == nil || *file.CreatorUuid != ss.UserUuid {
			return nil, fmt.Errorf("%w: файл недоступен текущему пользователю", err_const.ErrAccessDenied)
		}
	}

	// Проверяем доступ к остальным (реальным) объектам
	available, err := s.repo.CheckObjectGettable(ctx, file.ObjectRef)
	if err != nil {
		return nil, fmt.Errorf("не удалось проверить доступ к объекту: %w", err)
	}

	if !available {
		return nil, fmt.Errorf("%w: файл недоступен текущему пользователю", err_const.ErrAccessDenied)
	}

	return file, nil
}

// listByObjectRef получает список файлов (не временных), принадлежащих указанному объекту
func (s *FileStorageService) listNotTempByObjectRef(ctx context.Context, objectRef *models.ObjectRef) (dto.Files, error) {
	// Проверяем доступ к запрашиваемому объекту
	available, err := s.repo.CheckObjectGettable(ctx, objectRef)
	if err != nil {
		return nil, fmt.Errorf("не удалось проверить доступ к объекту: %w", err)
	}

	if !available {
		return nil, fmt.Errorf("%w: запрашиваемый объект недоступен текущему пользователю", err_const.ErrAccessDenied)
	}

	filesList, err := s.repo.ListNotTempByObjectRef(ctx, objectRef)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список файлов: %w", err)
	}

	return filesList, nil
}

// createFile создает файл
func (s *FileStorageService) createFile(ctx context.Context, newFile *dto.NewFile, isTemp bool, content io.Reader) (*dto.File, error) {
	ss, ok := session.GetFromCtx(ctx)
	if !ok || ss.UserUuid == "" {
		return nil, fmt.Errorf("файлы могут создавать только авторизованные пользователи")
	}

	// По умолчанию берем uuid создателя из newFile; если он там не указан - берем uuid текущего пользователя
	creatorUuid := newFile.CreatorUuid
	if creatorUuid == nil {
		creatorUuid = &ss.UserUuid
	}

	// Проверяем, разрешено ли создание файла с указанным объектом objectRef
	err := s.checkFileCreationAccess(ctx, newFile.ObjectRef)
	if err != nil {
		return nil, err
	}

	// Проверяем, можно ли создать файл с указанным типом для объекта objectRef
	err = checkFileTypeAvailable(newFile.Type, newFile.ObjectRef)
	if err != nil {
		return nil, err
	}

	if err := file_type.Check(newFile.Filename, newFile.MIMEType, newFile.Type.AvailableFileTypes()...); err != nil {
		return nil, fmt.Errorf("ошибка при создании файла %s: %w", newFile.Filename, err)
	}

	// Новый файл создается в бакете объекта
	fileCreate := &dto.FileCreate{
		Filename:       newFile.Filename,
		MIMEType:       newFile.MIMEType,
		Description:    newFile.Description,
		CreatorUuid:    creatorUuid,
		ObjectRef:      newFile.ObjectRef,
		Type:           newFile.Type,
		Temp:           isTemp,
		SequenceNumber: nil,
	}

	createdFile, err := s.repo.Create(ctx, fileCreate)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать файл %s: %w", newFile.Filename, err)
	}

	s3File := &dto.S3File{
		BucketName: fs_utils.MakeBucketName(newFile.ObjectRef),
		Filename:   fs_utils.MakeFilename(createdFile),
		Metadata:   fs_utils.MakeMetadata(newFile),
	}

	_, err = s.fileStorage.SaveFile(ctx, s3File, content)
	if err != nil {
		return nil, fmt.Errorf("не удалось сохранить файл %s: %w", newFile.Filename, err)
	}

	return createdFile, nil
}

// Проверка доступа для создания файла с указанным objectRef
func (s *FileStorageService) checkFileCreationAccess(ctx context.Context, objectRef *models.ObjectRef) error {
	// Система может создавать файлы любому объекту
	if access.IsSystem(ctx) {
		return nil
	}

	// Проверяем наличие прав на создание файла с указанным objectRef
	err := checkPermissionsForFileCreation(ctx, objectRef)
	if err != nil {
		return err
	}

	// Проверяем возможность обновления указанного objectRef
	available, err := s.repo.CheckObjectUpdatable(ctx, objectRef)
	if err != nil {
		return fmt.Errorf("не удалось проверить возможность обновления объекта: %w", err)
	}
	if !available {
		return fmt.Errorf("%w: запрашиваемый объект недоступен для обновления текущему пользователю", err_const.ErrAccessDenied)
	}

	return nil
}

// Проверка прав на создание файла в указанном объекте
func checkPermissionsForFileCreation(ctx context.Context, objectRef *models.ObjectRef) error {
	ss, ok := session.GetFromCtx(ctx)
	if !ok {
		return err_const.ErrMissingSession
	}

	var neededPermission permissions.Type
	switch objectRef.Type {
	case models.ObjectTypeChapters:
		// Если объект - глава, то нужно право обновления глав
		neededPermission = permissions.ChaptersUpdate
	case models.ObjectTypeParts:
		// Если объект - часть, то нужно право обновления частей
		neededPermission = permissions.PartsUpdate
	case models.ObjectTypeUsers:
		if objectRef.Ref == ss.UserUuid {
			// Если объект - текущий пользователь, то нужно право обновления собственного пользователя
			neededPermission = permissions.UsersUpdateSelf
		} else {
			// Если объект - любой другой пользователь, то нужно право обновления любого пользователя
			neededPermission = permissions.UsersUpdate
		}
	default:
		return fmt.Errorf("неразрешенный тип объекта: %s", objectRef.Type)
	}

	return access.CheckPermissions(ss.Permissions, neededPermission)
}

// Проверка соответствия типа файла и объекта
func checkFileTypeAvailable(fileType files.Type, objectRef *models.ObjectRef) error {
	typesMapping := map[models.ObjectType][]files.Type{
		models.ObjectTypeChapters: {files.Image},
		models.ObjectTypeParts:    {files.Image},
		models.ObjectTypeUsers:    {files.Avatar},
	}

	availableTypes, ok := typesMapping[objectRef.Type]
	if !ok {
		return fmt.Errorf("неразрешенный тип объекта: %s", objectRef.Type)
	}

	if !slices.Contains(availableTypes, fileType) {
		return fmt.Errorf("тип файла %s недоступен для объекта %s", fileType, objectRef.Type)
	}

	return nil
}
