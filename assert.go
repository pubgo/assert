package assert

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

const callDepth = 2

var MaxStack = 10000

func Bool(b bool, format string, args ...interface{}) {
	if b {
		_e := fmt.Sprintf(format, args...)

		_ke := NewKErr()
		_ke.AddStack(funcCaller() + _e)
		_ke.SetErr(errors.New(_e))
		panic(_ke)
	}
}

func Err(err error, format string, args ...interface{}) {
	if !reflect.ValueOf(err).IsValid() || reflect.ValueOf(err).IsNil() {
		return
	}

	_ke := NewKErr()
	_ke.AddStack(funcCaller() + fmt.Sprintf(format, args...))
	_ke.SetErr(err)
	panic(_ke)
}

func MustNotError(err error) {
	if !reflect.ValueOf(err).IsValid() || reflect.ValueOf(err).IsNil() {
		return
	}

	_ke := NewKErr()
	_ke.AddStack(funcCaller() + err.Error())
	_ke.SetErr(err)
	panic(_ke)
}

func NotNil(err error) {
	if !reflect.ValueOf(err).IsValid() || reflect.ValueOf(err).IsNil() {
		return
	}

	_ke := NewKErr()
	_ke.SetErr(err)
	panic(_ke)
}

func P(d ...interface{}) {
	for _, i := range d {
		dt, err := json.MarshalIndent(i, "", "\t")
		MustNotError(err)
		fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
	}
}

var True = Bool

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
					err = &KErr{err: d}
				case string:
					err = &KErr{err: errors.New(d)}
				}
			}
		}()
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return
}
