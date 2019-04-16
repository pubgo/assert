package assert

import (
	"testing"
)

func TestName(t *testing.T) {
	a := &Assert{}
	a.P("hello")
}
