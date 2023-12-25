package test

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/redis"
)

type RedisTestContainer struct {
	testcontainers.Container
	MappedPort string
	Host       string
}

func (m RedisTestContainer) GetURI() string {
	return fmt.Sprintf("redis://%s:%s/0?protocol=3", m.Host, m.MappedPort)
}

func NewRedisTestContainer() (*RedisTestContainer, error) {
	ctx := context.Background()

	container, err := redis.RunContainer(ctx,
		testcontainers.WithImage("docker.io/redis:7"),
		redis.WithSnapshotting(10, 1),
		redis.WithLogLevel(redis.LogLevelVerbose),
	)
	if err != nil {
		log.Error(err)
	}
	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "6379")
	if err != nil {
		return nil, err
	}

	return &RedisTestContainer{
		Container:  container,
		MappedPort: mappedPort.Port(),
		Host:       host,
	}, nil
}
