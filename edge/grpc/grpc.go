package grpcserver

import (
	"time"

	grpc_xray "github.com/aws/aws-xray-sdk-go/xray"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func NewClientConn(addr, serviceName string) (*grpc.ClientConn, error) {
	return grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // replaced form grpc.WithInsecure()
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    time.Second * 10,
			Timeout: time.Second * 5,
		}),
		grpc.WithUnaryInterceptor(
			grpc_xray.UnaryClientInterceptor(grpc_xray.WithSegmentNamer(grpc_xray.NewFixedSegmentNamer(serviceName))),
		),
	)
}

func NewServer() *grpc.Server {
	return grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				MaxConnectionIdle:     15 * time.Second,
				MaxConnectionAge:      30 * time.Second,
				MaxConnectionAgeGrace: 15 * time.Second,
				Time:                  15 * time.Second,
				Timeout:               10 * time.Second,
			},
		),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             5 * time.Second,
				PermitWithoutStream: true,
			},
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_xray.UnaryServerInterceptor(grpc_xray.WithSegmentNamer(grpc_xray.NewFixedSegmentNamer("myApp"))),
			),
		),
	)
}
