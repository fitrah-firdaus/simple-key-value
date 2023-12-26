package test

import (
	"context"
	"github.com/gofiber/fiber/v2/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MySQLTestContainer struct {
	testcontainers.Container
	MappedPort string
	Host       string
}

func (m MySQLTestContainer) GetURI() string {
	return "root:password@tcp(" + m.Host + ":" + m.MappedPort + ")/kv"
}

func NewMySQLTestContainer() (*MySQLTestContainer, error) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        "mysql:8.2.0",
		ExposedPorts: []string{"3306/tcp", "33060/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": "password",
			"MYSQL_DATABASE":      "kv",
		},
		WaitingFor: wait.ForListeningPort("3306/tcp"),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Error(err)
	}
	host, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "3306")
	if err != nil {
		return nil, err
	}

	return &MySQLTestContainer{
		Container:  container,
		MappedPort: mappedPort.Port(),
		Host:       host,
	}, nil
}
