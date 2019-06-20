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
		caller: funcCaller(callDepth),
		msg:    _m,
		err:    errors.New(_m),
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
		caller: funcCaller(callDepth),
		msg:    _m.msg,
		err:    errors.New(_m.msg),
		tag:    _m.tag,
		m:      _m.M,
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

		m.err = _err.(error)
		m.msg = m.err.Error()
	case *KErr:
		m = e
	case error:
		m.msg = e.Error()
		m.err = e
	}

	_m := &M{M: make(map[string]interface{})}
	fn(_m)

	var _tag = If(_m.tag == "", m.tag, _m.tag).(string)
	m.tag = ""

	if len(_m.M) == 0 {
		_m.M = nil
	}

	panic(&KErr{
		sub:    m,
		caller: funcCaller(callDepth),
		msg:    _m.msg,
		err:    m.tErr(),
		tag:    _tag,
		m:      _m.M,
	})
}

func Wrap(err error, msg string, args ...interface{}) error {
	return &KErr{
		caller: funcCaller(callDepth),
		msg:    fmt.Sprintf(msg, args...),
		err:    err,
	}
}

func Expect(fn func(), msg string, args ...interface{}) {
	assertFn(fn)
	err := KTry(fn)
	if IsNil(err) {
		return
	}

	var m = &KErr{}
	switch e := err.(type) {
	case *KErr:
		m = e
	case error:
		m.msg = e.Error()
		m.err = e
	}

	_m := fmt.Sprintf(msg, args...)
	panic(&KErr{
		sub:    m,
		caller: funcCaller(8),
		msg:    _m,
		err:    m.tErr(),
	})
}

func ErrWrap(err interface{}, msg string, args ...interface{}) {
	if IsNil(err) {
		return
	}

	var m = &KErr{}
	switch e := err.(type) {
	case FnT:
		assertFn(e)
		T(reflect.TypeOf(e).NumOut() != 1, "the func out num error")
		_err := e()[0].Interface()
		if IsNil(_err) {
			return
		}

		m.err = _err.(error)
		m.msg = m.err.Error()
	case *KErr:
		m = e
	case error:
		m.msg = e.Error()
		m.err = e
	}

	_m := fmt.Sprintf(msg, args...)
	panic(&KErr{
		sub:    m,
		caller: funcCaller(callDepth),
		msg:    _m,
		err:    m.tErr(),
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

		m.err = _err.(error)
		m.msg = m.err.Error()
	case *KErr:
		m = e
	case error:
		m.err = e
		m.msg = m.err.Error()
	}

	var _tag = m.tag
	m.tag = ""

	panic(&KErr{
		sub:    m,
		caller: funcCaller(callDepth),
		err:    m.tErr(),
		tag:    _tag,
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
