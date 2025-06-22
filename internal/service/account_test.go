package service

import (
	"context"
	"github.com/ashkanamani/dummygame/internal/entity"
	"github.com/ashkanamani/dummygame/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAccountService_CreateOrUpdateWithUserExists(t *testing.T) {
	accRep := repository.NewMockAccountRepository(t)
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 33)).Return(
		entity.Account{ID: 33, FirstName: "Ashkan"}, nil,
	).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "RadioNishtman"
	})).Return(nil).Once()

	newAcc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        33,
		FirstName: "RadioNishtman",
	})
	assert.NoError(t, err)
	assert.Equal(t, false, created)
	assert.Equal(t, int64(33), newAcc.ID)
	assert.Equal(t, "RadioNishtman", newAcc.FirstName)
}

func TestAccountService_CreateOrUpdateWithUserDoesNotExist(t *testing.T) {
	accRep := repository.NewMockAccountRepository(t)
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 33)).Return(
		entity.Account{}, repository.ErrorNotFound,
	).Once()

	accRep.On("Save", mock.Anything, mock.MatchedBy(func(acc entity.Account) bool {
		return acc.FirstName == "RadioNishtman"
	})).Return(nil).Once()

	newAcc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        33,
		FirstName: "RadioNishtman",
	})
	assert.NoError(t, err)
	assert.Equal(t, true, created)
	assert.Equal(t, int64(33), newAcc.ID)
	assert.Equal(t, "RadioNishtman", newAcc.FirstName)
}

func TestAccountService_CreateOrUpdateWithUserHasNotChanged(t *testing.T) {
	accRep := repository.NewMockAccountRepository(t)
	s := NewAccountService(accRep)

	accRep.On("Get", mock.Anything, entity.NewID("account", 33)).Return(
		entity.Account{ID: 33, FirstName: "RadioNishtman"}, nil,
	).Once()

	newAcc, created, err := s.CreateOrUpdate(context.Background(), entity.Account{
		ID:        33,
		FirstName: "RadioNishtman",
	})
	assert.NoError(t, err)
	assert.Equal(t, false, created)
	assert.Equal(t, int64(33), newAcc.ID)
	assert.Equal(t, "RadioNishtman", newAcc.FirstName)
}
