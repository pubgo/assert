package assert

import (
	"fmt"
	"strings"
)

func NewKErr(stack []string, err error) *KErr {
	return &KErr{_stacks: stack, err: err}
}

type KErr struct {
	_stacks []string
	err     error
}

func (e *KErr) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return ""
}

func (e *KErr) GetStacks() []string {
	return e._stacks
}

func (e *KErr) LogStacks() {
	fmt.Println(strings.Join(e._stacks, "\n"))
}
