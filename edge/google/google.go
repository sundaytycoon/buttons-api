package google

import (
	"context"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	buttonsapi "github.com/sundaytycoon/buttons-api"
	"github.com/sundaytycoon/buttons-api/internal/config"
)

const UserInfoAPIEndpoint = "https://www.googleapis.com/oauth2/v3/userinfo"

type Client struct {
	cfg oauth2.Config
}

func New(cfg *config.Config) *Client {
	protocol := "http"
	if cfg.Env != buttonsapi.ValueEnvLocal {
		protocol = "https"
	}

	return &Client{
		cfg: oauth2.Config{
			RedirectURL:  fmt.Sprintf("%s://%s%s", protocol, cfg.ApplicationHTTP.InternalDSN, cfg.Google.OAuthCallbackURL),
			ClientID:     cfg.Google.ClientID,
			ClientSecret: cfg.Google.ClientSecret,
			Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
			Endpoint:     google.Endpoint,
		},
	}
}

type OAuthCallbackResponse struct {
	AccessToken  string
	RefreshToken string
	TokenType    string
	Expiry       time.Time

	Email   string
	Picture string
}

func (c *Client) OAuthRedirectURL(state string) string {
	return c.cfg.AuthCodeURL(
		state,
		oauth2.AccessTypeOffline, // For get refresh_token
		oauth2.ApprovalForce,     // For get refresh_token
	)
}

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
		return nil, buttonsapi.ErrRefreshTokenIsNotValid
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
		return nil, buttonsapi.ErrEmailIsNotVerified
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
