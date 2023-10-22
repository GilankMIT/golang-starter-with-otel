package middleware

import (
	"fmt"
	"go-boilerplate/common/util/constants"
	"go-boilerplate/common/util/logutil"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

func WithOtelTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		url := r.RequestURI

		ctx = otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(r.Header))

		tracer := otel.Tracer(constants.APP_NAME)

		ctx = context.WithValue(ctx, constants.TRACER_CTX_KEY, tracer)

		ctx = context.WithValue(ctx, constants.LOG_APPENDER_CTX_KEY, logutil.API_APPENDER)

		var span trace.Span
		ctx, span = tracer.Start(ctx, url)
		defer span.End()

		startTime := time.Now()
		logutil.LogInfo(ctx, fmt.Sprintf("new http request %s", url))

		next.ServeHTTP(w, r.WithContext(ctx))

		endTime := time.Now().Sub(startTime).Milliseconds()
		logutil.LogInfo(ctx, fmt.Sprintf("http request end %s, %vms", url, endTime))
	})
}
