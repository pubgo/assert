package assert

import (
	"errors"
	"fmt"
	"reflect"
)

func isNil(err error) bool {
	switch _e := err.(type) {
	case *KErr:
		return _e.err == nil
	case error:
		return _e == nil
	default:
		panic("unknown type")
	}
	return true
}

func Type(err interface{}) {
	P(err)
	fmt.Println(reflect.TypeOf(err).String(), funcCaller())
	fmt.Println("******************************")
}

func _Try(fn func()) (err *KErr) {
	Bool(fn == nil, "the func is nil")

	_v := reflect.TypeOf(fn)
	Bool(_v.Kind() != reflect.Func, "the params type(%s) is not func", _v.String())

	defer func() {
		defer func() {
			if r := recover(); r != nil {
				switch d := r.(type) {
				case *KErr:
					err = d
				case error:
					err.SetErr(d)
				case string:
					err.SetErr(errors.New(d))
				}
			}
		}()
		fn()
	}()
	return nil
}
