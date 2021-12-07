package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/dig"

	adapterservicedb "github.com/sundaytycoon/profile.me-server/internal/adapter/servicedb"
	repositoryuser "github.com/sundaytycoon/profile.me-server/internal/repository/user"
	serviceuser "github.com/sundaytycoon/profile.me-server/internal/service/user"
	servicedbstorage "github.com/sundaytycoon/profile.me-server/internal/storage/servicedb"
)

type Handler struct {
	userService userService
}

func New(params struct {
	dig.In
	ServiceDB *adapterservicedb.Adapter
}) *Handler {

	repositoryUser := repositoryuser.New(params.ServiceDB, servicedbstorage.New())
	serviceUser := serviceuser.New(repositoryUser)

	return &Handler{
		userService: serviceUser,
	}
}

func (h *Handler) RouteHTTP(r *chi.Mux) {
	r.Get("/user/:id", h.GetUser)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(400)
		return
	}

	u, err := h.userService.Get(r.Context(), id)
	if err != nil {
		w.WriteHeader(502)
		render.Respond(w, r, map[string]interface{}{
			"message": "occurred an error at api server",
		})
		return
	}
	w.WriteHeader(200)
	render.Respond(w, r, u)
}
