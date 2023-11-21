package template

import (
	"context"
	"go-boilerplate/common/util/constants"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type FuncCallbackPreCheck func(request any) error

type FuncCallbackProcess func(ctx context.Context, request any) (any, error)

type FuncCallbackPostProcess func(request, result any)

type FuncCallbackPostProcessWithReturn func(request, result any) (res any)

func createNewSpan(ctx context.Context, serviceName string) (appendedCtx context.Context, span trace.Span) {
	tracerCtx := ctx.Value(constants.TRACER_CTX_KEY)
	var tracer trace.Tracer
	if tracerCtx == nil {
		tracer = otel.Tracer(constants.APP_NAME)
		ctx = context.WithValue(ctx, constants.TRACER_CTX_KEY, tracer)
	} else {
		tracer = tracerCtx.(trace.Tracer)
	}

	ctx, span = tracer.Start(ctx, serviceName)

	return context.WithValue(ctx, constants.TRACER_CTX_KEY, tracer), span
}
