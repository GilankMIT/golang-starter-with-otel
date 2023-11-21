package ledger

import (
	"context"
	"go-boilerplate/common/model/service/request"
)

type LedgerService interface {
	Debit(ctx context.Context, command request.LedgerEntryCommand) error
	Credit(ctx context.Context, command request.LedgerEntryCommand) error
}
