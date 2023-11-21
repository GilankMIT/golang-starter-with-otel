package template

import (
	"context"
	"encoding/json"
	"fmt"
	"go-boilerplate/common/util/constants"
	"go-boilerplate/common/util/jsonutil"
	"go-boilerplate/common/util/logutil"
	"go.opentelemetry.io/otel/attribute"
)

func ServiceTemplateExec(ctx context.Context, serviceName string, req any, fPre FuncCallbackPreCheck, f FuncCallbackProcess, fPost FuncCallbackPostProcessWithReturn) (err error) {

	ctx, span := createNewSpan(ctx, serviceName)
	defer func() {
		span.End()
	}()

	ctx = context.WithValue(ctx, constants.LOG_APPENDER_CTX_KEY, logutil.SERVICE_APPENDER)

	logutil.LogInfo(ctx, fmt.Sprintf("%s service param ", serviceName), jsonutil.ToJsonString(req))

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
	}

	if fPost != nil && res != nil {
		//postprocess
		res = fPost(req, res)
	}

	logutil.LogInfo(ctx, fmt.Sprintf("%s service result ", serviceName), jsonutil.ToJsonString(res))

	resJson, _ := json.Marshal(res)
	reqJson, _ := json.Marshal(req)

	span.SetAttributes(attribute.KeyValue{Key: "result", Value: attribute.StringValue(string(resJson))})
	span.SetAttributes(attribute.KeyValue{Key: "request", Value: attribute.StringValue(string(reqJson))})

	return nil
}
