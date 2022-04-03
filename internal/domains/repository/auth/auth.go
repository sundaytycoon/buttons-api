package auth

import (
	"context"

	"github.com/pkg/errors"

	"github.com/sundaytycoon/buttons-api/internal/utils/er"

	buttonsapi "github.com/sundaytycoon/buttons-api"
	"github.com/sundaytycoon/buttons-api/edge/google"
	"github.com/sundaytycoon/buttons-api/internal/model"
)

type authStorage interface {
}

type googleClient interface {
	OAuthRedirectURL(string) string
	OAuthCallback(context.Context, string) (*google.OAuthCallbackResponse, error)
}

// repository  직접 구현 한 곳
type repository struct {
	googleClient googleClient
	authStorage  authStorage
}

func New(googleClient googleClient, authStorage authStorage) *repository {
	return &repository{
		googleClient: googleClient,
		authStorage:  authStorage,
	}
}

func (r *repository) GetOAuthRedirectURL(provider, fromHost string) (string, error) {
	if provider == buttonsapi.Google {
		return r.googleClient.OAuthRedirectURL(fromHost), nil
	} else {
		return "", er.WithNamedErr(errors.New("'provider' service is not defined"), buttonsapi.ErrGoogleOAuthCallbackInternalError)
	}
}

func (r *repository) GetUserInfoFromProvider(ctx context.Context, provider, code string) (*model.UserToken, error) {
	op := er.GetOperator()

	var userToken *model.UserToken
	var err error
	if provider == buttonsapi.Google {
		t, err := r.googleClient.OAuthCallback(ctx, code)
		if err != nil {
			if er.IsSource(err, google.ErrEmailIsNotVerified) {
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
		err = er.WithNamedErr(errors.New("'provider' service is not defined"), buttonsapi.ErrInvalidRequest)
		return nil, er.WrapOp(err, op)
	}

	return userToken, nil
}
