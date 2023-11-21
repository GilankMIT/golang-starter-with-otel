package domainmodel

import (
	"go-boilerplate/common/util/money"
	domainenum "go-boilerplate/core/domain/enum"
	"time"
)

type Account struct {
	AccountNo           string
	Direction           string
	AssetType           domainenum.AssetTypeEnum
	CurrentBalance      money.Money
	OpeningTime         time.Time
	LastTransactionTime time.Time
	Status              domainenum.AccountStatusEnum
}
