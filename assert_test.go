package assert

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func a1() {
	defer Panic(func(m *M) {
		m.Msg("test SWrap")
	})

	SWrap(errors.New("sbhbhbh"), func(m *M) {
		m.Msg("test shhh")
		m.M["ss"] = 1
		m.M["input"] = 1
	})
}

func TestName(t *testing.T) {
	defer Debug()

	P(IfEquals(1, 2, 34))
	P(IfEquals(nil, nil))
	P(IfEquals(0, ""))
	P(IfEquals("", ""))

	ErrHandle(KTry(a1), func(err *KErr) {
		err.P()
	})

	Throw(KTry(func() {
		SWrap(KTry(func() {
			SWrap(KTry(a1), func(m *M) {
				m.Msg("ok111")
				m.Tag("test tag")
			})
		}), func(m *M) {
			m.Msg("test 123")
		})
	}))
}

func TestTry(t *testing.T) {
	defer Debug()

	Cfg.Debug = true

	T(true, "sss")
}

func TestIf(t *testing.T) {
	fmt.Println(If(true, 1, "ss").(int))
	fmt.Println(If(true, FnOf(ToInt, "2"), "ss").(int))
	fmt.Println(If(true, FnOf(ToInt, "2"), "ss").(int))
	fmt.Println(reflect.TypeOf(FnOf(ToInt, "2")).Name())
}

func TestTask(t *testing.T) {
	defer Debug()

	Throw(Wrap(KTry(func() {
		ErrWrap(errors.New("dd"), "err ")
	}), "test wrap"))
}

func test123() {
	defer Panic(func(m *M) {
		m.Msg("test panic %d", 33)
	})

	ErrWrap(errors.New("hello error"), "sss")
}

func TestExpect11(t *testing.T) {
	defer Debug()

	Cfg.Debug=true

	test123()
}

func TestIsNil(t *testing.T) {
	defer Debug()

	var ss = func() map[string]interface{} {
		return make(map[string]interface{})
	}

	var ss1 = func() map[string]interface{} {
		return nil
	}

	var s = 1
	var ss2 map[string]interface{}
	t.Log(IsNil(1))
	t.Log(IsNil(1.2))
	t.Log(IsNil(nil))
	t.Log(IsNil("ss"))
	t.Log(IsNil(map[string]interface{}{}))
	t.Log(IsNil(ss()))
	t.Log(IsNil(ss1()))
	t.Log(IsNil(&s))
	t.Log(IsNil(ss2))
}

func TestResponce(t *testing.T) {
	defer Resp(func(err *KErr) {
		err.Tag()
		err.StackTrace()
	})

	T(true, "data handle")
}
