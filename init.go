package assert

import (
	"fmt"
	"reflect"
)

func Type(err interface{}) {
	P(err)
	fmt.Println(reflect.TypeOf(err).String(), funcCaller())
	fmt.Println("******************************")
}

func _Try(fn func()) (err error) {
	Bool(fn == nil, "the func is nil")

	_v := reflect.TypeOf(fn)
	Bool(_v.Kind() != reflect.Func, "the params type(%s) is not func", _v.String())

	defer func() {
		defer func() {
			m := &KErr{}
			if r := recover(); r != nil {
				switch d := r.(type) {
				case *KErr:
					m = d
				case error:
					m.Err = d
				default:
					panic("type error, must be *KErr type")
				}
			}

			if m.Err == nil {
				err = nil
			} else {
				err = m
			}
		}()
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return
}
