package assert

import (
	"encoding/json"
	"github.com/juju/errors"
	"reflect"
	"strings"
)

const callDepth = 2

type Assert struct {
}

func (t *Assert) P(d ...interface{}) {
	for _, i := range d {
		dt, err := json.MarshalIndent(i, "", "\t")
		if err != nil {
			panic(err.Error())
		}
		println(reflect.ValueOf(i).String(), "->", string(dt))
	}
}

func (t *Assert) Bool(b bool, format string, args ...interface{}) {
	if b {
		err := errors.NewErr(format, args...)
		err.SetLocation(callDepth)
		panic(strings.Join(err.StackTrace(), "\n"))
	}
}

func (t *Assert) Err(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	e := errors.Annotatef(err, format, args...)
	e.(*errors.Err).SetLocation(callDepth)
	panic(errors.ErrorStack(e))
}

func (t *Assert) MustNotError(err error) {
	if err != nil {
		e := errors.Trace(err)
		e.(*errors.Err).SetLocation(callDepth)
		panic(errors.ErrorStack(e))
	}
}
