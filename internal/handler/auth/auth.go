package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"google.golang.org/grpc/metadata"
	"io/ioutil"
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

const UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:5002/api/v1/auth/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
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
		RedirectUrl: googleOauthConfig.AuthCodeURL("http://localhost:3000"),
	}
	log.Trace().Str("operator", "GetRedirectURL").Interface("req", req).Interface("res", res).Send()
	return res, nil
}

func (h *Handler) GetCallback(ctx context.Context, req *v1pb.GetCallbackRequest) (*v1pb.GetCallbackResponse, error) {
	res := &v1pb.GetCallbackResponse{}
	token, err := googleOauthConfig.Exchange(ctx, req.Code)
	if err != nil {
		fmt.Println("GetCallback = 1")
		return nil, err
	}
	fmt.Println(token)
	t, err := googleOauthConfig.TokenSource(ctx, token).Token()
	if err != nil {
		fmt.Println("GetCallback = 2")
		return nil, err
	}
	fmt.Println(t)
	client := googleOauthConfig.Client(ctx, token)
	userInfoResp, err := client.Get(UserInfoAPIEndpoint)
	if err != nil {
		fmt.Println("GetCallback = 3")
		return nil, err
	}
	defer userInfoResp.Body.Close()
	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		fmt.Println("GetCallback = 3")
		return nil, err
	}
	log.Trace().
		Str("user_info", string(userInfo)).
		Str("operator", "GetCallback").
		Interface("token", token).
		Interface("req", req).
		Interface("res", res).
		Send()

	// FIXME: https://stackoverflow.com/questions/49878855/how-to-do-a-302-redirect-in-grpc-gateway
	metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{
		"status":   "302",
		"Location": req.State,
	}))
	return res, nil
}
