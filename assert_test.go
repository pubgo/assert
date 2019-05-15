package assert

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func a1() error {
	return KTry(func() {
		SWrap(errors.New("sbhbhbh"), func(m *M) {
			m.Msg("test shhh")
			m.M["ss"] = 1
			m.M["input"] = 1
		})

		TT(true, func(m *M) {
			m.Msg("好东西%d", 1)
		})
	})
}

func TestName(t *testing.T) {
	P(IfEquals(1, 2, 34))
	P(IfEquals(nil, nil))
	P(IfEquals(0, ""))
	P(IfEquals("", ""))

	fmt.Println(IsNil(a1()))

	P(KTry(func() {
		Throw(KTry(func() {
			SWrap(KTry(func() {
				SWrap(a1(), func(m *M) {
					m.Msg("ok111")
					m.Tag("test tag")
				})
			}), func(m *M) {
				m.Msg("test 123")
			})
		}))
	}))
}

func TestType(t *testing.T) {
	var ss map[string]interface{}
	var s interface{}
	for _, i := range ObjOf("1", 0, errors.New(""), nil, []string{}, ss, s) {
		fmt.Println(IsNil(i))
	}
}

func TestTry(t *testing.T) {
	P(Try(func() {
		T(true, "sss")
	}))

	P(KTry(func() {
		T(true, "sss")
	}))
}

func TestIf(t *testing.T) {
	fmt.Println(If(true, 1, "ss").(int))
	fmt.Println(If(true, FnOf(ToInt, "2"), "ss").(int))
	fmt.Println(If(true, FnOf(ToInt, "2"), "ss").(int))
	fmt.Println(reflect.TypeOf(FnOf(ToInt, "2")).Name())
}
func TestTask(t *testing.T) {
	var ss interface{}
	sss, ok := ss.(FnT)
	fmt.Println(sss, ok)
}

func TestTasks(t *testing.T) {

	_fn := TaskOf(func(i int) {
		fmt.Println(i)
		T(i == 90, "90 error")
	}, func(err error) {
		Throw(err)
	})

	var task = NewTask(50, time.Second+time.Millisecond*10)
	for i := 0; i < 10000; i++ {
		if err := task.Do(_fn, i); err != nil {
			fmt.Println(err)
			break
		}
	}

	task.Wait()
}
