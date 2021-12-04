package server

import (
	"github.com/spf13/cobra"
	adaptermysql "github.com/sundaytycoon/profile.me-server/internal/adapter/mysql"
	"github.com/sundaytycoon/profile.me-server/internal/config"
	"github.com/sundaytycoon/profile.me-server/internal/server"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
	"go.uber.org/dig"
)

func ServerCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "server",
		Aliases: []string{"s"},
		Short:   "server application",
		RunE: func(c *cobra.Command, _ []string) error {
			return c.Help()
		},
	}
	c.AddCommand(&cobra.Command{
		Use:     "start",
		Aliases: []string{"s"},
		Short:   "start api application",
		RunE: func(c *cobra.Command, _ []string) error {
			return ServerStart()
		},
	})
	return c
}

func ServerStart() error {
	// build DI and Invoke server application
	d := dig.New()
	er.PanicError(d.Provide(config.New))
	er.PanicError(d.Provide(server.New))
	er.PanicError(d.Provide(adaptermysql.New))
	er.PanicError(d.Invoke(Main))

	return nil
}

func Main(params struct {
	dig.In
	Server *server.Server
}) error {
	go params.Server.Start()

	params.Server.Stop()
	return nil
}
