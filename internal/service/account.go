package service

import (
	"context"
	"github.com/ashkanamani/dummygame/internal/entity"
	"github.com/ashkanamani/dummygame/internal/repository"
)

type AccountService struct {
	accounts repository.AccountRepository
}

func NewAccountService(accounts repository.AccountRepository) *AccountService {
	return &AccountService{
		accounts: accounts,
	}
}

func (a *AccountService) Register(ctx context.Context, account entity.Account) error {
	return a.accounts.Save(ctx, account)
}
