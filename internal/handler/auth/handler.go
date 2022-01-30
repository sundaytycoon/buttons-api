package auth

import (
	"context"
	"github.com/sundaytycoon/buttons-api/edge/google"
	adapterbatchdb "github.com/sundaytycoon/buttons-api/internal/adapter/batchdb"
	"github.com/sundaytycoon/buttons-api/internal/config"
	repositoryauth "github.com/sundaytycoon/buttons-api/internal/repository/auth"
	serviceauth "github.com/sundaytycoon/buttons-api/internal/service/auth"
	"go.uber.org/dig"
	"time"
)

type authService interface {
	GetWebOAuthRedirectURL(provider, fromHost string) (string, error)
	GetWebCallback(ctx context.Context, provider, code string) (string, string, error)
}

type Handler struct {
	config      *config.Config
	authService authService

	timeoutMillis time.Duration
}

func New(params struct {
	dig.In
	Config       *config.Config
	ServiceDB    *adapterbatchdb.Adapter
	GoogleClient *google.Client
}) *Handler {

	authRepository := repositoryauth.New(
		params.GoogleClient,
	)

	return &Handler{
		timeoutMillis: 5 * time.Second,
		authService:   serviceauth.New(authRepository),
	}
}
