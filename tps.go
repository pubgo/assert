package assert

import "reflect"

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

	if !reflect.ValueOf(p).IsValid() {
		return true
	}

	return reflect.ValueOf(p).IsNil()
}
