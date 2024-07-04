package modules

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"waterfall-backend/internal/modules/domain"
	"waterfall-backend/internal/modules/domain/auth"
	"waterfall-backend/internal/modules/domain/chapter"
	"waterfall-backend/internal/modules/domain/comment"
	"waterfall-backend/internal/modules/domain/part"
	"waterfall-backend/internal/modules/domain/release"
	"waterfall-backend/internal/modules/domain/user"
	"waterfall-backend/internal/modules/features/asynq_server"
	"waterfall-backend/internal/modules/features/config"
	"waterfall-backend/internal/modules/features/echo"
	"waterfall-backend/internal/modules/features/logger"
	"waterfall-backend/internal/modules/features/migrations"
	"waterfall-backend/internal/modules/features/task_manager"
	"waterfall-backend/internal/modules/io/graphql"
	"waterfall-backend/internal/modules/io/http"
	"waterfall-backend/internal/modules/io/worker"
	"waterfall-backend/internal/modules/services/email"
	"waterfall-backend/internal/modules/services/fs"
	"waterfall-backend/internal/modules/services/scheduler"
	"waterfall-backend/internal/modules/services/token"
	"waterfall-backend/internal/modules/stores/db"
	"waterfall-backend/internal/modules/stores/minio"
	"waterfall-backend/internal/modules/stores/redis"
)

var (
	ServerModule = fx.Options(
		// Вспомогательные модули
		logger.Module,
		config.Module,

		// Базы данных
		db.Module,
		minio.Module,
		redis.Module,
		migrations.Module,

		// Модули для http сервера
		echo.Module,
		http.Module,
		graphql.Module,

		// Модуль для создания асинхронных задач
		task_manager.Module,

		// Общие сервисы
		fs.ServiceModule,
		scheduler.ServiceModule,
		token.Module,
		email.ServiceModule,

		// Доменные сервисы
		auth.ServiceModule,
		chapter.ServiceModule,
		part.ServiceModule,
		release.ServiceModule,
		user.ServiceModule,
		comment.ServiceModule,

		// Модуль со всеми доменными сервисами
		domain.ServicesModule,

		// Слой io
		fs.HTTPModule,
		auth.HTTPModule,

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)

	WorkerModule = fx.Options(
		// Вспомогательные модули
		logger.Module,
		config.Module,

		// Базы данных
		db.Module,
		minio.Module,
		redis.Module,
		migrations.Module,

		// Модуль для сервера асинхронных задач
		asynq_server.Module,
		worker.Module,

		// Модуль для создания асинхронных задач
		task_manager.Module,

		// Общие сервисы
		fs.ServiceModule,
		token.Module,
		email.ServiceModule,

		// Доменные сервисы
		auth.ServiceModule,
		chapter.ServiceModule,
		part.ServiceModule,
		release.ServiceModule,
		user.ServiceModule,
		comment.ServiceModule,

		// Слой io
		release.WorkerModule,
		fs.WorkerModule,

		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
	)
)
