package assert

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const callDepth = 2

var _stacks []string

func _log(s string) {
	_stacks = append(_stacks, s)
}

func Bool(b bool, format string, args ...interface{}) {
	if b {
		_log(funcCaller() + fmt.Sprintf(format, args...))
		panic("")
	}
}

func Err(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	_log(funcCaller() + fmt.Sprintf(format, args...))
	panic(err)
}

func MustNotError(err error) {
	if err == nil {
		return
	}

	_log(funcCaller() + err.Error())
	panic(err)
}

func NotNil(err error) {
	if err == nil {
		return
	}
	panic(err)
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

func GetStacks() []string {
	return _stacks
}

func LogStacks() {
	fmt.Println(strings.Join(_stacks, "\n"))
}
