package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sundaytycoon/profile.me-server/internal/config"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
	"go.uber.org/dig"
)

type Adapter struct {
	db *sql.DB
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
		db: db,
	}, nil
}

func (a *Adapter) Close() error {
	return a.db.Close()
}
