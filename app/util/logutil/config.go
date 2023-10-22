package logutil

import (
	"go-otel-example/app/util/constants"
	"go.uber.org/zap"
)

const (
	LOG_BASE_PATH = "./log"

	APP_DEFAULT_LOG_PATH     = LOG_BASE_PATH + "/" + constants.APP_NAME + "-default.log"
	APP_INFO_LOG_PATH        = LOG_BASE_PATH + "/" + constants.APP_NAME + "-info.log"
	APP_ERROR_LOG_PATH       = LOG_BASE_PATH + "/" + constants.APP_NAME + "-error.log"
	APP_SERVICE_LOG_PATH     = LOG_BASE_PATH + "/" + constants.APP_NAME + "-service.log"
	APP_INTEGRATION_LOG_PATH = LOG_BASE_PATH + "/" + constants.APP_NAME + "-integration.log"
	APP_API_LOG_PATH         = LOG_BASE_PATH + "/" + constants.APP_NAME + "-api.log"
)

func BuildZapLoggerConfig(logPath []string) *zap.Logger {
	logConfig := zap.NewProductionConfig()
	logConfig.OutputPaths = logPath

	encodeerConfig := zap.NewProductionEncoderConfig()
	encodeerConfig.CallerKey = ""
	encodeerConfig.TimeKey = ""
	logConfig.EncoderConfig = encodeerConfig
	logger, _ := logConfig.Build()

	return logger
}

// Appender config
func NewZapLoggerDefault() *zap.Logger {
	return BuildZapLoggerConfig([]string{
		"stdout",
		APP_DEFAULT_LOG_PATH,
	})
}

func NewZapLoggerInfo() *zap.Logger {
	return BuildZapLoggerConfig([]string{
		"stdout",
		APP_INFO_LOG_PATH,
	})
}

func NewZapLoggerError() *zap.Logger {
	return BuildZapLoggerConfig([]string{
		"stderr",
		APP_ERROR_LOG_PATH,
	})
}

func NewZapLoggerService() *zap.Logger {
	return BuildZapLoggerConfig([]string{
		"stdout",
		APP_SERVICE_LOG_PATH,
	})
}

func NewZapLoggerIntegration() *zap.Logger {
	return BuildZapLoggerConfig([]string{
		"stdout",
		APP_INTEGRATION_LOG_PATH,
	})
}

func NewZapLoggerAPI() *zap.Logger {
	return BuildZapLoggerConfig([]string{
		"stdout",
		APP_API_LOG_PATH,
	})
}
