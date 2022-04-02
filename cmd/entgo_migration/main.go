package main

import (
	"github.com/Netflix/go-env"
	"github.com/rs/zerolog/log"

	"github.com/sundaytycoon/buttons-api/internal/config"
	adapterservicedb "github.com/sundaytycoon/buttons-api/internal/storage/servicedb"

	buttonsapi "github.com/sundaytycoon/buttons-api"
	edgemysql "github.com/sundaytycoon/buttons-api/edge/mysql"
)

var (
	cfg config.Config
)

func init() {
	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Interface("config", cfg).Msg("entgo_migration start")
}

func main() {
	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		log.Panic().Err(err).Send()
	}

	if cfg.Env != buttonsapi.ValueEnvLocal {
		log.Panic().Msg("migration is only allowed at local")
	}

	db := edgemysql.MustNew(cfg.ServiceDB.DSN, cfg.ServiceDB.MaxIdleConns, cfg.ServiceDB.MaxOpenConns)
	defer db.Close()

	serviceDb := adapterservicedb.New(db, cfg.Debug)
	if err := serviceDb.Migrate(); err != nil {
		log.Panic().Err(err).Msg("failed migration at ENV=local")
	}

}
