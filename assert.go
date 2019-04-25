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
		_e := fmt.Sprintf(format, args...)
		panic(&KErr{_stacks: []string{funcCaller() + _e}, err: errors.New(_e)})
	}
}

func Err(err error, format string, args ...interface{}) {
	if reflect.ValueOf(err).IsNil() {
		return
	}
	_s := funcCaller() + fmt.Sprintf(format, args...)
	_ke := &KErr{_stacks: []string{_s}}
	switch e := err.(type) {
	case *KErr:
		_ke.err = e.err
		_ke._stacks = append(_ke._stacks, e._stacks...)
	case error:
		_ke.err = e
	}

	panic(_ke)
}

func MustNotError(err error) {
	if reflect.ValueOf(err).IsNil() {
		return
	}

	_s := funcCaller()
	_ke := &KErr{_stacks: []string{_s}}
	switch e := err.(type) {
	case *KErr:
		_ke.err = e.err
		_ke._stacks = append(_ke._stacks, e._stacks...)
	case error:
		_ke.err = e
	}
	panic(_ke)
}

func NotNil(err error) {
	if reflect.ValueOf(err).IsNil() {
		return
	}

	_ke := &KErr{}
	switch e := err.(type) {
	case *KErr:
		_ke.err = e.err
		_ke._stacks = append(_ke._stacks, e._stacks...)
	case error:
		_ke.err = e
	}
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
				case error:
					err = &KErr{err: d}
				case string:
					err = &KErr{err: errors.New(d)}
				case *KErr:
					err = d
				}
			}
		}()
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}()
	return
}
