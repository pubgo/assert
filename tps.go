package assert

import (
	"reflect"
	"strconv"
	"time"
)

func StrOf(args ...string) []string {
	return args
}

func ObjOf(args ...interface{}) []interface{} {
	return args
}

func IsPtr(p interface{}) bool {
	if IsNil(p) {
		return false
	}

	return reflect.TypeOf(p).Kind() == reflect.Ptr
}

func IsErr(p interface{}) bool {
	if IsNil(p) {
		return false
	}

	_, ok := p.(error)
	return ok
}

func IsNil(p interface{}) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			b = false
		}
	}()

	if p == nil {
		return true
	}

	if !reflect.ValueOf(p).IsValid() {
		return true
	}

	return reflect.ValueOf(p).IsNil()
}

func ToInt(p string) int {
	defer Panic(func(m *M) {
		m.Msg("ToInt Error")
	})

	r, err := strconv.Atoi(p)
	ErrWrap(err, "can not convert %s to int", p)
	return r
}

func ToTime(t string) time.Time {
	defer Panic(func(m *M) {
		m.Msg("ToTime Error")
	})

	tt, err := time.Parse("2006-01-02 15:04:05", t)
	ErrWrap(err, "time parse error")
	return tt
}
