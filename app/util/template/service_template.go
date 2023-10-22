package template

import (
	"context"
	"encoding/json"
	"fmt"
	"go-otel-example/app/core/shared/enum/app_code"
	"go-otel-example/app/core/shared/model"
	"go-otel-example/app/util/constants"
	"go-otel-example/app/util/logutil"
	"go.opentelemetry.io/otel/attribute"
)

func ServiceTemplateExec(ctx context.Context, serviceName string, req any, fPre FuncCallbackPreCheck, f FuncCallbackProcess, fPost FuncCallbackPostProcess) (err error) {

	ctx, span := createNewSpan(ctx, serviceName)
	defer func() {
		span.End()
	}()

	ctx = context.WithValue(ctx, constants.LOG_APPENDER_CTX_KEY, logutil.SERVICE_APPENDER)

	logutil.LogInfo(ctx, fmt.Sprintf("%s service invoke param ", serviceName), req)

	//precheck
	err = fPre(req)
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
	logutil.LogInfo(ctx, fmt.Sprintf("%s service invoke result ", serviceName), res)

	if fPost != nil {
		//postprocess
		fPost(req, res)
		err = model.NewError(app_code.SUCCESS, "success")
	}

	resJson, _ := json.Marshal(res)
	reqJson, _ := json.Marshal(req)

	span.SetAttributes(attribute.KeyValue{Key: "result", Value: attribute.StringValue(string(resJson))})
	span.SetAttributes(attribute.KeyValue{Key: "request", Value: attribute.StringValue(string(reqJson))})

	return nil
}
