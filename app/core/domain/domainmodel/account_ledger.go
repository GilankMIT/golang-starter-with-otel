package domainmodel

import (
	"go-boilerplate/common/util/money"
	"time"
)

type AccountLedger struct {
	AccountNo string
	Amount    money.Money
	Direction string
	BizDate   time.Time
}
