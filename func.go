package assert

import (
	"reflect"
	"strconv"
	"time"
)

func If(b bool, t, f interface{}) interface{} {

	if b {
		if _t, ok := t.(FnT); ok {
			return _t()[0].Interface()
		}
		return t
	}

	if _f, ok := f.(FnT); ok {
		return _f()[0].Interface()
	}
	return f
}

func IsAllNil(args ...interface{}) bool {
	for _, _a := range args {
		if !IsNil(_a) {
			return false
		}
	}
	return true
}

func IsAllNotNil(args ...interface{}) bool {
	for _, _a := range args {
		if IsNil(_a) {
			return false
		}
	}
	return true
}

func IfEquals(args ...interface{}) bool {
	if len(args) < 2 {
		return true
	}

	if IsAllNil(args...) {
		return true
	}

	_t := args[0]
	if IsNil(_t) {
		return false
	}

	for _, _a := range args[1:] {
		if IsNil(_t) {
			return false
		}

		if _t != _a {
			return false
		}
	}
	return true
}

func IfIn(a interface{}, args ...interface{}) bool {
	if IsNil(a) == !IsAllNotNil(args...) {
		return true
	}

	if IsNil(a) {
		return false
	}

	_a := reflect.TypeOf(a).Kind()

	for _, arg := range args {
		if IsNil(arg) {
			return false
		}

		if _a == reflect.TypeOf(arg).Kind() {
			return true
		}
	}

	return false
}

func IfNotIn(a interface{}, args ...interface{}) bool {
	return !IfIn(a, args...)
}

func ToInt(p string) int {
	r, err := strconv.Atoi(p)
	SWrap(err, "can not convert %s to int,error(%s)", p, err)
	return r
}

func FnCost(f func()) time.Duration {
	t1 := time.Now()
	f()
	return time.Now().Sub(t1)
}

func Retry(num int, fn func()) (err error) {
	_t := fibonacci()
	for i := 0; i < num; i++ {
		if err = KTry(fn); err == nil {
			return
		}
		time.Sleep(time.Second * time.Duration(_t()))
	}
	return
}

func WaitFor(fn func(dur time.Duration) bool) {
	var _b = true
	for i := 0; _b; i++ {
		if err := Try(func() {
			_b = fn(time.Second * time.Duration(i))
		}).(*KErr); err != nil {
			err.Caller = funcCaller()
			err.Panic()
		}

		if !_b {
			return
		}

		time.Sleep(time.Second)
	}
	return
}

func Ticker(fn func(dur time.Time) time.Duration) {
	_dur := time.Duration(0)
	for i := 0; ; i++ {
		if err := Try(func() {
			_dur = fn(time.Now())
		}).(*KErr); err != nil {
			err.Caller = funcCaller()
			err.Panic()
		}

		if _dur < 0 {
			return
		}

		if _dur == 0 {
			_dur = time.Second
		}

		time.Sleep(_dur)
	}
}

type FnT func() []reflect.Value

func FnOf(fn interface{}, args ...interface{}) FnT {
	assertFn(fn)

	t := reflect.ValueOf(fn)
	return func() []reflect.Value {
		var vs []reflect.Value
		for i, p := range args {
			var _v reflect.Value
			if IsNil(p) {
				if t.Type().IsVariadic() {
					i = 0
				}
				_v = reflect.New(t.Type().In(i)).Elem()
			} else {
				_v = reflect.ValueOf(p)
			}

			vs = append(vs, _v)
		}
		return t.Call(vs)
	}
}

func AssertFn(fn interface{}) {
	assertFn(fn)
}
