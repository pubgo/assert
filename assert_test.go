package assert

import (
	"errors"
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
	})

	_Try(func() {
		Err(a1(), "oo")
	})

	_Try(func() {
		NotNil(a1())
	})
}
