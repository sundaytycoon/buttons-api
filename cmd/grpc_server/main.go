package main

import (
	"net"

	"github.com/Netflix/go-env"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/health"
	health_check "google.golang.org/grpc/health/grpc_health_v1"

	buttonsapi "github.com/sundaytycoon/buttons-api"
	v1grpc "github.com/sundaytycoon/buttons-api/api/proto/v1"
	edgegoogle "github.com/sundaytycoon/buttons-api/edge/google"
	edgegrpc "github.com/sundaytycoon/buttons-api/edge/grpc"
	edgemysql "github.com/sundaytycoon/buttons-api/edge/mysql"
	"github.com/sundaytycoon/buttons-api/internal/config"
	repositoryauth "github.com/sundaytycoon/buttons-api/internal/domains/repository/auth"
	serviceauth "github.com/sundaytycoon/buttons-api/internal/domains/service/auth"
	handlergrpc "github.com/sundaytycoon/buttons-api/internal/handler/grpc"
	storageservicedb "github.com/sundaytycoon/buttons-api/internal/storage/servicedb"
	storageauth "github.com/sundaytycoon/buttons-api/internal/storage/servicedb/auth"
	utilslifecycle "github.com/sundaytycoon/buttons-api/internal/utils/lifecycle"
	utilsrecovery "github.com/sundaytycoon/buttons-api/internal/utils/recovery"
)

var (
	cfg config.Config
)

func init() {
	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		log.Panic().Err(err).Send()
	}
	log.Info().Interface("config", cfg).Msg("grpc_server start")

	if cfg.Env != buttonsapi.ValueEnvLocal {
	}
}

func main() {
	var options []utilslifecycle.RunOption
	if cfg.Env != buttonsapi.ValueEnvLocal {
		defer xray.SdkDisabled()
	}

	// fatal detection
	defer func() {
		if t := recover(); t != nil {
			err := utilsrecovery.RecoverFn(t)
			if err != nil {
				log.Error().Err(err).Str("type", "fatal").Msg("anomaly terminated")
				if cfg.Env != buttonsapi.ValueEnvLocal {
					// FIXME: noti somewhere
				}
			}
		}
	}()

	// service database
	db := edgemysql.MustNew(
		cfg.ServiceDB.DSN,
		cfg.ServiceDB.MaxIdleConns,
		cfg.ServiceDB.MaxOpenConns,
	)
	defer db.Close()
	serviceDb := storageservicedb.New(db, cfg.Debug)

	grpcServer := edgegrpc.NewServer()

	health_check.RegisterHealthServer(grpcServer, health.NewServer())
	v1grpc.RegisterAuthServiceServer(grpcServer, handlergrpc.NewAuthHandler(serviceauth.New(
		repositoryauth.New(edgegoogle.New(&cfg), storageauth.New(serviceDb)),
	)))

	lis, err := net.Listen("tcp", cfg.ApplicationGRPC.InternalDSN)
	if err != nil {
		log.Panic().Err(err).Msg("failed to start grpc server")
	}
	options = append(options, utilslifecycle.WithGRPC(grpcServer, lis))

	if err = utilslifecycle.Run(options...); err != nil {
		log.Panic().Err(err).Msg("failed for serving grpc server")
	}
}
