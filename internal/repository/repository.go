package repository

import (
	"context"
	"errors"
	"github.com/ashkanamani/dummygame/internal/entity"
)

var ErrorNotFound = errors.New("entity not found")

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, entity T) error
}
