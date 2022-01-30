package google

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	buttonsapi "github.com/sundaytycoon/buttons-api"
	"github.com/sundaytycoon/buttons-api/internal/config"
	"github.com/sundaytycoon/buttons-api/pkg/er"
	"go.uber.org/dig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"time"
)

const UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"

type Client struct {
	cfg oauth2.Config
}

func New(params struct {
	dig.In
	Config *config.Config
}) *Client {
	return newClient(params.Config.Google)
}

func newClient(googleCfg *config.Google) *Client {
	return &Client{
		cfg: oauth2.Config{
			RedirectURL:  googleCfg.OAuthCallbackURL,
			ClientID:     googleCfg.ClientID,
			ClientSecret: googleCfg.ClientSecret,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (c *Client) OAuthRedirectURL(state string) string {
	return c.cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

type OAuthCallbackResponse struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	Expiry       time.Time

	Email   string
	Picture string
}

func (c *Client) OAuthCallback(ctx context.Context, code string) (*OAuthCallbackResponse, error) {
	op := er.GetOperator()

	token, err := c.cfg.Exchange(ctx, code)
	if err != nil {
		err = er.WithNamedErr(err, buttonsapi.ErrGoogleOAuthCallbackInternalError)
		return nil, er.WrapOp(err, op)
	}
	t, err := c.cfg.TokenSource(ctx, token).Token()
	if err != nil {
		err = er.WithNamedErr(err, buttonsapi.ErrGoogleOAuthCallbackInternalError)
		return nil, er.WrapOp(err, op)
	}
	if t.RefreshToken == "" {
		err = er.New(
			"google email is not verified",
			buttonsapi.ErrGoogleOAuthCallbackInternalError,
		)
		return nil, er.WrapOp(err, op)
	}

	client := c.cfg.Client(ctx, t)
	userInfoResp, err := client.Get(UserInfoAPIEndpoint)
	if err != nil {
		err = er.WithNamedErr(err, buttonsapi.ErrGoogleOAuthCallbackInternalError)
		return nil, er.WrapOp(err, op)
	}
	defer userInfoResp.Body.Close()

	type userInfo struct {
		Email         string `json:"email"`
		Picture       string `json:"picture"`
		EmailVerified bool   `json:"email_verified"`
	}
	u := userInfo{}
	err = jsoniter.NewDecoder(userInfoResp.Body).Decode(&u)
	if err != nil {
		err = er.WithNamedErr(err, buttonsapi.ErrGoogleOAuthCallbackInternalError)
		return nil, er.WrapOp(err, op)
	}
	if !u.EmailVerified {
		err = er.New("google email is not verified", buttonsapi.ErrGoogleOAuthCallbackEmailIsNotValid)
		return nil, er.WrapOp(err, op)
	}

	return &OAuthCallbackResponse{
		AccessToken:  t.AccessToken,
		RefreshToken: t.RefreshToken,
		TokenType:    t.TokenType,
		Expiry:       t.Expiry,
		Email:        u.Email,
		Picture:      u.Picture,
	}, nil
}
