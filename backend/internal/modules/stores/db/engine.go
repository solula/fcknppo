package db

import (
	"fmt"
	"go.uber.org/zap"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/stores/db/ent"
	_ "waterfall-backend/internal/modules/stores/db/ent/runtime"
)

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --target ./ent --feature  sql/execquery,sql/modifier,sql/upsert,intercept,privacy,schema/snapshot ./schema

func NewDBClient(cfg config.Config, logger *zap.Logger) (*ent.Client, error) {
	client, err := connectDB(cfg, logger)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к базе данных: %w", err)
	}

	return client, nil
}
