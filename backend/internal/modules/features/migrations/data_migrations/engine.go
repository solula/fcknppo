package data_migrations

import (
	"github.com/minio/minio-go/v7"
	"waterfall-backend/internal/modules/stores/db/ent"
)

type Engine struct {
	db    *ent.Client
	minio *minio.Client
}

func NewEngine(db *ent.Client, minio *minio.Client) *Engine {
	return &Engine{
		db:    db,
		minio: minio,
	}
}
