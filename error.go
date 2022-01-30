package buttonsapi

import "errors"

var (
	ErrBadRequest                         = errors.New("bad request, request value is not valid")
	ErrGoogleOAuthCallbackInternalError   = errors.New("google authentication has an error")
	ErrGoogleOAuthCallbackEmailIsNotValid = errors.New("requested user's email is not valid")
)
