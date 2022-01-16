package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"os"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/grpc"

	v1pb "github.com/sundaytycoon/buttons-api/gen/go/buttons/api/v1"
	adapterbatchdb "github.com/sundaytycoon/buttons-api/internal/adapter/batchdb"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:5002/auth/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

type Handler struct {
	authService sessionService
	v1pb.UnimplementedAuthServiceServer

	timeoutMillis time.Duration
}

func New(params struct {
	dig.In
	ServiceDB *adapterbatchdb.Adapter
}) *Handler {
	return &Handler{
		timeoutMillis: time.Second * 5,
	}
}

func (h *Handler) Name() string {
	return "AuthHandler"
}

func (h *Handler) Close() error {
	return nil
}

func (h *Handler) Register(grpcServer grpc.ServiceRegistrar) {
	v1pb.RegisterAuthServiceServer(grpcServer, h)
}

func (h *Handler) Connect(grpcEndpoint string, mux *runtime.ServeMux) error {
	op := er.GetOperator()

	ctx, cancel := context.WithTimeout(context.Background(), h.timeoutMillis)
	defer cancel()
	conn, err := grpc.DialContext(
		ctx,
		grpcEndpoint,
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		return er.WrapOp(err, op)
	}

	return v1pb.RegisterAuthServiceHandler(ctx, mux, conn)
}

func (h *Handler) GetRedirectURL(ctx context.Context, req *v1pb.GetRedirectURLRequest) (*v1pb.GetRedirectURLResponse, error) {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	//cookie := &http.Cookie{Name: "oauthstate", Value: state, Expires: time.Now().Add(1 * 24 * time.Hour)}

	res := &v1pb.GetRedirectURLResponse{
		Provider:    req.Provider,
		RedirectUrl: googleOauthConfig.AuthCodeURL(state),
	}
	log.Trace().Interface("req", req).Interface("res", res).Send()
	return res, nil
}
