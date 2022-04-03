package buttonsapi

import (
	"errors"

	oapiv1 "github.com/sundaytycoon/buttons-api/api/oapi/v1"
	"github.com/sundaytycoon/buttons-api/internal/storage/servicedb/ent"
	"github.com/sundaytycoon/buttons-api/internal/utils/er"
)

var ()

var (
	// client side error
	ErrInvalidRequest = errors.New("error, invalid request")
	ErrRequestTimeout = errors.New("error, request timeout")

	// server side error
	ErrInternalServer                     = errors.New("error, internal server error")
	ErrGoogleOAuthCallbackInternalError   = errors.New("google authentication has an error")
	ErrGoogleOAuthCallbackEmailIsNotValid = errors.New("requested user's email is not valid")

	ErrRefreshTokenIsNotValid = errors.New("google oauth2: refresh_token is not loaded")
	ErrEmailIsNotVerified     = errors.New("google oauth2: email is not verified")
)

var httpErrToStatus = map[error]int{
	ErrInvalidRequest:                     oapiv1.StatusBadRequest,
	ErrRequestTimeout:                     oapiv1.StatusRequestTimeout,
	ErrInternalServer:                     oapiv1.StatusInternalServerError,
	ErrGoogleOAuthCallbackInternalError:   oapiv1.StatusInternalServerError,
	ErrGoogleOAuthCallbackEmailIsNotValid: oapiv1.StatusInternalServerError,
	ErrRefreshTokenIsNotValid:             oapiv1.StatusInternalServerError,
	ErrEmailIsNotVerified:                 oapiv1.StatusInternalServerError,
}

func HTTPErrorToStatusCode(err error) int {
	if err == nil {
		return oapiv1.StatusInternalServerError
	}
	if v, ok := httpErrToStatus[err]; ok {
		return v
	}

	sourceErr := er.GetSourceErr(err)
	if ent.IsValidationError(sourceErr) {
		return oapiv1.StatusInternalServerError
	} else if ent.IsConstraintError(sourceErr) {
		return oapiv1.StatusInternalServerError
	} else if ent.IsNotFound(sourceErr) {
		return oapiv1.StatusInternalServerError
	} else if ent.IsNotLoaded(sourceErr) {
		return oapiv1.StatusInternalServerError
	} else if ent.IsNotSingular(sourceErr) {
		return oapiv1.StatusInternalServerError
	} else if oapiv1.IsErrorTypeOfOAPI(sourceErr) {
		return oapiv1.StatusBadRequest
	}

	return oapiv1.StatusInternalServerError
}
