package db

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"time"
	"waterfall-backend/internal/models/session"
	logger_pkg "waterfall-backend/internal/modules/features/logger"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/trace_driver"
)

func NewLogger(lg *zap.Logger, config trace_driver.Config) trace_driver.ILogger {
	return &logger{
		lg:     lg,
		Config: config,
	}
}

type logger struct {
	lg *zap.Logger
	trace_driver.Config
}

func (l *logger) LogMode(level trace_driver.LogLevel) trace_driver.ILogger {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

func (l logger) Error(ctx context.Context, msg string, data map[string]any) {
	if l.LogLevel >= trace_driver.Error {
		l.lg.Error(msg, l.makeFields(ctx, data)...)
	}
}

func (l logger) Warn(ctx context.Context, msg string, data map[string]any) {
	if l.LogLevel >= trace_driver.Warn {
		l.lg.Warn(msg, l.makeFields(ctx, data)...)
	}
}
func (l logger) Info(ctx context.Context, msg string, data map[string]any) {
	if l.LogLevel >= trace_driver.Info {
		l.lg.Info(msg, l.makeFields(ctx, data)...)
	}
}

func (l logger) Debug(ctx context.Context, msg string, data map[string]any) {
	if l.LogLevel >= trace_driver.Debug {
		l.lg.Debug(msg, l.makeFields(ctx, data)...)
	}
}

func (l logger) Trace(ctx context.Context, begin time.Time, fc func() string, err error, data map[string]any) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		strElapced := fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)

		fields := l.makeFields(ctx, data)

		switch {
		case err != nil && l.LogLevel >= trace_driver.Error:
			var loggerFunc func(msg string, fields ...zap.Field)
			logger := l.lg.With(fields...)
			if ent.IsNotFound(err) {
				loggerFunc = logger.Debug
			} else {
				loggerFunc = logger.Error
			}

			sql := fc()
			loggerFunc("Trace SQL",
				zap.String("sql", sql),
				zap.String("elapsedTime", strElapced),
				zap.Error(err))
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= trace_driver.Warn:
			sql := fc()
			l.lg.With(fields...).With(
				zap.String("sql", sql),
				zap.String("elapsedTime", strElapced),
				zap.Duration("slowThreshold", l.SlowThreshold)).
				Warn("Trace SQL")
		case l.LogLevel >= trace_driver.Debug:
			sql := fc()
			l.lg.With(fields...).With(
				zap.String("sql", sql),
				zap.String("elapsedTime", strElapced)).
				Debug("Trace SQL")
		}
	}
}

func (l logger) makeFields(ctx context.Context, data map[string]any) []zap.Field {
	var fields []zap.Field

	for k, v := range data {
		fields = append(fields, zap.Any(k, v))
	}

	ss, ok := session.GetFromCtx(ctx)
	if !ok {
		return fields
	}

	fields = append(fields, logger_pkg.UserFields(ss)...)

	return fields
}
