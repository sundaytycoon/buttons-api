package handlergrpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	buttonsapi "github.com/sundaytycoon/buttons-api"
	v1grpc "github.com/sundaytycoon/buttons-api/api/proto/v1"
)

type authService interface {
	GetWebOAuthRedirectURL(context.Context, string, string) (string, error)
	GetWebCallback(context.Context, string, string, string) (string, string, error)
}

func NewAuthHandler(authService authService) *authHandler {
	return &authHandler{
		authService: authService,
	}
}

type authHandler struct {
	authService authService
	v1grpc.UnimplementedAuthServiceServer
}

func (h *authHandler) GetWebRedirectURL(ctx context.Context, req *v1grpc.GetWebRedirectURLRequest) (*v1grpc.GetWebRedirectURLResponse, error) {
	res, err := h.authService.GetWebOAuthRedirectURL(ctx, req.Provider, req.Service)
	if err != nil {
		st, _ := status.New(codes.Internal, err.Error()).WithDetails(req)
		return nil, st.Err()
	}
	return &v1grpc.GetWebRedirectURLResponse{
		Provider:    req.Provider,
		RedirectUrl: res,
	}, nil
}

func (h *authHandler) GetWebGoogleCallback(ctx context.Context, req *v1grpc.GetWebGoogleCallbackRequest) (*v1grpc.GetWebGoogleCallbackResponse, error) {
	// provider, code, state
	host, tempToken, err := h.authService.GetWebCallback(ctx, buttonsapi.Google, req.Code, req.State)
	if err != nil {
		st, _ := status.New(codes.Internal, err.Error()).WithDetails(req)
		return nil, st.Err()
	}
	fmt.Println(host, tempToken)
	return &v1grpc.GetWebGoogleCallbackResponse{}, nil
}
