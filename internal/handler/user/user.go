package user

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"

	v1pb "github.com/sundaytycoon/buttons-api/gen/go/proto/rpc/v1"
	adapterservicedb "github.com/sundaytycoon/buttons-api/internal/adapter/servicedb"
	repositoryuser "github.com/sundaytycoon/buttons-api/internal/repository/user"
	serviceuser "github.com/sundaytycoon/buttons-api/internal/service/user"
	servicedbstorage "github.com/sundaytycoon/buttons-api/internal/storage/servicedb"
)

type Handler struct {
	userService userService
	v1pb.UnimplementedUserServiceServer
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

func (h *Handler) Get(ctx context.Context, req *v1pb.UserMessage) (*v1pb.UserMessage, error) {
	log.Trace().Str("name", req.Name).Str("id", req.Id).Msg("message")
	return req, nil
}
