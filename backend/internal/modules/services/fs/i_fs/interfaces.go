package i_fs

import (
	"context"
	"waterfall-backend/internal/models"
)

// IBucketCreator интерфейс создания бакета в s3
type IBucketCreator interface {
	CreateBucket(ctx context.Context, objectRef *models.ObjectRef) error
}
