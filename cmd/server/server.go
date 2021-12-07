package server

import (
	"github.com/spf13/cobra"
	"go.uber.org/dig"

	"github.com/sundaytycoon/profile.me-server/infrastructure/httpserver"
	adapterservicedb "github.com/sundaytycoon/profile.me-server/internal/adapter/servicedb"
	"github.com/sundaytycoon/profile.me-server/internal/config"
	handleruser "github.com/sundaytycoon/profile.me-server/internal/handler/user"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
)

func ServerStart() error {
	// build DI and Invoke server application
	d := dig.New()
	er.PanicError(d.Provide(config.New))
	er.PanicError(d.Provide(adapterservicedb.New))
	er.PanicError(d.Provide(handleruser.New))

	er.PanicError(d.Invoke(Main))

	return nil
}

func Main(params struct {
	dig.In
	Config      *config.Config
	UserHandler *handleruser.Handler
}) error {

	httpServer := httpserver.New(params.Config)
	httpServer.SetHandler(params.UserHandler)

	go httpServer.Start()

	httpServer.Stop()
	return nil
}

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
