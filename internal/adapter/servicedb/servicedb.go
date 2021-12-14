package servicedb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/dig"

	"github.com/sundaytycoon/buttons-api/infrastructure/mysql"
	"github.com/sundaytycoon/buttons-api/internal/config"
	"github.com/sundaytycoon/buttons-api/pkg/er"
	"github.com/sundaytycoon/buttons-api/pkg/retry"
)

type Adapter struct {
	MySQL *mysql.Client
}

func New(params struct {
	dig.In
	Config *config.Config
}) (*Adapter, error) {
	op := er.GetOperator()

	mysqlClient, err := mysql.New(params.Config.ServiceDB.DSN())
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	return &Adapter{
		MySQL: mysqlClient,
	}, nil
}

func (a *Adapter) Conn(ctx context.Context) (*sql.Conn, error) {
	op := er.GetOperator()

	v, err := retry.Retry(5, 1*time.Second, func() (interface{}, error) {
		conn, err := a.MySQL.DB.Conn(ctx)
		if err != nil {
			return nil, err
		}
		if _, err := a.MySQL.DB.QueryContext(ctx, "SELECT 1 + 1"); err != nil {
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

func (a *Adapter) Close() error {
	return a.MySQL.Close()
}
