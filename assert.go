package assert

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

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

func T(b bool, msg string, args ...interface{}) {
	if !b {
		return
	}

	_m := fmt.Sprintf(msg, args...)
	panic(&KErr{
		Caller: funcCaller(callDepth),
		Msg:    _m,
		Err:    errors.New(_m),
	})
}

func TT(b bool, fn func(m *M)) {
	if !b {
		return
	}

	_m := &M{M: make(map[string]interface{})}
	fn(_m)

	if len(_m.M) == 0 {
		_m.M = nil
	}

	panic(&KErr{
		Caller: funcCaller(callDepth),
		Msg:    _m.msg,
		Err:    errors.New(_m.msg),
		Tag:    _m.tag,
		M:      _m.M,
	})
}

func SWrap(err interface{}, fn func(m *M)) {
	if IsNil(err) {
		return
	}

	var m = &KErr{}
	switch e := err.(type) {
	case FnT:
		assertFn(e)
		T(reflect.TypeOf(e).NumOut() != 1, "the func num out error")
		_err := e()[0].Interface()
		if IsNil(_err) {
			return
		}

		m.Err = _err.(error)
		m.Msg = m.Err.Error()
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
		Caller: funcCaller(callDepth),
		Msg:    _m.msg,
		Err:    m.tErr(),
		Tag:    _tag,
		M:      _m.M,
	})
}

func Wrap(err error, msg string, args ...interface{}) error {
	return &KErr{
		Caller: funcCaller(callDepth),
		Msg:    fmt.Sprintf(msg, args...),
		Err:    err,
	}
}

func ErrWrap(err interface{}, msg string, args ...interface{}) {
	if IsNil(err) {
		return
	}

	var m = &KErr{}
	switch e := err.(type) {
	case FnT:
		assertFn(e)
		T(reflect.TypeOf(e).NumOut() != 1, "the func num out error")
		_err := e()[0].Interface()
		if IsNil(_err) {
			return
		}

		m.Err = _err.(error)
		m.Msg = m.Err.Error()
	case *KErr:
		m = e
	case error:
		m.Msg = e.Error()
		m.Err = e
	}

	_m := fmt.Sprintf(msg, args...)
	panic(&KErr{
		Sub:    m,
		Caller: funcCaller(callDepth),
		Msg:    _m,
		Err:    m.tErr(),
	})
}

func Throw(err interface{}) {
	if IsNil(err) {
		return
	}

	var m = &KErr{}
	switch e := err.(type) {
	case FnT:
		assertFn(e)
		T(reflect.TypeOf(e).NumOut() != 1, "the func num out error")
		_err := e()[0].Interface()
		if IsNil(_err) {
			return
		}

		m.Err = _err.(error)
		m.Msg = m.Err.Error()
	case *KErr:
		m = e
	case error:
		m.Err = e
		m.Msg = m.Err.Error()
	}

	var _tag = m.Tag
	m.Tag = ""

	panic(&KErr{
		Sub:    m,
		Caller: funcCaller(callDepth),
		Err:    m.tErr(),
		Tag:    _tag,
	})
}

func P(d ...interface{}) {
	for _, i := range d {
		if IsNil(i) {
			continue
		}

		if dt, err := json.MarshalIndent(i, "", "\t"); err != nil {
			panic(err)
		} else {
			fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
		}
	}
}
