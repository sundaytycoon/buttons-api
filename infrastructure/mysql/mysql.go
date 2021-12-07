package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sundaytycoon/profile.me-server/internal/config"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
	"github.com/sundaytycoon/profile.me-server/pkg/testdockercontainer"
)

type Client struct {
	DB *sql.DB
}

func MockNew(mysqlDocker *testdockercontainer.DockerContainer) (*Client, error) {
	d := &config.Database{
		Host:     mysqlDocker.ExternalHost,
		Port:     mysqlDocker.ExternalPort,
		User:     mysqlDocker.Get("user"),
		Password: mysqlDocker.Get("password"),
		Name:     mysqlDocker.Get("name"),
		Dialect:  mysqlDocker.Get("dialect"),
	}
	return New(d.DSN())
}

func New(dsn string) (*Client, error) {
	op := er.GetOperator()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	return &Client{
		DB: db,
	}, nil
}

func (a *Client) Close() error {
	return a.DB.Close()
}
