package batchdb

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
	DB                                      *sql.DB
	connectionValidation                    bool
	connectionValidationSQL                 string
	connectionValidationRetryTimes          int64
	connectionValidationRetryDuringEachTime time.Duration
}

func New(params struct {
	dig.In
	Config *config.Config
}) (*Adapter, error) {
	op := er.GetOperator()

	mysqlDB, err := mysql.New(params.Config.BatchDB.DSN())
	if err != nil {
		return nil, er.WrapOp(err, op)
	}

	return &Adapter{
		DB:                                      mysqlDB,
		connectionValidation:                    params.Config.BatchDB.ConnectionValidation,
		connectionValidationSQL:                 params.Config.BatchDB.ConnectionValidationSQL,
		connectionValidationRetryTimes:          params.Config.BatchDB.ConnectionValidationRetryTimes,
		connectionValidationRetryDuringEachTime: params.Config.BatchDB.ConnectionValidationRetryDuringEachTime,
	}, nil
}

func (a *Adapter) Conn(ctx context.Context) (*sql.Conn, error) {
	op := er.GetOperator()
	if !a.connectionValidation {
		return a.DB.Conn(ctx)
	}

	v, err := retry.Retry(int(a.connectionValidationRetryTimes), a.connectionValidationRetryDuringEachTime, func() (interface{}, error) {
		conn, err := a.DB.Conn(ctx)
		if err != nil {
			return nil, err
		}
		if _, err := a.DB.QueryContext(ctx, a.connectionValidationSQL); err != nil {
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
	return a.DB.Close()
}
