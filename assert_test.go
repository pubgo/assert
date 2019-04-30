package assert

import (
	"testing"
)

func a1() error {
	return _Try(func() {
		Err(ErrOf("sbhbhbh"))
		Bool(true, "好东西%d", 1)
	})
}

func TestName(t *testing.T) {
	P(IfEquals(1, 2, 34))
	P(IfEquals(nil, nil))
	P(IfEquals(0, ""))
	P(IfEquals("", ""))

	P(_Try(func() {
		Err(_Try(func() {
			ErrWrap(_Try(func() {
				ErrWrap(a1(), "ok111")
			}), "test 123")
		}))
	}))
}

func TestIfNil(t *testing.T) {
	t.Log(IfNil(nil, nil, "test").(string))
}
