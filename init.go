package assert

import (
	"errors"
	"fmt"
	"reflect"
)

func assertFn(fn interface{}) {
	ST(IsNil(fn), "the func is nil")

	_v := reflect.TypeOf(fn)
	ST(_v.Kind() != reflect.Func, "the params(%s) is not func type", _v.String())
}

func Try(fn func()) (r interface{}) {
	defer func() {
		defer func() {
			r = recover()
		}()
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	assertFn(fn)
	return
}

func KTry(fn func()) (err error) {
	m := &KErr{}
	if r := Try(fn); r != nil {
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
