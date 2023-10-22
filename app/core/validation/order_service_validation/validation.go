package order_service_validation

import (
	"go-boilerplate/core/shared/enum/app_code"
	"strings"
)

func TrxIdIsPayable(trxId string) error {
	//sample logic, not actual business requirement
	if strings.Contains(trxId, "99") {
		return app_code.ORDER_ALREADY_PAID
	}
	return nil
}
