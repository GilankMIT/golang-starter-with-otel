package request

import "go-boilerplate/common/util/money"

type LedgerEntryCommand struct {
	AccountNo      string
	AssetType      string
	Direction      string
	AdditionalData map[string]string
	Amount         money.Money
}
