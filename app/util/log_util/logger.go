package log_util

import (
	"context"
	"fmt"
	"go-otel/app/util/constants"
	"go.opentelemetry.io/otel/trace"
	"log"
	"os"
	"time"
)

const (
	LOG_LEVEL_INFO  = "INFO"
	LOG_LEVEL_ERROR = "ERROR"
)

func LogInfo(ctx context.Context, msgs ...any) {
	msg, logger := buildLog(ctx, LOG_LEVEL_INFO, msgs...)
	logger.Printf(msg)
}

func LogError(ctx context.Context, msgs ...any) {
	msg, logger := buildLog(ctx, LOG_LEVEL_ERROR, msgs...)
	logger.Printf(msg)
}

func buildLog(ctx context.Context, logLevel string, msgs ...any) (string, *log.Logger) {
	message := ""
	for _, msg := range msgs {
		message += fmt.Sprintf("%v", msg)
	}

	logger := ctx.Value(constants.LOGGER_CTX_KEY)
	if logger == nil {
		logger = createGlobalLogger()
	}

	//otel trace
	span := trace.SpanFromContext(ctx)
	message = fmt.Sprintf("[%s](%s,%s)%s,%s",
		logLevel, span.SpanContext().TraceID(), span.SpanContext().SpanID(), time.Now().Format(time.RFC3339),
		message)

	return message, logger.(*log.Logger)
}

func createGlobalLogger() *log.Logger {
	return log.New(os.Stdout, "", 0)
}
