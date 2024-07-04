package db

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	"go.uber.org/zap"
	"time"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/trace_driver"

	_ "waterfall-backend/internal/modules/stores/db/ent/runtime"
	// _ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
)

func connectDB(cfg config.Config, logger *zap.Logger) (*ent.Client, error) {
	db, err := sql.Open(dialect.Postgres, getConnectionString(cfg))
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к БД: %w", err)
	}

	err = db.DB().Ping()
	if err != nil {
		return nil, fmt.Errorf("ошибка при подключении к БД: %w", err)
	}

	logLevel := trace_driver.Warn
	if cfg.TraceSQLCommands {
		logLevel = trace_driver.Debug
	}

	// Устанавливаем драйвер с трассировкой SQL команд
	traceDriver := trace_driver.NewTraceDriver(db, NewLogger(
		logger,
		trace_driver.Config{
			SlowThreshold: time.Duration(cfg.SQLSlowThreshold) * time.Second,
			LogLevel:      logLevel,
		}))

	// Формируем опции подключения
	var opts []ent.Option
	opts = append(opts, ent.Driver(traceDriver))

	client := ent.NewClient(opts...)

	return client, nil
}

func getConnectionString(cfg config.Config) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBName,
	)
}
