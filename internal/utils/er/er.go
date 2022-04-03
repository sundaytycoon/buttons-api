package er

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

type Error struct {
	Ops       []string
	Message   string
	SourceErr error
	NamedErr  error
}

func WithMessage(err error, msg string) error {
	e := newError(err)
	e.Message = msg
	return e
}

func WithNamedErr(err error, named error) error {
	e := newError(err)
	e.NamedErr = named
	return e
}

func WithSourceErr(err error, source error) error {
	e := newError(err)
	e.SourceErr = source
	return e
}

func WrapOp(err error, op string) error {
	e := newError(err)
	e.Ops = append(e.Ops, op)
	return e
}

func newError(err error) *Error {
	if err == nil {
		err = errors.New("empty error")
	}
	if e, ok := err.(*Error); ok {
		return e
	}
	return &Error{SourceErr: err, NamedErr: err, Message: err.Error()}
}

func IsSource(err, sourceErr error) bool {
	if err != nil || sourceErr != nil {
		return false
	}
	se := newError(err)
	te := newError(sourceErr)

	return errors.Is(se.SourceErr, te.SourceErr)
}

func IsNamed(err, namedErr error) bool {
	if err == nil || namedErr == nil {
		return false
	}
	se := newError(err)
	te := newError(namedErr)

	return errors.Is(se.NamedErr, te.NamedErr)
}

func (e *Error) Error() string {
	if e.SourceErr == nil {
		e.SourceErr = errors.New("")
	}
	if e.NamedErr == nil {
		e.NamedErr = errors.New("")
	}

	var ops []string
	if os.Getenv("GO_PKG_ER_DEBUG") == "true" {
		ops = append(ops, e.NamedErr.Error())
		ops = append(ops, e.SourceErr.Error())
		ops = append(ops, e.Ops...)
	} else {
		ops = append(ops, e.NamedErr.Error())
	}
	ops = append(ops, e.Message)
	return strings.Join(ops, "\n")
}

func GetNamedErr(err error) error {
	e := newError(err)
	return e.NamedErr
}

func GetSourceErr(err error) error {
	e := newError(err)
	return e.SourceErr
}

func GetOperator() string {
	pc, _, _, _ := runtime.Caller(1)
	caller := runtime.FuncForPC(pc).Name()
	splits := strings.Split(caller, "/")
	return strings.Join(splits[3:], ".")
}

func PanicError(err error, msg ...string) {
	if err != nil {
		panic(fmt.Errorf("%w, %v", err, msg))
	}
}
