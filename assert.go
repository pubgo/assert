package assert

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
		_ke.Panic()
	}
}

func Err(err error, format string, args ...interface{}) {
	if !reflect.ValueOf(err).IsValid() || reflect.ValueOf(err).IsNil() {
		return
	}

	if isNil(err) {
		return
	}

	_ke := NewKErr()
	_ke.AddStack(funcCaller() + fmt.Sprintf(format, args...))
	_ke.SetErr(err)
	_ke.Panic()
}

func MustNotError(err error) {
	if !reflect.ValueOf(err).IsValid() || reflect.ValueOf(err).IsNil() {
		return
	}

	if isNil(err) {
		return
	}

	_ke := NewKErr()
	_ke.AddStack(funcCaller() + err.Error())
	_ke.SetErr(err)
	_ke.Panic()
}

func NotNil(err error) {
	if !reflect.ValueOf(err).IsValid() || reflect.ValueOf(err).IsNil() {
		return
	}

	if isNil(err) {
		return
	}

	_ke := NewKErr()
	_ke.SetErr(err)
	_ke.Panic()
}

func P(d ...interface{}) {
	for _, i := range d {
		if dt, err := json.MarshalIndent(i, "", "\t"); err != nil {
			panic(err)
		} else {
			log.Println(reflect.ValueOf(i).String(), "->", string(dt))
		}

	}
}

var True = Bool
