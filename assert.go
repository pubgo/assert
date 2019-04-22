package assert

import (
	"encoding/json"
	"fmt"
	"reflect"
)

const callDepth = 2

func Bool(b bool, format string, args ...interface{}) {
	if b {
		panic(funcCaller() + fmt.Sprintf(format, args...))
	}
}

func Err(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}
	panic(err.Error() + "\n" + funcCaller() + fmt.Sprintf(format, args...))
}

func MustNotError(err error) {
	if err != nil {
		panic(err.Error() + "\n" + funcCaller())
	}
}

func P(d ...interface{}) {
	for _, i := range d {
		dt, err := json.MarshalIndent(i, "", "\t")
		MustNotError(err)
		fmt.Println(reflect.ValueOf(i).String(), "->", string(dt))
	}
}

var True = Bool
