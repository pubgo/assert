package assert

import (
	"github.com/pubgo/gotry"
	"testing"
)

func a1() error {
	return gotry.Try(func() {
		Bool(true, "好东西%d", 1) //strings.Join(CallerInfo(), "\n")
	}).Error()
}

func TestName(t *testing.T) {
	P(IfEquals(1, 2, 34))
	P(IfEquals(nil, nil))
	P(IfEquals(0, ""))
	P(IfEquals("", ""))
	gotry.Try(func() {
		Err(a1(), "ok111")
		MustNotError(a1())
	}).P()

}
