package assert

import (
	"log"
	"strings"
)

func NewKErr() *KErr {
	return &KErr{}
}

type KErr struct {
	_stacks [] string
	err     error
}

func (e *KErr) SetErr(err error) {
	switch _e := err.(type) {
	case *KErr:
		e.err = _e.err
		for _, s := range _e._stacks {
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
	e._stacks = append(e._stacks, stack)
}

func (e *KErr) Error() string {
	if e.IsNil() {
		return ""
	}

	return e.err.Error()
}

func (e *KErr) IsNil() bool {
	return e.err == nil
}

func (e *KErr) Panic() {
	panic(e)
}

func (e *KErr) GetStacks() []string {
	if e.IsNil() {
		return []string{}
	}

	return e._stacks
}

func (e *KErr) LogStacks() {
	if e.IsNil() {
		return
	}

	for _, _s := range e.GetStacks() {
		log.Println(_s)
	}
	log.Println("error: ", e.Error())
	log.Println("************************")
}
