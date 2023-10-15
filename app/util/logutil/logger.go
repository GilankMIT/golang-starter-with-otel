package logutil

import (
	"context"
	"fmt"
	"go-otel/app/util/constants"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"log"
	"os"
)

const (
	LOG_LEVEL_INFO  = "INFO"
	LOG_LEVEL_WARN  = "WARN"
	LOG_LEVEL_ERROR = "ERROR"

	//appender
	DEFAULT_APPENDER     = "DEFAULT"
	INFO_APPENDER        = "INFO"
	ERROR_APPENDER       = "ERROR"
	SERVICE_APPENDER     = "SERVICE"
	INTEGRATION_APPENDER = "INTEGRATION"
	API_APPENDER         = "API"
)

var DefaultLogger = NewZapLoggerDefault()
var LOG_APPENDER = map[string]*zap.Logger{
	"":                   DefaultLogger,
	DEFAULT_APPENDER:     DefaultLogger,
	INFO_APPENDER:        NewZapLoggerInfo(),
	ERROR_APPENDER:       NewZapLoggerError(),
	SERVICE_APPENDER:     NewZapLoggerService(),
	INTEGRATION_APPENDER: NewZapLoggerIntegration(),
	API_APPENDER:         NewZapLoggerAPI(),
}

func LogInfo(ctx context.Context, msgs ...any) {
	msg, logger := buildLog(ctx, LOG_LEVEL_INFO, msgs...)

	span := trace.SpanFromContext(ctx)
	logger.Info(msg,
		zap.String("traceId", span.SpanContext().TraceID().String()),
		zap.String("spanId", span.SpanContext().SpanID().String()))
}

func LogError(ctx context.Context, msgs ...any) {
	msg, logger := buildLog(ctx, LOG_LEVEL_ERROR, msgs...)

	span := trace.SpanFromContext(ctx)
	logger.Error(msg, zap.String("traceId", span.SpanContext().TraceID().String()),
		zap.String("spanId", span.SpanContext().SpanID().String()))
}

func LogWarn(ctx context.Context, msgs ...any) {
	msg, logger := buildLog(ctx, LOG_LEVEL_WARN, msgs...)

	span := trace.SpanFromContext(ctx)
	logger.Warn(msg, zap.String("traceId", span.SpanContext().TraceID().String()),
		zap.String("spanId", span.SpanContext().SpanID().String()))
}

func buildLog(ctx context.Context, logLevel string, msgs ...any) (string, *zap.Logger) {
	message := ""
	for _, msg := range msgs {
		message += fmt.Sprintf("%v", msg)
	}

	if logLevel == LOG_LEVEL_ERROR {
		logger := LOG_APPENDER[ERROR_APPENDER]
		return message, logger
	}

	logAppenderCtx, ok := ctx.Value(constants.LOG_APPENDER_CTX_KEY).(string)
	if !ok {
		logAppenderCtx = ""
	}

	logger := LOG_APPENDER[logAppenderCtx]
	return message, logger
}

func createGlobalLogger() *log.Logger {
	return log.New(os.Stdout, "", 0)
}
