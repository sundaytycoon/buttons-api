package servicedb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/dig"

	"github.com/sundaytycoon/buttons-api/edge/mysql"
	"github.com/sundaytycoon/buttons-api/internal/config"
	"github.com/sundaytycoon/buttons-api/pkg/er"
	"github.com/sundaytycoon/buttons-api/pkg/retry"
)

type Adapter struct {
	MySQL                                   *mysql.Client
	ConnectionValidation                    bool
	ConnectionValidationSQL                 string
	ConnectionValidationRetryTimes          int64
	ConnectionValidationRetryDuringEachTime time.Duration
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
		MySQL:                                   mysqlClient,
		ConnectionValidation:                    true,
		ConnectionValidationSQL:                 "SELECT 1+1",
		ConnectionValidationRetryTimes:          5,
		ConnectionValidationRetryDuringEachTime: 1 * time.Second,
	}, nil
}

func (a *Adapter) Conn(ctx context.Context) (*sql.Conn, error) {
	op := er.GetOperator()
	if !a.ConnectionValidation {
		return a.MySQL.DB.Conn(ctx)
	}

	v, err := retry.Retry(int(a.ConnectionValidationRetryTimes), a.ConnectionValidationRetryDuringEachTime, func() (interface{}, error) {
		conn, err := a.MySQL.DB.Conn(ctx)
		if err != nil {
			return nil, err
		}
		if _, err := a.MySQL.DB.QueryContext(ctx, a.ConnectionValidationSQL); err != nil {
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
