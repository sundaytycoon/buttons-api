package servicedb

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/fatih/color"
	"go.uber.org/dig"

	"github.com/sundaytycoon/buttons-api/edge/mysql"
	"github.com/sundaytycoon/buttons-api/ent"
	"github.com/sundaytycoon/buttons-api/internal/config"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

type Adapter struct {
	EntClient *ent.Client
}

func New(params struct {
	dig.In
	Config *config.Config
}) (*Adapter, error) {
	op := er.GetOperator()

	logCtx := func(ctx context.Context, i ...interface{}) {
		if params.Config.Debug {
			Println("MYSQL[LOG]", i)
		}
	}
	drv, err := DB(params.Config.ServiceDB.DSN())
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	dbgDrv := dialect.DebugWithContext(drv, logCtx)
	client := ent.NewClient(ent.Driver(dbgDrv))

	return &Adapter{
		EntClient: client,
	}, nil
}

func DB(dsn string) (*entsql.Driver, error) {
	op := er.GetOperator()
	mysqlDB, err := mysql.New(dsn)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	drv := entsql.OpenDB("mysql", mysqlDB)
	return drv, nil
}

func Println(title string, i ...interface{}) {
	blue := color.New(color.FgBlue).SprintfFunc()
	fmt.Println(blue(title), i)
}

func MockDB() (*entsql.Driver, error) {
	op := er.GetOperator()
	drv, err := entsql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	return drv, nil
}
