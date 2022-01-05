package httpserver

//
//import (
//	"net"
//	"net/http"
//
//	"github.com/rs/zerolog/log"
//
//	"github.com/sundaytycoon/buttons-api/internal/config"
//	"github.com/sundaytycoon/buttons-api/pkg/er"
//
//	"github.com/go-chi/chi/v5"
//	"github.com/go-chi/chi/v5/middleware"
//)
//
//type Server struct {
//	http *http.Server
//	mux  *chi.Mux
//}
//
//func New(cfg *config.Config) *Server {
//
//	r := chi.NewRouter()
//
//	// A good base middleware stack
//	r.Use(middleware.RequestID)
//	r.Use(middleware.RealIP)
//	r.Use(middleware.Logger)
//	r.Use(middleware.Recoverer)
//
//	return &Server{
//		http: &http.Server{
//			Addr:    net.JoinHostPort(cfg.HTTPEndPoint.Host, cfg.HTTPEndPoint.Port),
//			Handler: r,
//		},
//		mux: r,
//	}
//}
//
//type Handler interface {
//	RouteHTTP(r *chi.Mux)
//}
//
//func (s *Server) SetHandler(h Handler) {
//	h.RouteHTTP(s.mux)
//}
//
//func (s *Server) Start() error {
//	op := er.GetOperator()
//
//	if err := s.http.ListenAndServe(); err != nil {
//		if err == http.ErrServerClosed {
//			log.Info().Msgf("Shutdown http, done!!")
//		} else {
//			panic(er.WrapOp(err, op))
//		}
//	}
//	return nil
//}
