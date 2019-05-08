package assert

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

const callDepth = 2

type M struct {
	msg string
	tag string
	M   map[string]interface{}
}

func (t *M) Msg(format string, args ...interface{}) {
	t.msg = fmt.Sprintf(format, args...)
}

func (t *M) Tag(tag string) {
	t.tag = tag
}

func ST(b bool, msg string, args ...interface{}) {
	if b {
		_m := fmt.Sprintf(msg, args...)
		panic(&KErr{
			Caller: funcCaller(),
			Msg:    _m,
			Err:    errors.New(_m),
		})
	}
}

func T(b bool, fn func(m *M)) {
	if b {
		_m := &M{M: make(map[string]interface{})}
		fn(_m)

		if len(_m.M) == 0 {
			_m.M = nil
		}

		panic(&KErr{
			Caller: funcCaller(),
			Msg:    _m.msg,
			Err:    errors.New(_m.msg),
			Tag:    _m.tag,
			M:      _m.M,
		})
	}
}

func ErrWrap(err error, fn func(m *M)) {

	if err == nil {
		return
	}

	var m = &KErr{}
	switch e := err.(type) {
	case *KErr:
		m = e
	case error:
		m.Msg = e.Error()
		m.Err = e
	}

	_m := &M{M: make(map[string]interface{})}
	fn(_m)

	var _tag = If(_m.tag == "", m.Tag, _m.tag).(string)
	m.Tag = ""

	if len(_m.M) == 0 {
		_m.M = nil
	}

	panic(&KErr{
		Sub:    m,
		Caller: funcCaller(),
		Msg:    _m.msg,
		Err:    m.tErr(),
		Tag:    _tag,
		M:      _m.M,
	})
}

func SWrap(err error, msg string, args ...interface{}) {
	if err == nil {
		return
	}

	var m = &KErr{}
	switch e := err.(type) {
	case *KErr:
		m = e
	case error:
		m.Msg = e.Error()
		m.Err = e
	}

	_m := fmt.Sprintf(msg, args...)
	panic(&KErr{
		Sub:    m,
		Caller: funcCaller(),
		Msg:    _m,
		Err:    m.tErr(),
	})
}

func Throw(err error) {
	if err == nil {
		return
	}

	var m = &KErr{}
	switch e := err.(type) {
	case *KErr:
		m = e
	case error:
		m.Err = e
		m.Msg = e.Error()
	}

	var _tag = m.Tag
	m.Tag = ""

	panic(&KErr{
		Sub:    m,
		Caller: funcCaller(),
		Err:    m.tErr(),
		Tag:    _tag,
	})
}

func P(d ...interface{}) {
	for _, i := range d {
		if i == nil {
			continue
		}

		if dt, err := json.MarshalIndent(i, "", "\t"); err != nil {
			panic(err)
		} else {
			fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
		}
	}
}
