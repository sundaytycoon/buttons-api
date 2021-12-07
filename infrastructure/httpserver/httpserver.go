package httpserver

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog/log"

	"github.com/sundaytycoon/profile.me-server/internal/config"
	"github.com/sundaytycoon/profile.me-server/pkg/er"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	http *http.Server
	mux  *chi.Mux
}

func New(cfg *config.Config) *Server {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	return &Server{
		http: &http.Server{
			Addr:    net.JoinHostPort(cfg.HTTPEndPoint.Host, cfg.HTTPEndPoint.Port),
			Handler: r,
		},
		mux: r,
	}
}

type Handler interface {
	RouteHTTP(r *chi.Mux)
}

func (s *Server) SetHandler(h Handler) {
	h.RouteHTTP(s.mux)
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
