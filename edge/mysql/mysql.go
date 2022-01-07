package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/sundaytycoon/buttons-api/internal/config"
	"github.com/sundaytycoon/buttons-api/pkg/er"
	"github.com/sundaytycoon/buttons-api/pkg/testdockercontainer"
)

func MockNew(mysqlDocker *testdockercontainer.DockerContainer) (*sql.DB, error) {
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

func New(dsn string) (*sql.DB, error) {
	op := er.GetOperator()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	err = db.Ping()
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	return db, nil
}
