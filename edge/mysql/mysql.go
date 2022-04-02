package mysql

import (
	"database/sql"

	"github.com/aws/aws-xray-sdk-go/xray"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

func MustNew(dsn string, maxIdleConns, maxOpenConns int) *sql.DB {
	db, err := xray.SQLContext("mysql", dsn)
	if err != nil {
		log.Panic().Err(err).Msgf("cannot connect with service db [%s]", dsn)
	}
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)

	err = db.Ping()
	if err != nil {
		log.Panic().Err(err).
			Str("dsn", dsn).
			Int("maxIdleConns", maxIdleConns).
			Int("maxOpenConns", maxOpenConns).
			Send()
	}
	return db
}
