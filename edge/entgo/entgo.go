package entgo

import (
	"database/sql"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"go.uber.org/dig"

	"github.com/sundaytycoon/buttons-api/ent"
	"github.com/sundaytycoon/buttons-api/internal/config"
)

func New(
	params struct {
		dig.In
		Config *config.Config
	},
) (*ent.Client, error) {
	dsn := params.Config.ServiceDB.DSN()
	db, err := sql.Open(params.Config.ServiceDB.Dialect, dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(2)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(time.Second * 120)
	db.SetConnMaxLifetime(time.Second * 120)
	drv := entsql.OpenDB("mysql", db)
	return ent.NewClient(ent.Driver(drv)), nil
}
