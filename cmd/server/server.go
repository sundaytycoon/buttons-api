package server

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"google.golang.org/grpc"

	"github.com/sundaytycoon/buttons-api/edge/grpcgw"
	"github.com/sundaytycoon/buttons-api/edge/grpcserver"
	adapterservicedb "github.com/sundaytycoon/buttons-api/internal/adapter/servicedb"
	"github.com/sundaytycoon/buttons-api/internal/config"
	handleruser "github.com/sundaytycoon/buttons-api/internal/handler/user"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

func Main() error {
	// build DI and Invoke server application
	d := dig.New()
	er.PanicError(d.Provide(config.New))
	er.PanicError(d.Provide(adapterservicedb.New))
	er.PanicError(d.Provide(handleruser.New))

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
}) error {
	app := grpcserver.New()
	gw := grpcgw.New()
	grpcAppHandlers := []grpcserver.GRPCHandler{params.UserHandler}
	grpcGWHandlers := []grpcgw.GRPCHandler{params.UserHandler}
	httpEndpoint := net.JoinHostPort(params.Config.HTTPEndPoint.Host, params.Config.HTTPEndPoint.Port)
	grpcEndpoint := net.JoinHostPort(params.Config.GRPCEndPoint.Host, params.Config.GRPCEndPoint.Port)

	go func() {
		app.SetHandlers(grpcAppHandlers...)
		if err := app.Start(grpcEndpoint); err != nil {
			fmt.Println("grpc Start")
			fmt.Println(err)
		}
	}()

	go func() {
		if err := gw.Start(httpEndpoint); err != nil {
			fmt.Println("http mux gateway Start")
			fmt.Println(err)

		}
	}()

	go func() {
		if err := gw.ConnectWithHandlers(grpcEndpoint, grpcGWHandlers...); err != nil {
			fmt.Println("ConnectWithHandlers")
			fmt.Println(err)
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

//
//func ServerStart(params struct {
//	dig.In
//	Config      *config.Config
//	UserHandler *handleruser.Handler
//}) error {
//
//	fmt.Println(1)
//	mux := runtime.NewServeMux(
//		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
//			Marshaler: &runtime.JSONPb{
//				MarshalOptions: protojson.MarshalOptions{
//					UseProtoNames:   true,
//					EmitUnpopulated: true,
//				},
//				UnmarshalOptions: protojson.UnmarshalOptions{
//					DiscardUnknown: true,
//				},
//			},
//		}),
//	)
//	httpEndpoint := net.JoinHostPort("0.0.0.0", "4002")
//	httpServer := &http.Server{
//		Addr:    httpEndpoint,
//		Handler: mux,
//	}
//
//	go func() {
//		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
//			fmt.Println(err)
//		}
//		fmt.Println(2)
//	}()
//
//	grpcServer := grpc.NewServer(
//		grpc.UnaryInterceptor(
//			grpcmiddleware.ChainUnaryServer(
//				func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//					if v, ok := req.(*v1pb.UserMessage); ok {
//						fmt.Println("---------")
//						fmt.Println(v.Name)
//					}
//					return handler(ctx, req)
//				},
//			),
//		),
//		grpc.KeepaliveParams(keepalive.ServerParameters{
//			MaxConnectionIdle:     15 * time.Second,
//			MaxConnectionAge:      30 * time.Second,
//			MaxConnectionAgeGrace: 15 * time.Second,
//			Time:                  15 * time.Second,
//			Timeout:               3 * time.Second,
//		}),
//		grpc.KeepaliveEnforcementPolicy(
//			keepalive.EnforcementPolicy{
//				MinTime:             5 * time.Second,
//				PermitWithoutStream: true,
//			},
//		),
//	)
//
//	fmt.Println(4)
//	grpcEndpoint := net.JoinHostPort("0.0.0.0", "4001")
//	grpcListener, err := net.Listen("tcp", grpcEndpoint)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	reflection.Register(grpcServer)
//	hchk.RegisterHealthServer(grpcServer, health.NewServer())
//	v1pb.RegisterUserServiceServer(grpcServer, params.UserHandler)
//	go func() {
//		if err = grpcServer.Serve(grpcListener); err != nil && err != grpc.ErrServerStopped {
//			fmt.Println(err)
//		}
//	}()
//
//	go func() {
//		<-time.After(3 * time.Second)
//		fmt.Println(11)
//		conn, err := grpc.DialContext(context.Background(), grpcEndpoint, grpc.WithBlock(), grpc.WithInsecure())
//		if err != nil {
//			fmt.Println(err)
//		}
//
//		fmt.Println(13)
//		err = v1pb.RegisterUserServiceHandler(context.Background(), mux, conn)
//		if err != nil {
//			fmt.Println(err)
//		}
//	}()
//
//	fmt.Println(7)
//
//	//if err = httpListener.Close(); err != nil {
//	//	fmt.Println(err)
//	//}
//	//grpcServer.Stop()
//
//	<-time.After(10 * time.Hour)
//	//httpServer := httpserver.New(params.Config)
//	//httpServer.SetHandler(params.UserHandler)
//	//go httpServer.Start()
//	//httpServer.Stop()
//	return nil
//}

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
