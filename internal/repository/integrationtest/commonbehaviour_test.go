package integrationtest

import (
	"context"
	"fmt"
	"github.com/ashkanamani/dummygame/internal/entity"
	"github.com/ashkanamani/dummygame/internal/repository"
	"github.com/ashkanamani/dummygame/internal/repository/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

type testType struct {
	ID   string
	Name string
}

func (t testType) EntityID() entity.ID {
	return entity.NewID("testType", t.ID)
}

func TestCommonBehaviourSetAndGet(t *testing.T) {
	redisClient, err := redis.NewRedisClient(fmt.Sprintf("%s:%s", "127.0.0.1", redisPort))
	defer redisClient.Close()
	assert.NoError(t, err)

	ctx := context.Background()
	rcb := repository.NewRedisCommonBehaviour[testType](redisClient)

	err = rcb.Save(ctx, testType{
		ID:   "34",
		Name: "Arshia",
	})
	assert.NoError(t, err)
	err = rcb.Save(ctx, testType{
		ID:   "33",
		Name: "Ashkan",
	})
	assert.NoError(t, err)

	val, err := rcb.Get(ctx, entity.NewID("testType", "34"))
	assert.NoError(t, err)
	assert.Equal(t, "Arshia", val.Name)
	assert.Equal(t, "34", val.ID)

	val, err = rcb.Get(ctx, entity.NewID("testType", "33"))
	assert.NoError(t, err)
	assert.Equal(t, "Ashkan", val.Name)
	assert.Equal(t, "33", val.ID)

	_, err = rcb.Get(ctx, entity.NewID("testType", "32"))
	assert.ErrorIs(t, err, repository.ErrorNotFound)
}
