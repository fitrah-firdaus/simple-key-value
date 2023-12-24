package test

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

type MongoTestContainer struct {
	testcontainers.Container
	MappedPort string
	Host       string
}

func (m MongoTestContainer) GetURI() string {
	return fmt.Sprintf("mongodb://%s:%s/", m.Host, m.MappedPort)
}

func NewMongoTestContainer() (*MongoTestContainer, error) {
	ctx := context.Background()
	container, err := mongodb.RunContainer(ctx, testcontainers.WithImage("mongo:6"))
	if err != nil {
		log.Error(err)
	}
	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "27017")
	if err != nil {
		return nil, err
	}

	return &MongoTestContainer{
		Container:  container,
		MappedPort: mappedPort.Port(),
		Host:       host,
	}, nil
}
