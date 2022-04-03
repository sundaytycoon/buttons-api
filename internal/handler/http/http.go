package http

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	buttonsapi "github.com/sundaytycoon/buttons-api"
	v1oapi "github.com/sundaytycoon/buttons-api/api/oapi/v1"
	"github.com/sundaytycoon/buttons-api/internal/handler/http/helper"
	"github.com/sundaytycoon/buttons-api/internal/model"
	"github.com/sundaytycoon/buttons-api/internal/utils/er"
)

type authService interface {
	GetWebOAuthRedirectURL(context.Context, string, string) (string, error)
	GetWebCallback(context.Context, url.Values, string) (*model.User, error)
}

type handler struct {
	authService authService
}

// New returns a new handler.
func New(authService authService) *handler {
	return &handler{
		authService: authService,
	}
}

// GetAuthWebRedirectProvider get redirect url from proivder
// (GET /auth/web/redirect/{provider})
func (h *handler) GetAuthWebRedirectProvider(w http.ResponseWriter, r *http.Request, provider string) {
	redirectUrl, err := h.authService.GetWebOAuthRedirectURL(r.Context(), provider, "")
	if err != nil {
		_ = helper.RenderError(err, w, r)
		return
	}

	w.Header().Set(v1oapi.HeaderLocation, redirectUrl)
	w.WriteHeader(v1oapi.StatusFound)
}

// GetAuthWebCallbackProvider callback for provider(google, kakao)
// (GET /auth/web/callback/{provider})
func (h *handler) GetAuthWebCallbackProvider(w http.ResponseWriter, r *http.Request, provider string) {
	op := er.GetOperator()

	v, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		err = er.WrapOp(err, op)
		err = er.WithNamedErr(err, buttonsapi.ErrGoogleOAuthCallbackInternalError)
		_ = helper.RenderError(err, w, r)
		return
	}

	u, err := h.authService.GetWebCallback(r.Context(), v.Query(), provider)
	if err != nil {
		_ = helper.RenderError(err, w, r)
		return
	}

	fmt.Println(u)
}
