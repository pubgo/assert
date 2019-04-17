package assert

import (
	"testing"
)

func TestName(t *testing.T) {
	a := &Assert{}
	a.P(IfEquals(1, 2, 34))
	a.P(IfEquals(nil, nil))
	a.P(IfEquals(0, ""))
	a.P(IfEquals("", ""))
	a.P("hello")
}
