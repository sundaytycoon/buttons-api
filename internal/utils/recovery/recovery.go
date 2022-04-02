package recovery

import (
	"fmt"

	"github.com/pkg/errors"
)

// RecoverFn defines recover handler function.
func RecoverFn(p interface{}) error {
	var err error
	switch t := p.(type) {
	case error:
		err = t
	case string:
		err = errors.New(t)
	case fmt.Stringer:
		err = errors.New(t.String())
	default:
		err = fmt.Errorf("%v", t)
	}
	return err
}
