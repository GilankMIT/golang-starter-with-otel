package ledger

import (
	"context"
	"errors"
	"go-boilerplate/common/model/service/request"
	domainConstant "go-boilerplate/core/domain/constants"
	"go-boilerplate/core/domain/domainmodel"
	domainenum "go-boilerplate/core/domain/enum"
	"go-boilerplate/core/repository"
)

type ledgerService struct {
	accountRepo repository.AccountRepository
}

func NewLedgerService() LedgerService {
	return &ledgerService{}
}

func (l ledgerService) Debit(ctx context.Context, command request.LedgerEntryCommand) error {
	if command.Direction != domainConstant.DEBIT {
		return errors.New("unconsistent direction")
	}

	account, err := l.fetchAccount(command)
	if err != nil {
		return err
	}

	return l.journal(&account, &command)
}

func (l ledgerService) Credit(ctx context.Context, command request.LedgerEntryCommand) error {
	if command.Direction != domainConstant.CREDIT {
		return errors.New("unconsistent direction")
	}

	account, err := l.fetchAccount(command)
	if err != nil {
		return err
	}

	return l.journal(&account, &command)
}

func (l ledgerService) fetchAccount(command request.LedgerEntryCommand) (domainmodel.Account, error) {
	if command.AssetType != domainenum.ASSET_TYPE_BALANCE.Code {
		return l.accountRepo.LoadAccountByAccountNo(command.AccountNo)
	}

	//fetch by internal account config
	//logic ...

	return domainmodel.Account{}, nil
}

func (l ledgerService) journal(account *domainmodel.Account, commandEntry *request.LedgerEntryCommand) error {

	if commandEntry.Direction == account.Direction {
		account.CurrentBalance.Amount += commandEntry.Amount.Amount
	} else {
		account.CurrentBalance.Amount -= commandEntry.Amount.Amount
	}

	//persist to DB

	return nil
}
