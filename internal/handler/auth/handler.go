package auth

import (
	"context"
	"time"

	"go.uber.org/dig"

	"github.com/sundaytycoon/buttons-api/edge/google"
	"github.com/sundaytycoon/buttons-api/internal/config"
)

type authService interface {
	GetWebOAuthRedirectURL(provider, fromHost string) (string, error)
	GetWebCallback(ctx context.Context, provider, code, state string) (string, string, error)
}

type Handler struct {
	config      *config.Config
	authService authService

	timeoutMillis time.Duration
}

func New(
	params struct {
		dig.In
		Config       *config.Config
		GoogleClient *google.Client
	},
) *Handler {

	return &Handler{
		config: params.Config,

		timeoutMillis: 5 * time.Second,
	}
}
