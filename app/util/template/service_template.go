package template

import (
	"context"
	"encoding/json"
	"fmt"
	"go-otel/app/util/constants"
	log_util2 "go-otel/app/util/log_util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type FuncCallbackPreCheck func(request any) error

type FuncCallbackProcess func(ctx context.Context, request any) (any, error)

type FuncCallbackPostProcess func(request, result any)

func ServiceExec(ctx context.Context, serviceName string, req any, fPre FuncCallbackPreCheck, f FuncCallbackProcess, fPost FuncCallbackPostProcess) error {

	tracerCtx := ctx.Value(constants.TRACER_CTX_KEY)
	var tracer trace.Tracer
	if tracerCtx == nil {
		tracer = otel.Tracer(constants.SERVICE_NAME)
		ctx = context.WithValue(ctx, constants.TRACER_CTX_KEY, tracer)
	} else {
		tracer = tracerCtx.(trace.Tracer)
	}
	ctx = context.WithValue(ctx, constants.TRACER_CTX_KEY, tracer)

	var span trace.Span
	ctx, span = tracer.Start(ctx, serviceName)
	defer func() {
		span.End()
	}()

	log_util2.LogInfo(ctx, fmt.Sprintf("%s service invoke param ", serviceName), req)

	//precheck
	err := fPre(req)
	if err != nil {
		span.RecordError(err)
		return err
	}

	//process
	res, err := f(ctx, req)
	if err != nil {
		span.RecordError(err)
		return err
	}
	log_util2.LogInfo(ctx, fmt.Sprintf("%s service invoke result ", serviceName), res)

	if fPost != nil {
		//postprocess
		fPost(req, res)
	}

	resJson, _ := json.Marshal(res)
	reqJson, _ := json.Marshal(req)

	span.SetAttributes(attribute.KeyValue{Key: "result", Value: attribute.StringValue(string(resJson))})
	span.SetAttributes(attribute.KeyValue{Key: "request", Value: attribute.StringValue(string(reqJson))})

	return nil
}
