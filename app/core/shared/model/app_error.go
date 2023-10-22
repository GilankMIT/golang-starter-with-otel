package model

import (
	"fmt"
	"go-boilerplate/common/util/constants"
	"go-boilerplate/core/shared/enum/app_code"
)

type AppError struct {
	ErrorCode app_code.AppCode
	Message   string
}

func NewError(errorCode app_code.AppCode, additionalMsg string) AppError {
	return AppError{
		ErrorCode: errorCode,
		Message:   additionalMsg,
	}
}

func (a AppError) Error() string {
	return fmt.Sprintf("%s@%s", constants.APP_NAME, a.ErrorCode.GetFullCode())
}
