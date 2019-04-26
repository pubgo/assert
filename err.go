package assert

import (
	"fmt"
	"strings"
)

func NewKErr() *KErr {
	return &KErr{_stacks: make(chan string, MaxStack*2)}
}

type KErr struct {
	_stacks chan string
	err     error
}

func (e *KErr) SetErr(err error) {
	switch _e := err.(type) {
	case *KErr:
		e.err = _e.err
		close(_e._stacks)
		for s := range _e._stacks {
			e.AddStack(s)
		}
	case error:
		e.err = _e
	}
}

func (e *KErr) AddStack(stack string) {
	if len(e._stacks) > MaxStack {
		panic(strings.Join(e.GetStacks(), "\n"))
	}
	e._stacks <- stack
}

func (e *KErr) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return ""
}

func (e *KErr) GetStacks() (stack []string) {
	close(e._stacks)
	for s := range e._stacks {
		stack = append(stack, s)
	}
	return
}

func (e *KErr) LogStacks() {
	fmt.Println("error: ", e.Error())
	fmt.Println(strings.Join(e.GetStacks(), "\n"))
	fmt.Println()
}
