package assert

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const callDepth = 2

func ErrOf(msg string, args ...interface{}) *KErr {
	return &KErr{
		FuncCaller: funcCaller(),
		Msg:        fmt.Sprintf(msg, args...),
	}
}

func Bool(b bool, format string, args ...interface{}) {
	if b {
		panic(ErrOf(format, args...))
	}
}

func ErrWrap(err error, format string, args ...interface{}) {

	if err == nil {
		return
	}

	var m = &KErr{}
	switch e := err.(type) {
	case *KErr:
		m = e
	case error:
		m.Msg=e.Error()
		m.Err=e
	}

	panic(&KErr{
		Sub:        m,
		FuncCaller: funcCaller(),
		Msg:        fmt.Sprintf(format, args...),
		Err:        m.tErr(),
	})
}

func Err(err error) {
	if err == nil {
		return
	}

	var m = &KErr{}
	switch e := err.(type) {
	case *KErr:
		m = e
	case error:
		m.Err=e
		m.Msg=e.Error()
	}

	panic(&KErr{
		Sub:        m,
		FuncCaller: funcCaller(),
		Err:        m.tErr(),
	})
}

func P(d ...interface{}) {
	for _, i := range d {
		if dt, err := json.MarshalIndent(i, "", "\t"); err != nil {
			panic(err)
		} else {
			fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
		}
	}
}

var True = Bool
