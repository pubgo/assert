package assert

import "reflect"

func StrOf(args ...string) []string {
	return args
}

func ObjOf(args ...interface{}) []interface{} {
	return args
}

func IsNil(p interface{}) (b bool) {
	defer func() {
		defer func() {
			if err := recover(); err != nil {
				b = false
			}
		}()

		if !reflect.ValueOf(p).IsValid() {
			b = true
			return
		}

		b = reflect.ValueOf(p).IsNil()
	}()
	return
}

type FnT func() []reflect.Value

func FnOf(fn interface{}, args ...interface{}) FnT {
	assertFn(fn)

	t := reflect.ValueOf(fn)
	return func() []reflect.Value {
		var vs []reflect.Value
		for i, p := range args {
			if p == nil {
				if t.Type().IsVariadic() {
					i = 0
				}

				vs = append(vs, reflect.New(t.Type().In(i)).Elem())
			} else {
				vs = append(vs, reflect.ValueOf(p))
			}
		}
		return t.Call(vs)
	}
}
