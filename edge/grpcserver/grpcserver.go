package grpcserver

import (
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	hchk "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/sundaytycoon/buttons-api/insecure"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

type Server struct {
	grpcServer *grpc.Server
	handlers   []GRPCHandler
}

func New() *Server {
	grpcServer := grpc.NewServer(
		//grpc.UnaryInterceptor(
		//	grpcmiddleware.ChainUnaryServer(
		//		func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		//			if v, ok := req.(*v1pb.UserRe); ok {
		//				fmt.Println("---------")
		//				fmt.Println(v.Name)
		//			}
		//			return handler(ctx, req)
		//		},
		//	),
		//),
		grpc.Creds(credentials.NewServerTLSFromCert(&insecure.Cert)),
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

	return &Server{
		grpcServer: grpcServer,
	}
}

type GRPCHandler interface {
	Register(grpc.ServiceRegistrar)
	Name() string
	Close() error
}

func (s *Server) SetHandlers(handlers ...GRPCHandler) {
	reflection.Register(s.grpcServer)
	hchk.RegisterHealthServer(s.grpcServer, health.NewServer())
	for _, h := range handlers {
		h.Register(s.grpcServer)
	}
}

func (s *Server) Close() error {
	var err error
	for _, h := range s.handlers {
		e := h.Close()
		if e != nil {
			err = errors.Errorf("%v\n[%s]\n%v", err, h.Name(), e)
		}
	}
	s.grpcServer.GracefulStop()
	return err
}

func (s *Server) Start(endpoint string) error {
	op := er.GetOperator()

	grpcListener, err := net.Listen("tcp", endpoint)
	if err != nil {
		return er.WrapOp(err, op)
	}
	log.Info().Str("endpoint", endpoint).Msgf("run grpc application")
	if err = s.grpcServer.Serve(grpcListener); err != nil {
		if err != grpc.ErrServerStopped {
			return er.WrapOp(err, op)
		}
	}
	return nil
}
