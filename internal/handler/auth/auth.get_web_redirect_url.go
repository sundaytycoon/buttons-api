package auth

//
//func (h *Handler) GetWebRedirectURL(ctx context.Context, provider, service string) (*GetWebRedirectURLOut, error) {
//	op := er.GetOperator()
//
//	if service == "" {
//		err := er.New("service is required", buttonsapi.ErrBadRequest)
//		return nil, er.WrapOp(err, op)
//	}
//	if provider == "" {
//		err := er.New("provider is required", buttonsapi.ErrBadRequest)
//		return nil, er.WrapOp(err, op)
//	}
//
//	var fromHost string
//	if buttonsapi.Service(service) == buttonsapi.ButtonsAdmin {
//		fromHost = "http://" + h.config.AdminWeb.DSN
//	} else {
//		err := er.New("service is not valid", buttonsapi.ErrBadRequest)
//		return nil, er.WrapOp(err, op)
//	}
//
//	redirectURL, err := h.authService.GetWebOAuthRedirectURL(in.Provider, fromHost)
//	if err != nil {
//		return nil, er.WrapOp(err, op)
//	}
//	out := &GetWebRedirectURLOut{
//		Provider:    in.Provider,
//		RedirectURL: redirectURL,
//	}
//	return out, nil
//}
