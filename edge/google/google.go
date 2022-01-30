package google

import (
	"context"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/sundaytycoon/buttons-api/internal/config"
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

var (
	ErrRefreshTokenIsNotValid = errors.New("google oauth2: refresh_token is not loaded")
	ErrEmailIsNotVerified     = errors.New("google oauth2: email is not verified")
)

func (c *Client) OAuthCallback(ctx context.Context, code string) (*OAuthCallbackResponse, error) {
	token, err := c.cfg.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	t, err := c.cfg.TokenSource(ctx, token).Token()
	if err != nil {
		return nil, err
	}
	if t.RefreshToken == "" {
		return nil, ErrRefreshTokenIsNotValid
	}

	client := c.cfg.Client(ctx, t)
	userInfoResp, err := client.Get(UserInfoAPIEndpoint)
	if err != nil {
		return nil, err
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
		return nil, err
	}
	if !u.EmailVerified {
		return nil, ErrEmailIsNotVerified
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
