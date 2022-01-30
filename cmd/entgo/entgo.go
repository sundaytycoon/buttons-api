// Package entd
/*
- https://entgo.io/docs/generating-ent-schemas
- daily
*/
package entgo

import (
	"context"
	"github.com/sundaytycoon/buttons-api/ent/migrate"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/dig"

	adapterservicedb "github.com/sundaytycoon/buttons-api/internal/adapter/servicedb"
	"github.com/sundaytycoon/buttons-api/internal/config"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

func MigrationCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "entgo",
		Aliases: []string{"e"},
		Short:   "entgo go processs",
		RunE: func(c *cobra.Command, _ []string) error {
			return c.Help()
		},
	}
	c.AddCommand(&cobra.Command{
		Use:     "migration",
		Aliases: []string{"m"},
		Short:   "entgo go script migration",
		RunE: func(c *cobra.Command, _ []string) error {
			return Main()
		},
	})
	return c
}

func Main() error {
	// build DI and Invoke server application
	d := dig.New()
	er.PanicError(d.Provide(config.New))
	er.PanicError(d.Provide(adapterservicedb.New))

	er.PanicError(d.Invoke(BasicMigration))

	return nil
}

func BasicMigration(params struct {
	dig.In
	ServiceDB *adapterservicedb.Adapter
}) error {
	ctx := context.Background()
	start := time.Now()
	if err := params.ServiceDB.EntClient.Schema.Create(
		ctx,
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

	log.Info().
		Dur("duration", time.Since(start)).Send()

	return nil
}
