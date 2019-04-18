package assert

import (
	"fmt"
	"github.com/pubgo/gotry"
	"testing"
)

func a1(){
	Bool(true,"ok")
}

func TestName(t *testing.T) {
	a := &Assert{}
	a.P(IfEquals(1, 2, 34))
	a.P(IfEquals(nil, nil))
	a.P(IfEquals(0, ""))
	a.P(IfEquals("", ""))
	gotry.Try(func() {
		a1()
	}).Catch(func(err error) {
		fmt.Println(err.Error())
	})
	a1()


}
