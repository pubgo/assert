package assert

import (
	"fmt"
	"testing"
)

func a1() error {
	return _Try(func() {
		Bool(true, "好东西%d", 1)
	})
}

func TestName(t *testing.T) {
	P(IfEquals(1, 2, 34))
	P(IfEquals(nil, nil))
	P(IfEquals(0, ""))
	P(IfEquals("", ""))

	fmt.Println(_Try(func() {
		Err(a1(), "ok111")
	}))

	fmt.Println(_Try(func() {
		Err(a1(), "oo")
	}))
}
