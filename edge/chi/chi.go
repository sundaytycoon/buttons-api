package chi

import (
	"net/http"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	corsmiddleware "github.com/go-chi/cors"

	buttonsapi "github.com/sundaytycoon/buttons-api"
)

func New() chi.Router {
	r := chi.NewRouter()
	r.Use(chimiddleware.RealIP)
	r.Use(chimiddleware.Logger)
	r.Use(func(handler http.Handler) http.Handler {
		return xray.Handler(xray.NewFixedSegmentNamer(string(buttonsapi.ButtonsAPI)), handler)
	})
	r.Use(chimiddleware.Recoverer)
	r.Use(corsmiddleware.New(corsmiddleware.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           30,
	}).Handler)
	r.Use(chimiddleware.RequestID)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("OK"))
	})

	return r
}
