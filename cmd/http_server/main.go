package main

import (
	"net/http"

	"github.com/Netflix/go-env"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/rs/zerolog/log"

	httphelper "github.com/sundaytycoon/buttons-api/internal/handler/http/helper"

	buttonsapi "github.com/sundaytycoon/buttons-api"
	oapiv1 "github.com/sundaytycoon/buttons-api/api/oapi/v1"
	edgechi "github.com/sundaytycoon/buttons-api/edge/chi"
	edgegoogle "github.com/sundaytycoon/buttons-api/edge/google"
	edgemysql "github.com/sundaytycoon/buttons-api/edge/mysql"
	"github.com/sundaytycoon/buttons-api/internal/config"
	repositoryauth "github.com/sundaytycoon/buttons-api/internal/domains/repository/auth"
	serviceauth "github.com/sundaytycoon/buttons-api/internal/domains/service/auth"
	handlerhttp "github.com/sundaytycoon/buttons-api/internal/handler/http"
	storageservicedb "github.com/sundaytycoon/buttons-api/internal/storage/servicedb"
	storageauth "github.com/sundaytycoon/buttons-api/internal/storage/servicedb/auth"
	"github.com/sundaytycoon/buttons-api/internal/utils"
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
	log.Info().Interface("config", cfg).Msg("http_server start")

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

	httpHandler := handlerhttp.New(
		serviceauth.New(
			repositoryauth.New(
				edgegoogle.New(&cfg),
				storageauth.New(serviceDb),
			),
		),
	)

	edgeChi := edgechi.New()
	oapiv1.HandlerWithOptions(httpHandler, oapiv1.ChiServerOptions{
		BaseURL:    "/api/v1",
		BaseRouter: edgeChi,
		Middlewares: []oapiv1.MiddlewareFunc{
			// Priority is highest
			func(next http.HandlerFunc) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					// FIXME: chi cannot set ContentType in custom apis
					w.Header().Set(oapiv1.HeaderContentType, oapiv1.MIMEApplicationJSONCharsetUTF8)
					r = r.WithContext(utils.WithRequestMetadataFromHTTP(r, cfg.Env))
					next.ServeHTTP(w, r)
				}
			},
		},
		ErrorHandlerFunc: httphelper.OAPIErrorHandler,
	})

	options = append(options, utilslifecycle.WithHTTP(&http.Server{
		Addr:    cfg.ApplicationHTTP.InternalDSN,
		Handler: edgeChi,
	}))
	if err := utilslifecycle.Run(options...); err != nil {
		log.Panic().Err(err).Msg("failed for serving http server")
	}
}
