package repository

import (
	"context"
	"errors"
	"github.com/ashkanamani/dummygame/internal/entity"
	"github.com/ashkanamani/dummygame/pkg/jsonhelper"
	"github.com/redis/rueidis"
	"github.com/sirupsen/logrus"
)

// Check that RedisCommonBehaviour implements CommonBehaviour
var _ CommonBehaviour[entity.Entity] = &RedisCommonBehaviour[entity.Entity]{}

type CommonBehaviour[T entity.Entity] interface {
	Get(ctx context.Context, id entity.ID) (T, error)
	Save(ctx context.Context, entity entity.Entity) error
}

type RedisCommonBehaviour[T entity.Entity] struct {
	client rueidis.Client
}

func NewRedisCommonBehaviour[T entity.Entity]() *RedisCommonBehaviour[T] {
	return &RedisCommonBehaviour[T]{}
}

func (r *RedisCommonBehaviour[T]) Get(ctx context.Context, id entity.ID) (T, error) {
	var t T
	cmd := r.client.B().JsonGet().Key(id.String()).Path(".").Build()
	val, err := r.client.Do(ctx, cmd).ToString()
	if err != nil {
		// handle redis nil error
		if errors.Is(err, rueidis.Nil) {
			return t, ErrNotFound
		}
		logrus.WithError(err).WithField("id", id).Errorln("could not get from redis.")
		return t, err
	}
	return jsonhelper.Decode[T]([]byte(val)), nil
}

func (r *RedisCommonBehaviour[T]) Save(ctx context.Context, entity entity.Entity) error {
	cmd := r.client.B().JsonSet().Key(entity.EntityID().String()).
		Path("$").Value(string(jsonhelper.Encode(entity))).Build()
	if err := r.client.Do(ctx, cmd).Error(); err != nil {
		logrus.WithError(err).WithField("ent", entity).Error("could not save the entity to redis.")
		return err
	}
	return nil
}
