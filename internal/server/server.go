package server

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"
	"github.com/sundaytycoon/profile.me-server/internal/config"
	serviceuser "github.com/sundaytycoon/profile.me-server/internal/core/service/user"
	handleruser "github.com/sundaytycoon/profile.me-server/internal/handler/user"
	repositoryuser "github.com/sundaytycoon/profile.me-server/internal/repository/user"
	"github.com/sundaytycoon/profile.me-server/pkg/er"
	"go.uber.org/dig"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	http *http.Server
}

func New(params struct {
	dig.In
	Config *config.Config

	UserRepository *repositoryuser.Repository
}) *Server {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	userHandler := handleruser.New(serviceuser.New())

	r.Get("user/:id", userHandler.GetUser)

	return &Server{
		http: &http.Server{
			Addr:    net.JoinHostPort(params.Config.HTTPEndPoint.Host, params.Config.HTTPEndPoint.Port),
			Handler: r,
		},
	}
}

func (s *Server) Start() error {
	op := er.GetOperator()

	if err := s.http.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Info().Msgf("Shutdown http, done!!")
		} else {
			panic(er.WrapOp(err, op))
		}
	}
	return nil
}

// Close When it get sigterm, It'll gracefully closed till request is done or TCP connection reset.
func (s *Server) Stop() {
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSEGV)
	<-stop
	log.Info().Msg("i got the SIGTERM signal, gotta stop")
	log.Info().Msg("Shutdown http, start!!")
	if err := s.http.Shutdown(context.Background()); err != nil {
		log.Err(err).Msgf("http shutdown")
	}
	log.Info().Msg("gracefully shutdown!")
	close(stop)
}
