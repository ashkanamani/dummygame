package testhelper

import (
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/sirupsen/logrus"
	"os"
)

type RetryFunc func(resource *dockertest.Resource) error

func IsIntegrationTest() bool {
	return os.Getenv("INTEGRATION_TEST") == "true"
}

func StartDockerPool() *dockertest.Pool {
	pool, err := dockertest.NewPool("")
	if err != nil {
		logrus.WithError(err).Fatalf("could not construct pool")
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		logrus.WithError(err).Fatalf("could not connect to docker")
	}
	return pool
}

func StartDockerInstance(pool *dockertest.Pool, image, tag string, retryFunc RetryFunc, env ...string) *dockertest.Resource {
	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: image,
		Tag:        tag,
		Env:        env,
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}

	})
	if err != nil {
		logrus.WithError(err).Fatalf("could not start resource")
	}
	if err := resource.Expire(120); err != nil {
		logrus.WithError(err).Fatalf("could not set resource expiration")
	}

	if err := pool.Retry(func() error {
		return retryFunc(resource)
	}); err != nil {
		logrus.Fatalf("could not connect to the resource")
	}
	return resource
}
