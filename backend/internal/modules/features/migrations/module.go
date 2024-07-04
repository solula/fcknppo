package migrations

import (
	"go.uber.org/fx"
	"waterfall-backend/internal/modules/features/migrations/data_migrations"
	"waterfall-backend/internal/modules/features/migrations/seeds"
)

var Module = fx.Module("migrations",
	fx.Provide(data_migrations.NewEngine, seeds.NewEngine),
	fx.Invoke(InvokeMigrations),
)
