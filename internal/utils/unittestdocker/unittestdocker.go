package unittestdocker

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/sundaytycoon/buttons-api/internal/utils/er"
)

type DockerContainer struct {
	Container    testcontainers.Container
	ExternalHost string
	ExternalPort string
	InternalPort string
	InternalHost string
	meta         map[string]string
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

// Close have to close when you RunAnyContainer each unit test
func (c *DockerContainer) Close(ctx context.Context) error {
	op := er.GetOperator()
	if err := c.Container.Terminate(ctx); err != nil {
		return er.WrapOp(err, op)
	}
	return nil
}

func New(cont testcontainers.Container, exHost, exPort, inHost, inPort string, m map[string]string) *DockerContainer {
	fmt.Printf("exHost [%s] / exPort [%s] / inHost [%s] / inPort [%s]", exHost, exPort, inHost, inPort)
	fmt.Printf("meta data => %v", m)

	return &DockerContainer{
		Container:    cont,
		ExternalHost: exHost,
		ExternalPort: exPort,
		InternalHost: inHost,
		InternalPort: inPort,
		meta:         m,
	}
}

func RunMySQL(ctx context.Context, containerName string) (*DockerContainer, error) {
	op := er.GetOperator()

	user := "user"
	password := "password"
	database := "search-test"

	contName := fmt.Sprintf("%s-%s", containerName, uuid.New().String())
	cont, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Name:         contName,
			Image:        "mysql:8.0.28",
			ExposedPorts: []string{"3306/tcp"},
			Env: map[string]string{
				"MYSQL_USER":          user,
				"MYSQL_PASSWORD":      password,
				"MYSQL_ROOT_PASSWORD": password,
				"MYSQL_DATABASE":      database,
			},
			WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
		},
		Started: true,
	})
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	port, err := cont.MappedPort(ctx, "3306/tcp")
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	host, err := cont.ContainerIP(ctx)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	if err = cont.Start(ctx); err != nil {
		return nil, er.WrapOp(err, op)
	}
	for {
		if s, err := cont.State(ctx); err != nil {
		} else {
			if s.Running {
				break
			}
		}
		<-time.After(3 * time.Second)
	}

	return New(
		cont,
		"0.0.0.0",
		port.Port(),
		host,
		"3306",
		map[string]string{
			"user":     user,
			"password": password,
			"database": database,
		},
	), nil
}

func RunRedis(ctx context.Context, containerName string) (*DockerContainer, error) {
	op := er.GetOperator()
	contName := fmt.Sprintf("%s-%s", containerName, uuid.New().String())
	cont, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Name:         contName,
			Image:        "redis:6.2.6",
			ExposedPorts: []string{"6379/tcp"},
			WaitingFor:   wait.ForLog("* Ready to accept connections"),
		},
		Started: true,
	})
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	port, err := cont.MappedPort(ctx, "6379/tcp")
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	host, err := cont.ContainerIP(ctx)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	if err = cont.Start(ctx); err != nil {
		return nil, er.WrapOp(err, op)
	}
	for {
		if s, err := cont.State(ctx); err != nil {
		} else {
			if s.Running {
				break
			}
		}
		<-time.After(3 * time.Second)
	}

	return New(
		cont,
		"0.0.0.0",
		port.Port(),
		host,
		"6379",
		nil,
	), nil
}
