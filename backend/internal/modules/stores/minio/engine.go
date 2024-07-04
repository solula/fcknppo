package minio

import (
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"waterfall-backend/internal/modules/features/config"
)

func NewClient(cfg config.Config) (*minio.Client, error) {
	client, err := minio.New(cfg.MinioHost+":"+cfg.MinioPort, &minio.Options{
		Creds: credentials.NewStaticV4(cfg.MinioAccessKey, cfg.MinioSecretKey, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к Minio S3: %w", err)
	}

	return client, nil
}
