package template

import (
	"context"
	"encoding/json"
	"fmt"
	"go-otel/app/util/constants"
	"go-otel/app/util/logutil"
	"go.opentelemetry.io/otel/attribute"
)

func IntegrationTemplateExec(ctx context.Context, serviceName string, req any, fPre FuncCallbackPreCheck, f FuncCallbackProcess, fPost FuncCallbackPostProcess) error {

	ctx, span := createNewSpan(ctx, serviceName)
	defer func() {
		span.End()
	}()

	ctx = context.WithValue(ctx, constants.LOG_APPENDER_CTX_KEY, logutil.INTEGRATION_APPENDER)

	logutil.LogInfo(ctx, fmt.Sprintf("%s client invoke param ", serviceName), req)

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
	logutil.LogInfo(ctx, fmt.Sprintf("%s client invoke result ", serviceName), res)

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
