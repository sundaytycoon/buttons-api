package v1

import (
	"github.com/pkg/errors"
)

type errorChecker func(error) bool

func IsErrorTypeOfOAPI(err error) bool {
	if err == nil {
		return false
	}

	var oapiErrors = []errorChecker{
		IsRequiredParamError,
		IsRequiredHeaderError,
		IsUnmarshalingParamError,
		IsUnescapedCookieParamError,
		IsInvalidParamFormatError,
		IsTooManyValuesForParamError,
	}
	for _, fn := range oapiErrors {
		if fn(err) {
			return true
		}
	}
	return false
}

func IsRequiredParamError(err error) bool {
	if err == nil {
		return false
	}
	var e *RequiredParamError
	return errors.As(err, &e)
}
func IsRequiredHeaderError(err error) bool {
	if err == nil {
		return false
	}
	var e *RequiredHeaderError
	return errors.As(err, &e)
}
func IsUnmarshalingParamError(err error) bool {
	if err == nil {
		return false
	}
	var e *UnmarshalingParamError
	return errors.As(err, &e)
}
func IsUnescapedCookieParamError(err error) bool {
	if err == nil {
		return false
	}
	var e *UnescapedCookieParamError
	return errors.As(err, &e)
}
func IsInvalidParamFormatError(err error) bool {
	if err == nil {
		return false
	}
	var e *InvalidParamFormatError
	return errors.As(err, &e)
}
func IsTooManyValuesForParamError(err error) bool {
	if err == nil {
		return false
	}
	var e *TooManyValuesForParamError
	return errors.As(err, &e)
}
