package testdockercontainer

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type DockerContainer struct {
	Container    testcontainers.Container
	ExternalHost string
	ExternalPort string
	InternalPort string
	InternalHost string
	meta         map[string]string
}

func NewDockerContainer(cont testcontainers.Container, exHost, exPort, inHost, inPort string, m map[string]string) *DockerContainer {
	return &DockerContainer{
		Container:    cont,
		ExternalHost: exHost,
		ExternalPort: exPort,
		InternalHost: inHost,
		InternalPort: inPort,
		meta:         m,
	}
}

// Close have to close when you RunAnyContainer each unit test
func (c *DockerContainer) Close(ctx context.Context) error {
	if err := c.Container.Terminate(ctx); err != nil {
		return err
	}
	return nil
}

func (c *DockerContainer) Get(k string) string {
	if v, ok := c.meta[k]; ok {
		return v
	}
	return ""
}

func (c *DockerContainer) Set(k, v string) *DockerContainer {
	c.meta[k] = v
	return c
}

func RunMySQL(ctx context.Context, containerName, schemaName, user, password string) (*DockerContainer, error) {
	contName := fmt.Sprintf("%s-%s", containerName, uuid.New().String())
	cont, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Name:         contName,
			Image:        "mysql:latest",
			ExposedPorts: []string{"3306/tcp"},
			Env: map[string]string{
				"MYSQL_DATABASE":      schemaName,
				"MYSQL_USER":          user,
				"MYSQL_PASSWORD":      password,
				"MYSQL_ROOT_PASSWORD": password,
			},
			WaitingFor: wait.ForAll(
				wait.ForLog("ready for connections."),
				wait.ForListeningPort("3306/tcp"),
			),
		},
		Started: true,
	})

	port, err := cont.MappedPort(ctx, "3306/tcp")
	if err != nil {
		return nil, err
	}
	host, err := cont.ContainerIP(ctx)
	if err != nil {
		return nil, err
	}
	return NewDockerContainer(cont, "0.0.0.0", port.Port(), host, "3306", map[string]string{
		"user":     user,
		"password": password,
		"dialect":  "mysql",
		"schema":   schemaName,
		"name":     schemaName,
	}), nil
}
