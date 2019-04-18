package assert

import (
	"github.com/pubgo/gotry"
	"testing"
)

func a1()error {
	return gotry.Try(func() {
		Bool(true, "ok")
	}).Error()

}

func TestName(t *testing.T) {
	a := &Assert{}
	a.P(IfEquals(1, 2, 34))
	a.P(IfEquals(nil, nil))
	a.P(IfEquals(0, ""))
	a.P(IfEquals("", ""))
	gotry.Try(func() {
		MustNotError(a1())
	}).P()

}
