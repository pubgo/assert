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

func KTry(fn interface{}, args ...interface{}) error {
	m := kerrGet()
	defer kerrPut(m)

	if r := Try(fn, args...); r != nil {
		switch d := r.(type) {
		case *KErr:
			m = d
		case error:
			m.err = d
			m.msg = d.Error()
			m.caller = funcCaller(callDepth)
		case string:
			m.err = errors.New(d)
			m.msg = d
			m.caller = funcCaller(callDepth)
		default:
			m.msg = fmt.Sprintf("type error %v", d)
			m.err = errors.New(m.msg)
			m.caller = funcCaller(callDepth)
			m.tag = ErrTag.UnknownErr
		}
	}

	if m.err == nil {
		return nil
	}

	return m.copy()
}

func ErrHandle(err error, fn func(err *KErr)) {
	if IsNil(err) {
		return
	}

	fn(err.(*KErr))
}
