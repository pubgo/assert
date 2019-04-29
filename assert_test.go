package assert

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func a1() error {
	return _Try(func() {
		MustNotError(errors.New("sbhbhbh"))
		Bool(true, "好东西%d", 1)
	})
}

func TestName(t *testing.T) {
	P(IfEquals(1, 2, 34))
	P(IfEquals(nil, nil))
	P(IfEquals(0, ""))
	P(IfEquals("", ""))

	_Try(func() {
		Err(a1(), "ok111")
	}).LogStacks()

	//_Try(func() {
	//	Err(a1(), "oo")
	//}).LogStacks()
	//
	//_Try(func() {
	//	NotNil(a1())
	//}).LogStacks()
}

func TestPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(reflect.TypeOf(r).String())
			fmt.Println(r.(*KErr).err)
			fmt.Println(r)
		}
	}()
	//panic("ss")
	//panic([]string{"11", "33"})
	panic(&KErr{err: errors.New("sss")})
}
