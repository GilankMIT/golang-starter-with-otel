package middleware

import (
	"fmt"
	"go-otel/app/util/constants"
	log_util2 "go-otel/app/util/log_util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
	"net/http"
	"time"
)

func WithOtelTrace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		url := r.RequestURI

		tracer := otel.Tracer(constants.SERVICE_NAME)
		ctx = context.WithValue(ctx, constants.TRACER_CTX_KEY, tracer)

		var span trace.Span
		ctx, span = tracer.Start(ctx, url)
		defer span.End()

		startTime := time.Now()
		log_util2.LogInfo(ctx, fmt.Sprintf("new http request %s", url))

		next.ServeHTTP(w, r.WithContext(ctx))

		endTime := time.Now().Sub(startTime).Microseconds()
		log_util2.LogInfo(ctx, fmt.Sprintf("http request end %s, %vms", url, endTime))
	})
}
