package seeds

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"os"
	"waterfall-backend/internal/constants/files"
	"waterfall-backend/internal/models"
	files2 "waterfall-backend/internal/models/files"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/services/fs/dto"
	fs_utils "waterfall-backend/internal/modules/services/fs/utils"
	"waterfall-backend/internal/modules/stores/db/converters"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/file"
	"waterfall-backend/internal/modules/stores/db/utils"
)

func (s *Engine) files(ctx context.Context, _ config.Config) error {
	publicFileObjectRef := models.ObjectRefPublic
	bucketName := fs_utils.MakeBucketName(publicFileObjectRef)

	err := utils.WithTx(ctx, s.db, func(tx *ent.Tx) error {
		bucketExist, err := s.minio.BucketExists(ctx, bucketName)
		if err != nil {
			return fmt.Errorf("не удалось проверить существование бакета: %w", err)
		}

		// Создаем бакет только если его еще нет
		if !bucketExist {
			if err := s.minio.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
				return fmt.Errorf("не удалось создать бакет: %w", err)
			}
		}

		// Проверяем существование аватара по умолчанию
		exists, err := s.db.File.Query().Where(file.ID(files.DefaultAvatarUuid)).Exist(ctx)
		if err != nil {
			return fmt.Errorf("не удалось проверить существование аватара по умолчанию: %w", err)
		}

		// Если аватар существует, то ничего делать не требуется
		if exists {
			return nil
		}

		avatarFile, err := os.Open("static/default_avatar.jpg")
		if err != nil {
			return fmt.Errorf("не удалось открыть файл аватара: %w", err)
		}
		defer avatarFile.Close()

		newFile := &dto.NewFile{
			Filename:    avatarFile.Name(),
			MIMEType:    "image/jpeg",
			Description: "Сервисный файл аватара пользователей по умолчанию",
		}

		createdFile, err := s.db.File.Create().
			SetID(files.DefaultAvatarUuid).
			SetFilename(newFile.Filename).
			SetMimeType(newFile.MIMEType).
			SetDescription(newFile.Description).
			SetObjectType(publicFileObjectRef.Type).
			SetObjectRef(publicFileObjectRef.Ref).
			SetType(files2.Avatar).
			SetTemp(false).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("не удалось создать файл: %w", err)
		}

		s3File := &dto.S3File{
			BucketName: bucketName,
			Filename:   fs_utils.MakeFilename(converters.ToFileDTO(createdFile)),
			Metadata:   fs_utils.MakeMetadata(newFile),
		}

		_, err = s.minio.PutObject(ctx, s3File.BucketName, s3File.Filename, avatarFile, -1, minio.PutObjectOptions{
			UserMetadata: s3File.Metadata,
		})
		if err != nil {
			return fmt.Errorf("не удалось сохранить файл: %w", err)
		}

		return nil
	})

	return utils.DefaultErrorWrapper(err)
}
