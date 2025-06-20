package repository

import (
	"github.com/ashkanamani/dummygame/internal/entity"
	"github.com/redis/rueidis"
)

var _ AccountRepository = &AccountRedisRepository{}

type AccountRedisRepository struct {
	*RedisCommonBehaviour[entity.Account]
}

func NewAccountRedisRepository(client rueidis.Client) *AccountRedisRepository {
	return &AccountRedisRepository{
		NewRedisCommonBehaviour[entity.Account](client),
	}
}
