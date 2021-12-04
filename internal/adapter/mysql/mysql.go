package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sundaytycoon/profile.me-server/internal/config"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
	"github.com/sundaytycoon/profile.me-server/pkg/retry"
	"github.com/sundaytycoon/profile.me-server/pkg/testdockercontainer"
	"go.uber.org/dig"
)

type Adapter struct {
	DB *sql.DB
}

func MockNew(mysqlDocker *testdockercontainer.DockerContainer) (*Adapter, error) {
	return New(
		struct {
			dig.In
			ServiceDatabase *config.Database
		}{
			ServiceDatabase: &config.Database{
				Host:     mysqlDocker.ExternalHost,
				Port:     mysqlDocker.ExternalPort,
				User:     mysqlDocker.Get("user"),
				Password: mysqlDocker.Get("password"),
				Name:     mysqlDocker.Get("name"),
				Dialect:  mysqlDocker.Get("dialect"),
			},
		},
	)
}

func New(params struct {
	dig.In
	ServiceDatabase *config.Database
}) (*Adapter, error) {
	op := er.GetOperator()

	db, err := sql.Open("mysql", params.ServiceDatabase.DSN())
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	return &Adapter{
		DB: db,
	}, nil
}

func (a *Adapter) Close() error {
	return a.DB.Close()
}

func (a *Adapter) Conn(ctx context.Context) (*sql.Conn, error) {
	op := er.GetOperator()

	v, err := retry.Retry(5, 1*time.Second, func() (interface{}, error) {
		conn, err := a.DB.Conn(ctx)
		if err != nil {
			return nil, err
		}
		if _, err := a.DB.QueryContext(ctx, "SELECT 1 + 1"); err != nil {
			return nil, err
		}
		return conn, nil
	})
	if err != nil {
		return nil, err
	}

	if value, ok := v.(*sql.Conn); ok {
		return value, nil
	}
	return nil, fmt.Errorf("Failed logic at connection retry logic[%s]", op)
}
