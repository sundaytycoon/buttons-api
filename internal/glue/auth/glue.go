package auth

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	v1pb "github.com/sundaytycoon/buttons-api/gen/go/buttons/api/v1"
	handlerauth "github.com/sundaytycoon/buttons-api/internal/handler/auth"
	"github.com/sundaytycoon/buttons-api/pkg/er"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/url"
	"time"
)

type Glue struct {
	v1pb.UnimplementedAuthServiceServer
	Handler       *handlerauth.Handler
	timeoutMillis time.Duration
}

func New(handler *handlerauth.Handler) *Glue {
	return &Glue{
		timeoutMillis: time.Second * 5,
		Handler:       handler,
	}
}

func (h *Glue) Name() string {
	return "AuthHandler"
}

func (h *Glue) Close() error {
	return nil
}

func (h *Glue) Register(grpcServer grpc.ServiceRegistrar) {
	v1pb.RegisterAuthServiceServer(grpcServer, h)
}

func (h *Glue) Connect(grpcEndpoint string, mux *runtime.ServeMux) error {
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

func (g *Glue) GetWebRedirectURL(ctx context.Context, req *v1pb.GetWebRedirectURLRequest) (*v1pb.GetWebRedirectURLResponse, error) {
	out, err := g.Handler.GetWebRedirectURL(ctx, &handlerauth.GetWebRedirectURLIn{
		Provider: req.Provider,
		Service:  req.Service,
	})
	if err != nil {
		return nil, err
	}
	return &v1pb.GetWebRedirectURLResponse{
		Provider:    out.Provider,
		RedirectUrl: out.RedirectURL,
	}, nil
}

func (g *Glue) GetWebGoogleCallback(ctx context.Context, req *v1pb.GetWebGoogleCallbackRequest) (*v1pb.GetWebGoogleCallbackResponse, error) {
	out, err := g.Handler.GetWebCallback(ctx, &handlerauth.GetWebCallbackIn{})

	qs := url.Values{}
	qs.Add("provider", out.Provider)
	qs.Add("temporary_token", out.TemporaryToken)

	err = grpc.SendHeader(ctx, metadata.Pairs(
		"Status", "302",
		"Location", fmt.Sprintf("%s?%s", out.ToHost, qs.Encode()),
	))
	if err != nil {
		return nil, err
	}

	return &v1pb.GetWebGoogleCallbackResponse{}, nil
}
