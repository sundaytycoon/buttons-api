package servicedb

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/fatih/color"
	"github.com/rs/zerolog/log"

	"github.com/sundaytycoon/buttons-api/internal/utils/er"

	"github.com/sundaytycoon/buttons-api/internal/storage/servicedb/ent"
	"github.com/sundaytycoon/buttons-api/internal/storage/servicedb/ent/migrate"
)

type adapter struct {
	db     *sqlx.DB
	client *ent.Client
}

func New(db *sql.DB, debug bool) *adapter {
	drv := entsql.OpenDB("mysql", db)
	var options = []ent.Option{
		ent.Driver(drv),
	}
	if debug {
		logDriver := ent.Driver(dialect.DebugWithContext(drv, func(ctx context.Context, i ...interface{}) {
			Println("MYSQL[LOG]", i)
		}))
		options = append(options, logDriver)
	}

	return &adapter{
		client: ent.NewClient(options...),
		db:     sqlx.NewDb(db, "mysql"),
	}
}

func (a *adapter) Close() error {
	return a.client.Close()
}

func (a *adapter) DB() *sqlx.DB {
	return a.db
}

func (a *adapter) EntClient() *ent.Client {
	return a.client
}

func (a *adapter) WithTx(ctx context.Context, f func(context.Context) error) error {
	tx, err := a.client.Tx(ctx)
	if err != nil {
		return er.WithMessage(err, "begin transaction")
	}

	ctx = ent.NewTxContext(ctx, tx)

	if err = f(ctx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return er.WithMessage(rollbackErr, "rollback transaction")
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return er.WithMessage(err, "commit transaction")
	}

	return nil
}

func (a *adapter) Migrate() error {
	start := time.Now()
	if err := a.client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(false),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Error().
			Err(err).
			Dur("duration", time.Since(start)).
			Msg("Error while in migration")
		return err
	}

	log.Info().Dur("duration", time.Since(start)).Send()
	return nil
}

func Println(title string, i ...interface{}) {
	blue := color.New(color.FgBlue).SprintfFunc()
	fmt.Println(blue(title), i)
}
