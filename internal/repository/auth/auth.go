package auth

import (
	"context"
	buttonsapi "github.com/sundaytycoon/buttons-api"
	"github.com/sundaytycoon/buttons-api/edge/google"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

// Repository  직접 구현 한 곳
type Repository struct {
	googleClient googleClient
}

type googleClient interface {
	OAuthRedirectURL(state string) string
	OAuthCallback(ctx context.Context, code string) (*google.OAuthCallbackResponse, error)
}

func New(gc googleClient) *Repository {
	return &Repository{
		googleClient: gc,
	}
}

func (r *Repository) GetOAuthRedirectURL(provider, fromHost string) (string, error) {
	if provider == buttonsapi.Google {
		return r.googleClient.OAuthRedirectURL(fromHost), nil
	} else {
		return "", er.New("'provider' service is not defined", buttonsapi.ErrGoogleOAuthCallbackInternalError)
	}
}
