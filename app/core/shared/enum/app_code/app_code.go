package app_code

import (
	"fmt"
	"go-boilerplate/core/shared/constant"
)

var (
	SUCCESS            = NewAppCode(constant.LEVEL_INFO, constant.TYPE_BIZ, "000", "success")
	UNKNOWN_EXCEPTION  = NewAppCode(constant.LEVEL_ERROR, constant.TYPE_BIZ, "fff", "unknown exception")
	ORDER_ALREADY_PAID = NewAppCode(constant.LEVEL_ERROR, constant.TYPE_BIZ, "011", "order already paid")
)

type AppCode struct {
	ErrLevel int
	ErrType  int
	Code     string
	Desc     string
}

func NewAppCode(errLevel, errType int, Code, Desc string) AppCode {
	return AppCode{
		ErrLevel: errLevel,
		ErrType:  errType,
		Code:     Code,
		Desc:     Desc,
	}
}

func (e AppCode) GetFullCode() string {
	return fmt.Sprintf("APP%d%d%s", e.ErrLevel, e.ErrType, e.Code)
}
func (e AppCode) Error() string {
	return e.GetFullCode()
}
