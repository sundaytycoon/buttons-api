package auth

import (
	"context"

	buttonsapi "github.com/sundaytycoon/buttons-api"

	"github.com/sundaytycoon/buttons-api/pkg/er"
)

type GetWebRedirectURLIn struct {
	Provider string ``
	Service  string
}

type GetWebRedirectURLOut struct {
	Provider    string
	RedirectURL string
}

func (h *Handler) GetWebRedirectURL(ctx context.Context, in *GetWebRedirectURLIn) (*GetWebRedirectURLOut, error) {
	op := er.GetOperator()

	if in.Service == "" {
		err := er.New("service is required", buttonsapi.ErrBadRequest)
		return nil, er.WrapOp(err, op)
	}
	if in.Provider == "" {
		err := er.New("provider is required", buttonsapi.ErrBadRequest)
		return nil, er.WrapOp(err, op)
	}

	var fromHost string
	if buttonsapi.Service(in.Service) == buttonsapi.ButtonsAdmin {
		fromHost = h.config.ButtonsAdminWeb.HTTPAddress()
	} else {
		err := er.New("service is not valid", buttonsapi.ErrBadRequest)
		return nil, er.WrapOp(err, op)
	}

	redirectURL, err := h.authService.GetWebOAuthRedirectURL(in.Provider, fromHost)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	out := &GetWebRedirectURLOut{
		Provider:    in.Provider,
		RedirectURL: redirectURL,
	}
	return out, nil
}
