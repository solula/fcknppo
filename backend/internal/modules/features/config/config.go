package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// Config ...
type Config struct {
	JWTSecret                      string        `envconfig:"JWT_SECRET"`
	JWTAccessTokenLifetime         time.Duration `envconfig:"JWT_ACCESS_TOKEN_LIFETIME" default:"30m"`        // 30 минут
	JWTRefreshTokenLifetime        time.Duration `envconfig:"JWT_REFRESH_TOKEN_LIFETIME" default:"720h"`      // 720 часов это 30 дней
	EmailVerificationTokenLifetime time.Duration `envconfig:"EMAIL_VERIFICATION_TOKEN_LIFETIME" default:"4h"` // 4 часа по умолчанию

	GoogleClientID string `envconfig:"GOOGLE_CLIENT_ID"`
	VKServiceToken string `envconfig:"VK_SERVICE_TOKEN"`

	DBUser           string `envconfig:"DB_USER"`
	DBPass           string `envconfig:"DB_PASS"`
	DBHost           string `envconfig:"DB_HOST" default:"postgres"`
	DBPort           string `envconfig:"DB_PORT" default:"5432"`
	DBName           string `envconfig:"DB_NAME" default:"server_db"`
	DBSSLMode        string `envconfig:"DB_SSL_MODE" default:"disable" validate:"oneof=disable enable"`
	SQLSlowThreshold int    `envconfig:"SQL_SLOW_THRESHOLD" default:"600"`
	TraceSQLCommands bool   `envconfig:"TRACE_SQL_COMMANDS" default:"false"`
	AutoMigrate      bool   `envconfig:"AUTO_MIGRATE" default:"false"`

	MinioHost      string `envconfig:"MINIO_HOST" default:"s3"`
	MinioPort      string `envconfig:"MINIO_PORT" default:"9000"`
	MinioAccessKey string `envconfig:"MINIO_ACCESS_KEY"`
	MinioSecretKey string `envconfig:"MINIO_SECRET_KEY"`

	RedisHost string `envconfig:"REDIS_HOST" default:"redis"`
	RedisPort string `envconfig:"REDIS_PORT" default:"6379"`

	FilesTTL time.Duration `envconfig:"FILES_TTL" default:"4h"`

	LoggerType string `envconfig:"LOGGER_TYPE" default:"prod" validate:"oneof=dev prod"`
	LogLevel   string `envconfig:"LOG_LEVEL" default:"info" validate:"oneof=debug info warn error dpanic panic fatal"`

	HTTPServerHost         string        `envconfig:"HTTP_SERVER_HOST" default:"0.0.0.0"`
	HTTPServerPort         string        `envconfig:"HTTP_SERVER_PORT" default:"8080"`
	HTTPServerReadTimeOut  time.Duration `envconfig:"HTTP_SERVER_READ_TIMEOUT" default:"10m"`
	HTTPServerWriteTimeOut time.Duration `envconfig:"HTTP_SERVER_WRITE_TIMEOUT" default:"13m"`

	SystemEmailAddress  string `envconfig:"SYSTEM_EMAIL_ADDRESS"`
	SystemEmailPassword string `envconfig:"SYSTEM_EMAIL_PASSWORD"`

	DevMode       bool   `envconfig:"DEV_MODE" default:"false"`
	AdminEmail    string `envconfig:"ADMIN_EMAIL"`
	AdminPassword string `envconfig:"ADMIN_PASSWORD"`

	FrontendUrl string `envconfig:"FRONTEND_URL"` // Адрес фронтенда для формирования ссылок не него (без завершающего слеша)
}

func NewConfig(logger *zap.Logger, logLevel zap.AtomicLevel) (Config, error) {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		return Config{}, err
	}

	err = replaceLogger(logger, logLevel, config)
	if err != nil {
		return Config{}, err
	}

	//logger.Info("Получена конфигурация", zap.Any("config", config))

	return config, nil
}

func replaceLogger(logger *zap.Logger, logLevel zap.AtomicLevel, cfg Config) error {
	// Принудительно инициализируем уровень логирования из конфигурации
	err := logLevel.UnmarshalText([]byte(cfg.LogLevel))
	if err != nil {
		return err
	}

	// Подменяем логгер в зависимости от конфигурации
	var loggerConfig zap.Config
	switch cfg.LoggerType {
	case "prod":
		loggerConfig = prodLoggerConfig()
	case "dev":
		loggerConfig = devLoggerConfig()
	default:
		loggerConfig = prodLoggerConfig()
	}

	newLogger, err := loggerConfig.Build()
	if err != nil {
		return fmt.Errorf("не удалось подменить логгер: %w", err)
	}

	*logger = *newLogger

	return nil
}

func prodLoggerConfig() zap.Config {
	cfg := zap.NewProductionConfig()

	cfg.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  zapcore.OmitKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	return cfg
}

func devLoggerConfig() zap.Config {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.StacktraceKey = zapcore.OmitKey

	return cfg
}
