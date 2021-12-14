package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc/health"
	hchk "google.golang.org/grpc/health/grpc_health_v1"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	v1pb "github.com/sundaytycoon/profile.me-server/gen/go/proto/rpc/v1"
	adapterservicedb "github.com/sundaytycoon/profile.me-server/internal/adapter/servicedb"
	"github.com/sundaytycoon/profile.me-server/internal/config"
	handleruser "github.com/sundaytycoon/profile.me-server/internal/handler/user"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
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

func ServerStart(params struct {
	dig.In
	Config      *config.Config
	UserHandler *handleruser.Handler
}) error {

	fmt.Println(1)
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
			Marshaler: &runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseProtoNames:   true,
					EmitUnpopulated: true,
				},
				UnmarshalOptions: protojson.UnmarshalOptions{
					DiscardUnknown: true,
				},
			},
		}),
	)
	httpEndpoint := net.JoinHostPort("0.0.0.0", "4002")
	httpServer := &http.Server{
		Addr:    httpEndpoint,
		Handler: mux,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
		}
		fmt.Println(2)
	}()

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpcmiddleware.ChainUnaryServer(
				func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
					if v, ok := req.(*v1pb.UserMessage); ok {
						fmt.Println("---------")
						fmt.Println(v.Name)
					}
					return handler(ctx, req)
				},
			),
		),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle:     15 * time.Second,
			MaxConnectionAge:      30 * time.Second,
			MaxConnectionAgeGrace: 15 * time.Second,
			Time:                  15 * time.Second,
			Timeout:               3 * time.Second,
		}),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             5 * time.Second,
				PermitWithoutStream: true,
			},
		),
	)

	fmt.Println(4)
	grpcEndpoint := net.JoinHostPort("0.0.0.0", "4001")
	grpcListener, err := net.Listen("tcp", grpcEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	reflection.Register(grpcServer)
	hchk.RegisterHealthServer(grpcServer, health.NewServer())
	v1pb.RegisterUserServiceServer(grpcServer, params.UserHandler)
	go func() {
		if err = grpcServer.Serve(grpcListener); err != nil && err != grpc.ErrServerStopped {
			fmt.Println(err)
		}
	}()

	go func() {
		<-time.After(3 * time.Second)
		fmt.Println(11)
		conn, err := grpc.DialContext(context.Background(), grpcEndpoint, grpc.WithBlock(), grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(13)
		err = v1pb.RegisterUserServiceHandler(context.Background(), mux, conn)
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(7)

	//if err = httpListener.Close(); err != nil {
	//	fmt.Println(err)
	//}
	//grpcServer.Stop()

	<-time.After(10 * time.Hour)
	//httpServer := httpserver.New(params.Config)
	//httpServer.SetHandler(params.UserHandler)
	//go httpServer.Start()
	//httpServer.Stop()
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
			return Main()
		},
	})
	return c
}
