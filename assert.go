package assert

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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
	if Cfg.Debug {
		log.Println(_m, funcCaller(callDepth))
	}

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

	if Cfg.Debug {
		log.Println(_m.msg, funcCaller(callDepth))
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
	case *KErr:
		m = e
	case error:
		m.msg = e.Error()
		m.err = e
	default:
		m.msg = fmt.Sprintf("type error %#v", e)
		m.err = errors.New(m.msg)
		m.tag = ErrTag.UnknownErr
	}

	_m := &M{M: make(map[string]interface{})}
	fn(_m)

	if len(_m.M) == 0 {
		_m.M = nil
	}

	if Cfg.Debug {
		log.Println(_m.msg, funcCaller(callDepth))
	}

	panic(&KErr{
		sub:    m,
		caller: funcCaller(callDepth),
		msg:    _m.msg,
		err:    m.tErr(),
		tag:    m.tTag(_m.tag),
		m:      _m.M,
	})
}

func Wrap(err error, msg string, args ...interface{}) error {
	if IsNil(err) {
		return nil
	}

	var m = &KErr{}
	switch e := err.(type) {
	case *KErr:
		m = e
	case error:
		m.msg = e.Error()
		m.err = e
	default:
		m.msg = fmt.Sprintf("type error %#v", e)
		m.err = errors.New(m.msg)
		m.tag = ErrTag.UnknownErr
	}

	if Cfg.Debug {
		log.Println(fmt.Sprintf(msg, args...), funcCaller(callDepth))
	}

	return &KErr{
		sub:    m,
		caller: funcCaller(callDepth),
		msg:    fmt.Sprintf(msg, args...),
		err:    m.tErr(),
	}
}

func Debug() {
	ErrHandle(recover(), func(err *KErr) {
		err.P()
	})
}

func Resp(fn func(err *KErr)) {
	ErrHandle(recover(), fn)
}

func Panic(fn func(m *M)) {
	err := recover()
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
	case string:
		m.err = errors.New(e)
		m.msg = e
	default:
		m.msg = fmt.Sprintf("type error %#v", e)
		m.err = errors.New(m.msg)
		m.tag = ErrTag.UnknownErr
	}

	_m := &M{M: make(map[string]interface{})}
	fn(_m)

	if len(_m.M) == 0 {
		_m.M = nil
	}

	panic(&KErr{
		sub:    m,
		caller: funcCaller(4),
		err:    m.tErr(),
		msg:    _m.msg,
		tag:    m.tTag(_m.tag),
		m:      _m.M,
	})
}

func ErrWrap(err interface{}, msg string, args ...interface{}) {
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
	default:
		m.msg = fmt.Sprintf("type error %#v", e)
		m.err = errors.New(m.msg)
		m.tag = ErrTag.UnknownErr
	}

	_m := fmt.Sprintf(msg, args...)

	if Cfg.Debug {
		log.Println(_m, funcCaller(callDepth))
	}
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
	case *KErr:
		m = e
	case error:
		m.err = e
		m.msg = m.err.Error()
	default:
		m.msg = fmt.Sprintf("type error %#v", e)
		m.err = errors.New(m.msg)
		m.tag = ErrTag.UnknownErr
	}

	if Cfg.Debug {
		log.Println(m.msg, funcCaller(callDepth))
	}

	panic(&KErr{
		sub:    m,
		caller: funcCaller(callDepth),
		err:    m.tErr(),
		tag:    m.tTag(m.tag),
	})
}

func P(d ...interface{}) {
	for _, i := range d {
		if IsNil(i) {
			return
		}

		if dt, err := json.MarshalIndent(i, "", "\t"); err != nil {
			panic(err)
		} else {
			fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
		}
	}
}
