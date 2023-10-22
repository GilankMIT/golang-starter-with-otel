package model

import (
	"go-boilerplate/core/shared/enum/app_code"
	"reflect"
)

func ResolveAppError(err error, baseResp *BaseServiceResponse) {
	if err == nil {
		baseResp.ResponseCode = app_code.SUCCESS.GetFullCode()
		baseResp.IsSuccess = true
		return
	}

	baseResp.IsSuccess = false
	if reflect.TypeOf(err) == reflect.TypeOf(app_code.AppCode{}) {
		baseResp.ResponseCode = err.(app_code.AppCode).GetFullCode()
		baseResp.ResponseDesc = err.(app_code.AppCode).Desc
		return
	}

	baseResp.ResponseCode = app_code.UNKNOWN_EXCEPTION.GetFullCode()
	baseResp.ResponseDesc = app_code.UNKNOWN_EXCEPTION.Desc
	return
}
