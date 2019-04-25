package assert

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

const callDepth = 2

func Bool(b bool, format string, args ...interface{}) {
	if b {
		panic(NewKErr(funcCaller()+fmt.Sprintf(format, args...), nil))
	}
}

func Err(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}
	panic(NewKErr(funcCaller()+fmt.Sprintf(format, args...), err))
}

func MustNotError(err error) {
	if err == nil {
		return
	}
	panic(NewKErr(funcCaller(), err))
}

func NotNil(err error) {
	if err == nil {
		return
	}
	panic(NewKErr("", err))
}

func P(d ...interface{}) {
	for _, i := range d {
		dt, err := json.MarshalIndent(i, "", "\t")
		MustNotError(err)
		fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
	}
}

var True = Bool

func _Try(fn func()) (err error) {
	True(fn == nil, "the func is nil")

	_v := reflect.TypeOf(fn)
	True(_v.Kind() != reflect.Func, "the params type(%s) is not func", _v.String())

	defer func() {
		defer func() {
			if r := recover(); r != nil {
				switch d := r.(type) {
				case error:
					err = d
				case string:
					err = errors.New(d)
				}
			}
		}()
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return
}
