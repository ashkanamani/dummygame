package integrationtest

import (
	"fmt"
	"github.com/ashkanamani/dummygame/internal/repository/redis"
	"github.com/ashkanamani/dummygame/pkg/testhelper"
	"github.com/ory/dockertest/v3"
	"os"
	"testing"
)

var redisPort string

func TestMain(m *testing.M) {
	if !testhelper.IsIntegrationTest() {
		return
	}
	pool := testhelper.StartDockerPool()

	// Set up the redis container for tests
	redisResource := testhelper.StartDockerInstance(pool, "redis/redis-stack-server", "latest",
		func(res *dockertest.Resource) error {
			port := res.GetPort("6379/tcp")
			_, err := redis.NewRedisClient(fmt.Sprintf("%s:%s", "127.0.0.1", port))
			return err
		})
	redisPort = redisResource.GetPort("6379/tcp")

	// now run tests
	defer func() {
		_ = redisResource.Close()
	}()
	exitCode := m.Run()
	os.Exit(exitCode)
}
