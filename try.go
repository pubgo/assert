package assert

import (
	"errors"
	"fmt"
)

func Try(fn interface{}, args ...interface{}) (r interface{}) {
	defer func() {
		defer func() {
			r = recover()
		}()
		FnOf(fn, args...)()
	}()

	assertFn(fn)
	return
}

func KTry(fn interface{}, args ...interface{}) (err error) {
	m := &KErr{}
	if r := Try(fn, args...); r != nil {
		switch d := r.(type) {
		case *KErr:
			m = d
		case error:
			m.Err = d
			m.Msg = d.Error()
		case string:
			m.Err = errors.New(d)
			m.Msg = d
		default:
			panic(fmt.Sprintf("type error %v", d))
		}
	}

	if m.Err == nil {
		err = nil
	} else {
		err = m
	}
	return
}
