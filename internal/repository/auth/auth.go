package auth

import (
	"context"

	buttonsapi "github.com/sundaytycoon/buttons-api"
	"github.com/sundaytycoon/buttons-api/edge/google"
	"github.com/sundaytycoon/buttons-api/internal/constants/model"
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

func (r *Repository) GetUserInfoFromProvider(ctx context.Context, provider, code string) (*model.UserToken, error) {
	op := er.GetOperator()

	var userToken *model.UserToken
	var err error
	if provider == buttonsapi.Google {
		t, err := r.googleClient.OAuthCallback(ctx, code)
		if err != nil {
			if er.Is(err, google.ErrEmailIsNotVerified) {
				err = er.WithNamedErr(err, buttonsapi.ErrGoogleOAuthCallbackEmailIsNotValid)
			} else {
				err = er.WithNamedErr(err, buttonsapi.ErrGoogleOAuthCallbackInternalError)
			}
			return nil, er.WrapOp(err, op)
		}

		userToken = &model.UserToken{
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
			TokenType:    t.TokenType,
			Expiry:       t.Expiry,

			Email:   t.Email,
			Picture: t.Picture,
		}
	} else {
		err = er.New("'provider' service is not defined", buttonsapi.ErrBadRequest)
		return nil, er.WrapOp(err, op)
	}

	return userToken, nil
}
