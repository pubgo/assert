package assert

import (
	"encoding/json"
	"fmt"
	"github.com/juju/errors"
	"reflect"
	"strings"
)

type Assert struct {
}

func (t *Assert) P(d ...interface{}) {
	for _, i := range d {
		dt, err := json.MarshalIndent(i, "", "\t")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
	}
}

func (t *Assert) Bool(b bool, format string, args ...interface{}) {
	if b {
		err := errors.NewErr(format, args...)
		err.SetLocation(1)
		panic(strings.Join(err.StackTrace(), "\n"))
	}
}

func (t *Assert) Err(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	e := errors.Annotatef(err, format, args...)
	e.(*errors.Err).SetLocation(1)
	panic(errors.ErrorStack(e))
}

func (t *Assert) MustNotError(err error) {
	if err != nil {
		e := errors.Trace(err)
		e.(*errors.Err).SetLocation(1)
		panic(errors.ErrorStack(e))
	}
}
