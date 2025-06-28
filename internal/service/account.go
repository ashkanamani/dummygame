package service

import (
	"context"
	"errors"
	"github.com/ashkanamani/dummygame/internal/entity"
	"github.com/ashkanamani/dummygame/internal/repository"
	"time"
)

const (
	DefaultState = "home"
)

type AccountService struct {
	accounts repository.AccountRepository
}

func NewAccountService(accounts repository.AccountRepository) *AccountService {
	return &AccountService{
		accounts: accounts,
	}
}

// CreateOrUpdate creates new user in the data store or update the existing one
func (a *AccountService) CreateOrUpdate(ctx context.Context, account entity.Account) (entity.Account, bool, error) {
	savedAccount, err := a.accounts.Get(ctx, account.EntityID())
	// User exists in the database
	if err == nil {
		if savedAccount.Username != account.Username ||
			savedAccount.FirstName != account.FirstName ||
			savedAccount.DisplayName != account.DisplayName {
			savedAccount.Username = account.Username
			savedAccount.FirstName = account.FirstName
			savedAccount.DisplayName = account.DisplayName
			return savedAccount, false, a.accounts.Save(ctx, savedAccount)
		}
	}
	// User does not exist in the database
	if errors.Is(err, repository.ErrorNotFound) {
		account.JoinedAt = time.Now()
		account.State = DefaultState
		return account, true, a.accounts.Save(ctx, account)
	}
	return savedAccount, false, err
}

func (a *AccountService) Update(ctx context.Context, account entity.Account) error {
	return a.accounts.Save(ctx, account)
}
