package s3

import (
	"context"
	"errors"
	"fmt"
	"github.com/minio/minio-go/v7"
	"io"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/modules/services/fs/dto"
)

type S3 struct {
	client *minio.Client
}

func NewS3(client *minio.Client) *S3 {
	return &S3{
		client: client,
	}
}

func (t *S3) GetFile(ctx context.Context, bucketName string, fileName string) ([]byte, error) {
	reader, err := t.client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, wrap(err)
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, wrap(err)
	}

	return content, nil
}

func (t *S3) CreateBucket(ctx context.Context, bucketName string) error {
	bucketExist, err := t.client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("не удалось проверить существование бакета: %w", wrap(err))
	}

	if !bucketExist {
		if err := t.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
			return fmt.Errorf("не удалось создать бакет: %w", wrap(err))
		}
	}

	return nil
}

func (t *S3) SaveFile(ctx context.Context, file *dto.S3File, content io.Reader) (int, error) {
	info, err := t.client.PutObject(ctx, file.BucketName, file.Filename, content, -1, minio.PutObjectOptions{
		UserMetadata: file.Metadata,
	})
	if err != nil {
		return 0, wrap(err)
	}

	return int(info.Size), nil
}

func (t *S3) CopyFile(ctx context.Context, fromBucketName string, fromFilename string, toBucketName string, toFilename string) error {
	src := minio.CopySrcOptions{
		Bucket: fromBucketName,
		Object: fromFilename,
	}

	dst := minio.CopyDestOptions{
		Bucket: toBucketName,
		Object: toFilename,
	}

	_, err := t.client.CopyObject(ctx, dst, src)
	if err != nil {
		return wrap(err)
	}

	return nil
}

func (t *S3) DeleteFile(ctx context.Context, bucketName string, fileName string) error {
	return wrap(t.client.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{}))
}

func (t *S3) DeleteBucket(ctx context.Context, bucketName string) error {
	return wrap(t.client.RemoveBucketWithOptions(ctx, bucketName, minio.RemoveBucketOptions{ForceDelete: true}))

}

func wrap(err error) error {
	if err == nil {
		return nil
	}

	var minioErr minio.ErrorResponse
	if errors.As(err, &minioErr) {
		if minioErr.StatusCode == 404 {
			return err_const.ErrNotFound
		}
	}

	return err
}
