package user

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
	"google.golang.org/grpc"

	v1pb "github.com/sundaytycoon/buttons-api/gen/go/buttons/api/v1"

	adapterservicedb "github.com/sundaytycoon/buttons-api/internal/adapter/servicedb"
	repositoryuser "github.com/sundaytycoon/buttons-api/internal/repository/user"
	serviceuser "github.com/sundaytycoon/buttons-api/internal/service/user"
	servicedbstorage "github.com/sundaytycoon/buttons-api/internal/storage/servicedb"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

type Handler struct {
	userService   userService
	timeoutMillis time.Duration
	v1pb.UnimplementedUserServiceServer
}

func New(params struct {
	dig.In
	ServiceDB *adapterservicedb.Adapter
}) *Handler {

	repositoryUser := repositoryuser.New(params.ServiceDB, servicedbstorage.New())
	serviceUser := serviceuser.New(repositoryUser)

	return &Handler{
		timeoutMillis: 2000 * time.Millisecond,
		userService:   serviceUser,
	}
}

func (h *Handler) Name() string {
	return "UserHandler"
}

func (h *Handler) Close() error {
	return nil
}

func (h *Handler) Register(grpcServer grpc.ServiceRegistrar) {
	v1pb.RegisterUserServiceServer(grpcServer, h)
}

func (h *Handler) Connect(grpcEndpoint string, mux *runtime.ServeMux) error {
	op := er.GetOperator()

	ctx, cancel := context.WithTimeout(context.Background(), h.timeoutMillis)
	defer cancel()
	conn, err := grpc.DialContext(ctx, grpcEndpoint, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return er.WrapOp(err, op)
	}

	return v1pb.RegisterUserServiceHandler(ctx, mux, conn)
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

func (h *Handler) Get(ctx context.Context, req *v1pb.UserServiceGetRequest) (*v1pb.UserServiceGetResponse, error) {
	log.Trace().Str("name", req.Name).Str("id", req.Id).Msg("message")
	return &v1pb.UserServiceGetResponse{
		Id:   req.GetId(),
		Name: req.GetName(),
	}, nil
}
