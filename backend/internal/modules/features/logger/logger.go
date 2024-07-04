package logger

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, zap.AtomicLevel, error) {
	config := zap.NewProductionConfig()
	logger, err := config.Build()

	loggerLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

	return logger, loggerLevel, err
}

func InvokeLogger(_ *zap.Logger, lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return nil //logger.Sync()
		},
	})
}
