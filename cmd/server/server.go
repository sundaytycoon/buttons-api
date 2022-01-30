package server

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	glueauth "github.com/sundaytycoon/buttons-api/internal/glue/auth"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"google.golang.org/grpc"

	edgegoogle "github.com/sundaytycoon/buttons-api/edge/google"
	"github.com/sundaytycoon/buttons-api/edge/grpcgw"
	"github.com/sundaytycoon/buttons-api/edge/grpcserver"
	adapterbatchdb "github.com/sundaytycoon/buttons-api/internal/adapter/batchdb"
	adapterservicedb "github.com/sundaytycoon/buttons-api/internal/adapter/servicedb"
	"github.com/sundaytycoon/buttons-api/internal/config"
	handlerauth "github.com/sundaytycoon/buttons-api/internal/handler/auth"
	handleruser "github.com/sundaytycoon/buttons-api/internal/handler/user"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

func Main() error {
	// build DI and Invoke server application
	d := dig.New()
	er.PanicError(d.Provide(config.New))
	er.PanicError(d.Provide(edgegoogle.New))
	er.PanicError(d.Provide(adapterservicedb.New))
	er.PanicError(d.Provide(adapterbatchdb.New))
	er.PanicError(d.Provide(handleruser.New))
	er.PanicError(d.Provide(handlerauth.New))

	er.PanicError(d.Invoke(ServerStart))

	return nil
}

type Handler interface {
	Register(grpc.ServiceRegistrar)
	Connect(grpcEndpoint string, mux *runtime.ServeMux) error
	Name() string
	Close() error
}

func ServerStart(params struct {
	dig.In
	Config      *config.Config
	UserHandler *handleruser.Handler
	AuthHandler *handlerauth.Handler
}) error {
	app := grpcserver.New()
	gw := grpcgw.New()
	grpcAppHandlers := []grpcserver.GRPCHandler{
		params.UserHandler,
		glueauth.New(params.AuthHandler),
	}
	grpcGWHandlers := []grpcgw.GRPCHandler{
		params.UserHandler,
		glueauth.New(params.AuthHandler),
	}
	httpEndpoint := net.JoinHostPort(params.Config.HTTPEndPoint.Host, params.Config.HTTPEndPoint.Port)
	grpcEndpoint := net.JoinHostPort(params.Config.GRPCEndPoint.Host, params.Config.GRPCEndPoint.Port)

	go func() {
		app.SetHandlers(grpcAppHandlers...)
		if err := app.Start(grpcEndpoint); err != nil {
			log.Error().
				Err(err).
				Str("endpoint", grpcEndpoint).
				Msg("grpc is stopped")
		}
	}()

	go func() {
		if err := gw.Start(httpEndpoint); err != nil {
			log.Error().
				Err(err).
				Str("endpoint", grpcEndpoint).
				Msg("http mux is stopped")
		}
	}()

	go func() {
		if err := gw.ConnectWithHandlers(grpcEndpoint, grpcGWHandlers...); err != nil {
			log.Error().
				Err(err).
				Str("endpoint", grpcEndpoint).
				Msg("connector[grpc-http] is stopped")
		}
	}()

	shutdown(func() error {
		if err := gw.Close(); err != nil {
			if !er.Is(err, http.ErrServerClosed) {
				log.Fatal().Err(err).Str("closing", "gateway").Send()
			}
		}

		if err := app.Close(); err != nil {
			if !er.Is(err, grpc.ErrServerStopped) {
				log.Fatal().Err(err).Str("closing", "grpc, handlers's connector").Send()
			}
			return err
		}
		return nil
	})
	return nil
}

// Stop When it get sigterm, It'll gracefully closed till request is done or TCP connection reset.
func shutdown(callback func() error) {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSEGV)
	<-stop
	log.Info().Msg("i got the SIGTERM signal, gotta stop")
	log.Info().Msg("Shutdown an application, start!!")
	if err := callback(); err != nil {
		log.Err(err).Msgf("Shutdown an application, shutdown")
	}
	log.Info().Msg("gracefully shutdown!")
	close(stop)
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
			return Main()
		},
	})
	return c
}
