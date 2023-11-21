package repository

import "go-boilerplate/core/domain/domainmodel"

type AccountRepository interface {
	LoadAccountByAccountNo(accountNo string) (domainmodel.Account, error)
	UpdateAccount(account domainmodel.Account) error
}

type accountRepoImpl struct {
}

func NewAccountRepository() AccountRepository {
	return accountRepoImpl{}
}

func (a accountRepoImpl) LoadAccountByAccountNo(accountNo string) (domainmodel.Account, error) {
	panic("implement me")
}

func (a accountRepoImpl) UpdateAccount(account domainmodel.Account) error {
	panic("implement me")
}
